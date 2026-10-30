# 0007. AI assistant runs on Amazon Bedrock; model is configurable

- Status: accepted (model choice provisional)
- Date: 2026-10-30
- Resolves: SPIKE-02 (results in api/cmd/aispike/RESULTS.md)

## Context
The assistant must answer menu, price and delivery questions using tools only.
Claude access on the new cakeshop-dev account is pending (Anthropic use-case form / account access).

## Decision
- Use Amazon Bedrock (Converse API) from ap-south-1, not a direct model API:
  one AWS bill the client will own, IAM auth with no API key, Mumbai region.
- The model ID is configuration (env var `AI_MODEL_ID`), not code.
- Development uses Amazon Nova (<Lite or Pro, from results>) until Claude access is granted.
- Final model is chosen in Sprint 7 by running the eval suite against Nova and Claude
  (Haiku 5.5 global, or Haiku 4.5 India-only if the client wants data kept in India).

## Consequences
- No lock-in: switching model is a config change, checked by evals.
- Claude access must be requested in the client's production account too (Sprint 12).
- Results so far: <2–3 lines from RESULTS.md: speed, cost, any rule it broke>.
