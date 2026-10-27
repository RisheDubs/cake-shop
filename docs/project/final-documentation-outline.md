# Final Project Documentation (Phase 4) — Outline

Written in Sprint 13 from what accumulates in this repo. Each chapter lists where its raw material lives, so nothing has to be remembered from memory.

| # | Chapter | Audience | Built from |
|---|---|---|---|
| 1 | Executive summary: what was built and why | Client | sprint reviews, this outline |
| 2 | Project history: timeline, sprints, what changed and why | Client, you | `docs/sprints/*`, `CHANGELOG.md`, git tags |
| 3 | Architecture: diagrams, components, data, API | Developer | `docs/architecture/*` |
| 4 | Key decisions | Developer | `docs/architecture/adr/*` |
| 5 | Cost: build effort, running cost, cost per order, how to keep it low | Client | `docs/project/cost-log.md`, `time-log.md`, AWS Cost Explorer |
| 6 | How it works, step by step: customer journey, owner journey, AI assistant | Client, staff | flows in architecture doc, screenshots |
| 7 | Owner's manual: adding cakes, photos, blocking days, handling orders | Owner | admin UI |
| 8 | Operations runbook: deploys, rollbacks, alarms, backups, secrets rotation, "what to do if…" | Developer | `infra/`, CI workflows |
| 9 | Impact: orders via site vs before, time saved on phone/WhatsApp orders, AI chat stats | Client | metrics gathered from launch onward |
| 10 | Security and privacy summary | Client | architecture §7 |
| 11 | Known limitations and technical debt | Developer | issues labelled `debt` |
| 12 | Future expansion: mobile app, WhatsApp ordering, loyalty, multi-branch, analytics | Client | backlog `priority:could` |
| 13 | Appendix: environment variables, accounts and owners, glossary | Developer | `.env.example`, `infra/README.md` |

**Start collecting impact baselines in S0:** how many orders per week now, how they arrive (phone/WhatsApp/walk-in), time spent per order. Without a "before", there is no "impact".
