# 0001. Go for the backend

- Status: accepted
- Date: 2026-10-27

## Context
The developer wants to learn a new language (Rust, Go, Elixir and Gleam were considered). The backend must run cheaply on AWS, be done in about 3 months, and be maintainable by someone else later.

## Decision
Go, using the standard library HTTP router, deployed to AWS Lambda.

## Consequences
- Easy to learn, first-class AWS SDK and Lambda support, fast cold starts.
- Rust rejected for this project (learning curve vs deadline); Elixir rejected because it suits always-on servers better than Lambda; Gleam rejected because the ecosystem is too young for payments/AWS.
- Frontend stays TypeScript (Next.js), so the project uses two languages.
