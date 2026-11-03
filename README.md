# Cake Shop 🎂

Online cake shop with an AI ordering assistant. Cakes are the hero: customers browse, hover or tap a cake to see ingredients, sizes and prices, order with a delivery slot, and can ask the assistant ("Ask the baker") for recommendations, prices and delivery times.

| | |
|---|---|
| Backend | Go on AWS Lambda + API Gateway, DynamoDB, Cognito, S3 |
| AI | Claude (Haiku-class) with tool calling, facts only from the real menu |
| Frontend | Next.js + TypeScript on Amplify (Phase 3) |
| Mobile | Expo / React Native on the same API (optional) |
| Process | Solo Scrumban, 1-week sprints, all in GitHub |

## The Architecture
<img width="1095" height="576" alt="image" src="https://github.com/user-attachments/assets/4b1a5bb0-d5a0-47ce-ab4e-897b249e9a29" />


## Repository layout

```
cake-shop/
├── api/                 Go backend (one module, two Lambdas: api + chat)
│   ├── cmd/server/      entry point (runs locally and on Lambda unchanged)
│   ├── internal/
│   │   ├── cakes/       catalogue model, pricing rules, quotes
│   │   ├── orders/      slots, orders, state machine        (Sprint 3)
│   │   ├── payments/    gateway + webhooks                   (Sprint 5)
│   │   ├── chat/        AI assistant + tools                 (Sprint 6)
│   │   ├── store/       repositories: memory now, DynamoDB in Sprint 2
│   │   ├── httpapi/     routes, handlers, middleware
│   │   └── platform/    config and shared plumbing
│   ├── requests.http    try the API from VS Code
│   └── Makefile
├── infra/               AWS SAM template, deploy config
├── web/                 website (Phase 3)
├── mobile/              mobile app (optional)
├── docs/                architecture, process, sprints, project logs
├── scripts/             GitHub backlog setup
└── .github/             CI, issue and PR templates
```

## Quick start

**Install once:** [Go (latest)](https://go.dev/dl/), [Git](https://git-scm.com), [VS Code](https://code.visualstudio.com), [GitHub CLI](https://cli.github.com). From Sprint 2: [AWS CLI](https://aws.amazon.com/cli/) and [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html).

```bash
code cake-shop                 # open in VS Code, accept the recommended extensions
cd api
cp .env.example .env
go test ./...                  # all green
go run ./cmd/server            # http://localhost:8080/v1/cakes
```

Then open `api/requests.http` and click **Send Request**, or press **F5** to debug the API.

## Git setup (first time)

```bash
1
git add -A
git commit -m "chore: project scaffold, architecture v1, process"
git switch -c develop
gh repo create cake-shop --private --source=. --push
git push -u origin main develop
bash scripts/setup-github.sh   # labels, Sprint 0-13 milestones, full backlog as issues
```

Then protect `main` and `develop` and create the Project board, see [docs/process/git-workflow.md](docs/process/git-workflow.md).

## Where to read next

1. [Architecture blueprint](docs/architecture/README.md)
2. [How we work](docs/process/workflow.md) and the [sprint plan](docs/process/sprint-plan.md)
3. [Sprint 0](docs/sprints/sprint-00.md)
