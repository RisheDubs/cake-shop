# Definition of Done

An issue is **Done** only when all of these are true:

1. Acceptance criteria in the issue are met.
2. Code is merged to `develop` through a pull request.
3. CI is green (`gofmt`, `go vet`, `go test`).
4. New logic has tests (business rules: always; handlers: happy path + one error).
5. No secrets, keys or customer data committed.
6. Deployed to **dev** and checked there (from Sprint 2 onwards).
7. Docs updated if the API, architecture, costs or setup changed.
8. If a real choice was made, an ADR exists.

## Definition of Ready (before an issue enters a sprint)

- Clear "as a / I want / so that"
- Acceptance criteria written
- Size S or M (L must be split)
- No unknown blockers (otherwise create a spike first)
