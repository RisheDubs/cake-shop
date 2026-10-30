// Command aispike is the SPIKE-02 experiment (Sprint 0, Day 3).
//
// It asks Claude on Amazon Bedrock a customer question, lets Claude call two tools
// that use the REAL cake catalogue and pricing code (cakes.Service), and prints
// every step with timings and token counts. Throwaway code: the real assistant is
// built properly in Sprint 6 under internal/chat.
//
// Run (Git Bash, after `export AWS_PROFILE=cakeshop-dev && aws sso login`):
//
//	go run ./cmd/aispike -q "I need an eggless cake for 12 people. How much?"
//	go run ./cmd/aispike -model in.anthropic.claude-haiku-4-5-20251001-v1:0 -q "..."
package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/document"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"

	"cakeshop/api/internal/cakes"
	"cakeshop/api/internal/store"
)

const systemPrompt = `You are the ordering assistant for a cake shop in Bangalore.
Rules:
- Never state a price, size, ingredient or allergen from memory. Always call a tool.
- Prices from tools are in paise. Show them to the customer in rupees, e.g. 85000 -> ₹850.
- Keep answers short and friendly. If something is not on the menu, say so.`

func main() {
	question := flag.String("q", "I need an eggless cake for about 12 people. What do you have and how much is it?", "customer question")
	model := flag.String("model", "global.anthropic.claude-haiku-5-5", "Bedrock model or inference profile ID")
	region := flag.String("region", "ap-south-1", "AWS region to call Bedrock in")
	flag.Parse()

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(*region))
	if err != nil {
		log.Fatalf("load AWS config (did you run aws sso login?): %v", err)
	}
	client := bedrockruntime.NewFromConfig(cfg)
	shop := cakes.NewService(store.NewMemory(store.SeedCakes()))

	fmt.Printf("Model:    %s\nRegion:   %s\nCustomer: %s\n\n", *model, *region, *question)

	messages := []types.Message{{
		Role:    types.ConversationRoleUser,
		Content: []types.ContentBlock{&types.ContentBlockMemberText{Value: *question}},
	}}

	var totalIn, totalOut int32
	start := time.Now()

	// The tool-use loop: ask Claude; if it wants a tool, run it and send the result back.
	for turn := 1; turn <= 6; turn++ {
		t0 := time.Now()
		out, err := client.Converse(ctx, &bedrockruntime.ConverseInput{
			ModelId:         aws.String(*model),
			System:          []types.SystemContentBlock{&types.SystemContentBlockMemberText{Value: systemPrompt}},
			Messages:        messages,
			ToolConfig:      toolConfig(),
			InferenceConfig: &types.InferenceConfiguration{MaxTokens: aws.Int32(600)},
		})
		if err != nil {
			explain(err)
			os.Exit(1)
		}
		if out.Usage != nil {
			totalIn += aws.ToInt32(out.Usage.InputTokens)
			totalOut += aws.ToInt32(out.Usage.OutputTokens)
		}
		fmt.Printf("── call %d: %v, stop reason %q\n", turn, time.Since(t0).Round(time.Millisecond), out.StopReason)

		msg, ok := out.Output.(*types.ConverseOutputMemberMessage)
		if !ok {
			log.Fatal("unexpected output type")
		}
		messages = append(messages, msg.Value)

		if out.StopReason != types.StopReasonToolUse {
			for _, block := range msg.Value.Content {
				if text, ok := block.(*types.ContentBlockMemberText); ok {
					fmt.Printf("\nAssistant:\n%s\n", text.Value)
				}
			}
			break
		}

		// Run every tool Claude asked for, and collect the results in one user message.
		var results []types.ContentBlock
		for _, block := range msg.Value.Content {
			use, ok := block.(*types.ContentBlockMemberToolUse)
			if !ok {
				continue
			}
			name := aws.ToString(use.Value.Name)
			var input map[string]any
			_ = use.Value.Input.UnmarshalSmithyDocument(&input)
			fmt.Printf("   tool  %s %v\n", name, input)

			result, toolErr := runTool(shop, name, input)
			status := types.ToolResultStatusSuccess
			if toolErr != nil {
				result = map[string]any{"error": toolErr.Error()}
				status = types.ToolResultStatusError
			}
			results = append(results, &types.ContentBlockMemberToolResult{Value: types.ToolResultBlock{
				ToolUseId: use.Value.ToolUseId,
				Content:   []types.ToolResultContentBlock{&types.ToolResultContentBlockMemberJson{Value: document.NewLazyDocument(result)}},
				Status:    status,
			}})
		}
		messages = append(messages, types.Message{Role: types.ConversationRoleUser, Content: results})
	}

	fmt.Printf("\nTotal: %v, %d input tokens, %d output tokens\n",
		time.Since(start).Round(time.Millisecond), totalIn, totalOut)
}

// toolConfig describes the tools Claude may call. Same idea as Sprint 6, just fewer tools.
func toolConfig() *types.ToolConfiguration {
	spec := func(name, desc string, schema map[string]any) types.Tool {
		return &types.ToolMemberToolSpec{Value: types.ToolSpecification{
			Name:        aws.String(name),
			Description: aws.String(desc),
			InputSchema: &types.ToolInputSchemaMemberJson{Value: document.NewLazyDocument(schema)},
		}}
	}
	return &types.ToolConfiguration{Tools: []types.Tool{
		spec("list_cakes", "List cakes on sale with ingredients, allergens, tags and sizes (with how many people each serves and price in paise).",
			map[string]any{
				"type": "object",
				"properties": map[string]any{
					"tag":      map[string]any{"type": "string", "description": "Optional filter, e.g. eggless"},
					"occasion": map[string]any{"type": "string", "description": "Optional filter, e.g. birthday, diwali"},
				},
			}),
		spec("quote", "Exact price in paise for a cake, size and quantity.",
			map[string]any{
				"type": "object",
				"properties": map[string]any{
					"cake_id":  map[string]any{"type": "string"},
					"size_id":  map[string]any{"type": "string"},
					"quantity": map[string]any{"type": "integer", "minimum": 1},
				},
				"required": []string{"cake_id", "size_id", "quantity"},
			}),
	}}
}

// runTool calls the real shop code. The model never sees anything we don't return here.
func runTool(shop *cakes.Service, name string, in map[string]any) (any, error) {
	str := func(k string) string { s, _ := in[k].(string); return s }
	switch name {
	case "list_cakes":
		list, err := shop.List(cakes.Filter{Tag: str("tag"), Occasion: str("occasion"), OnlyAvailable: true})
		if err != nil {
			return nil, err
		}
		return toMap(map[string]any{"cakes": list})
	case "quote":
		qty := 1
		if n, ok := in["quantity"].(float64); ok {
			qty = int(n)
		}
		q, err := shop.Quote(str("cake_id"), str("size_id"), qty)
		if err != nil {
			return nil, err
		}
		return toMap(q)
	default:
		return nil, fmt.Errorf("unknown tool %q", name)
	}
}

// toMap turns our Go structs into plain JSON-shaped data (uses our json tags).
func toMap(v any) (any, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	var out any
	return out, json.Unmarshal(b, &out)
}

func explain(err error) {
	fmt.Fprintf(os.Stderr, "\nBedrock call failed: %v\n\n", err)
	msg := err.Error()
	var hint string
	switch {
	case strings.Contains(msg, "AccessDenied"):
		hint = "Model access not set up yet: submit the Anthropic use-case form in the Bedrock console (Day 3, Part 1), wait a few minutes, retry."
	case strings.Contains(msg, "ValidationException") && strings.Contains(strings.ToLower(msg), "model"):
		hint = "Model ID not available in this region. Check the ID in Bedrock console -> Model catalog / Cross-region inference."
	case strings.Contains(msg, "expired") || strings.Contains(msg, "SSO"):
		hint = "Your SSO session expired: run  aws sso login"
	case errors.Is(err, context.DeadlineExceeded):
		hint = "Timed out. Try again."
	}
	if hint != "" {
		fmt.Fprintln(os.Stderr, "Hint:", hint)
	}
}
