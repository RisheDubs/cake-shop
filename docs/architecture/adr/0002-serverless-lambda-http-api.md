# 0002. Serverless: API Gateway HTTP API + Lambda + Web Adapter

- Status: accepted
- Date: 2026-10-27

## Context
Traffic is tiny and bursty (50–100 customers). Paying for an always-on server wastes money and needs patching.

## Decision
API Gateway **HTTP API** in front of Go Lambdas (arm64). The Go code is a normal `net/http` server; the **AWS Lambda Web Adapter** layer runs it on Lambda unchanged. Two functions: `api` and `chat`.

## Consequences
- Running cost near zero; scales automatically.
- The same binary runs locally (`go run`) and on AWS, which makes learning and debugging simpler.
- Cold starts of ~100–300 ms are acceptable for this use.
- If the shop ever needs WebSockets or long-running work, revisit (App Runner / ECS).
