// Package chat is the AI assistant ("Ask the baker"). Built in Sprint 6.
//
// See docs/architecture/README.md, section 4.3, for the full design.
//
// It runs as its own Lambda (longer timeout, separate throttling) behind POST /v1/chat,
// and calls a Haiku-class Claude model through Amazon Bedrock or the Claude API
// (decided in SPIKE-02; an API key would live in SSM Parameter Store).
//
// The model never states prices, sizes, ingredients or delivery times from its own
// knowledge. It must call these tools, which are thin wrappers over the real services:
//
//	list_cakes(occasion?, tag?)                  -> cakes.Service.List
//	get_cake(cake_id)                            -> cakes.Service.Get
//	quote(cake_id, size_id, quantity)            -> cakes.Service.Quote
//	check_delivery_slots(date, cake_id)          -> orders.Service.AvailableSlots
//	create_draft_order(cake_id, size_id, qty,
//	                   slot_id, message_on_cake) -> orders.Service.CreateDraft (checkout link)
//
// Payment never happens inside the chat; the customer confirms on the checkout page.
// Conversation history is stored in DynamoDB (PK CHAT#<session>) with a 30-day TTL.
// Guardrails: max turns per session, max tokens per reply, per-IP rate limiting, and a
// system prompt that keeps it on cake-shop topics and replies in English (Hindi or
// Kannada too, if the client wants).
// Quality: an eval test set of real customer questions with expected tool calls runs
// in CI before any prompt change ships.
package chat
