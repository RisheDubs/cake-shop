# 0004. The AI assistant only uses tools for facts

- Status: accepted
- Date: 2026-10-27

## Context
An AI that invents a price, an ingredient (allergy risk) or a delivery time causes real harm and lost trust.

## Decision
The assistant gets prices, sizes, ingredients, allergens and slots exclusively through tool calls that wrap the same Go services the website uses. It may create a draft order and return a checkout link; it never takes payment. A Haiku-class Claude model is used for cost and speed.

## Consequences
- One source of truth; a menu edit updates web, app and AI at once.
- Needs an eval test set to catch prompt regressions.
- Slightly more engineering than a "just answer" chatbot.
