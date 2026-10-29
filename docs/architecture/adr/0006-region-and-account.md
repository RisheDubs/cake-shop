# 0006. Mumbai region and a separate AWS account

- Status: accepted
- Date: 2026-10-28
- Resolves: SPIKE-01

## Context
The client and all customers are in Bangalore, India. The developer already runs another project in `us-east-1` in an existing AWS account. Costs must be reportable per project, and the client should eventually own production.

## Decision
1. **Region: Asia Pacific (Mumbai) `ap-south-1`.** Fallback: Hyderabad `ap-south-2`.
2. **Account: a dedicated member account `cakeshop-dev`** in an AWS Organization owned by the developer's existing account. Production goes into a client-owned account at launch (Sprint 12).
3. All resources tagged `project=cakeshop` (set in `samconfig.toml`).

## Why
- Round trip Bangalore → Mumbai is roughly 20–40 ms; to `us-east-1` it is roughly 220–280 ms, paid on every API call.
- Customer data stays in India, which suits India's Digital Personal Data Protection Act.
- A separate account gives clean cost reporting, its own budget, no risk of touching the other project, and a simple handover (the account can leave the organization).

## Consequences
- The CloudFront custom-domain certificate must still be created in `us-east-1` (AWS requirement).
- Claude on Bedrock: Haiku 4.5 available with in-India inference; Haiku 5.5 only via the global cross-region profile (requests may be processed outside India). Decided in SPIKE-02 / ADR 0007.
- New member accounts appear in the main account's Budgets filters only after first usage, so the budget lives inside `cakeshop-dev`.
