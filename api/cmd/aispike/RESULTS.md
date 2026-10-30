# SPIKE-02 results: AI model for the assistant

Fill this in while running `go run ./cmd/aispike`. Then turn the conclusion into
`docs/architecture/adr/0007-ai-model.md`.

## Options tested

| | A: Haiku 5.5, global | B: Haiku 4.5, India only |
|---|---|---|
| Model ID | `global.anthropic.claude-haiku-5-5` | `in.anthropic.claude-haiku-4-5-20251001-v1:0` |
| Where requests are processed | Any AWS region worldwide | Mumbai + Hyderabad only |
| List price per 1M tokens (in / out) | $0.10 / $0.50 | $1.00 / $5.00 |

Prices: check the Amazon Bedrock pricing page on the day you test.

## Measurements

Run each question twice per model and note the second run (the first can be slower).

| # | Question | Model | Total time | Input tokens | Output tokens | Right tools? | Correct answer? |
|---|---|---|---|---|---|---|---|
| 1 | Eggless cake for ~12 people, what do you have and how much? | Nova Lite | 3.78 s | 2,792 | 193 |  list_cakes(eggless) → quote(pistachio-rose, 8in, 1) |  Pistachio Rose 8", serves 14, ₹1,600 |
| 1b | Eggless cake for ~12 people, how much? | Nova Lite | 3.52 s | 2,761 | 167 |  list_cakes(eggless) → quote(pistachio-rose, 8in, 1) |  ₹1,600 |
| 2 | Two 8-inch red velvet cakes? (expect ₹2,800) | Nova Lite | 3.52 s | 3,655 | 177 |  list_cakes → quote(red-velvet, 8in, 2) |  ₹2,800, but leaked `<thinking>` text |
| 3 | Do you have mango cake? (expect: not available) | Nova Lite | — | — | — |  failed |  error 424 ModelErrorException: invalid tool-use sequence |
| 4 | Does the pistachio cake have nuts? (expect: yes) | Nova Lite | 2.69 s | 2,034 | 114 |  list_cakes |  yes, but leaked `<thinking>` text |
| 5 | Weather in Bangalore? (expect: polite redirect) | Nova Lite | 2.16 s | 615 | 61 |  no tools (correct) |  redirected to cakes, but leaked `<thinking>` text |

## Notes (Nova Lite)

- Tool use works: it always looked up prices with `quote` instead of guessing, and every price was correct.
- Speed: first call ~2.0–2.3 s, follow-up calls ~0.6–0.8 s. Full answers in 2.2–3.8 s.
- Average successful conversation: ~2,370 input tokens and ~140 output tokens.
- **Problem 1: `<thinking>` leaks.** 3 of 5 answers included the model's internal `<thinking>…</thinking>` text.
  The real assistant must strip it before replying, or the prompt must stop it.
- **Problem 2: tool-use error.** "Do you have mango cake?" failed with HTTP 424 `ModelErrorException`
  (model produced an invalid tool-use sequence). Needs: retry once, then a friendly fallback message.
  Re-run to see if it is random or repeatable.
- Mango is unavailable, so `list_cakes` doesn't return it. The model must say "not on the menu right now"
  instead of inventing it. Could not check this because of the error.

## Cost per conversation

Average tokens per conversation × price:

| | A | B |
|---|---|---|
| Avg input tokens | | |
| Avg output tokens | | |
| Cost per conversation ($) | | |
| × 1,000 conversations a month ($) | | |

## Claude API directly (not tested, compared on paper)

| | Bedrock (A or B) | Claude API |
|---|---|---|
| Billing | Same AWS bill, which the client will own | Separate Anthropic account and invoice |
| Auth | IAM role on the Lambda, no key to store | API key stored in SSM |
| Data location | Global (A) or India (B) | Anthropic's infrastructure |
| Setup | Use-case form once per account/organization | Sign-up + credits |

## Notes

(Anything surprising: errors, a wrong tool call, how the answer read.)

## Decision

(Which option and why, in 3–4 sentences.)
