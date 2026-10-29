# Sprint Plan (roadmap)

1-week sprints, Sunday → Thursday. Dates assume a start on **Sun 1 Nov 2026**; shift them if you start later. Every item below is an issue in [`scripts/backlog.psv`](../../scripts/backlog.psv), imported to GitHub by `scripts/setup-github.sh`.

| Sprint | Dates | Goal | Demo to client |
|---|---|---|---|
| **S0** | 1–5 Nov | Foundation: AWS account, repo, board, region/AI/payment spikes, collect menu data | Architecture walkthrough, timeline sign-off |
| **S1** | 8–12 Nov | Learn Go on the real codebase; catalogue API solid locally | Menu JSON with their real cakes (local) |
| **S2** | 15–19 Nov | DynamoDB + first AWS deploy + auto-deploy from `develop` | Live dev URL serving their menu |
| **S3** | 22–26 Nov | Delivery slots and orders (no payment yet) | Place a test order via API |
| **S4** | 29 Nov–3 Dec | Auth + owner admin API (menu, photos, orders, blocked days) | Owner changes a price, it updates everywhere |
| **S5** | 6–10 Dec | Payments + confirmation emails. API v1 frozen | Full test order paid with test card |
| **S6** | 13–17 Dec | AI assistant core: tool loop, slots, draft orders, history | Ask the bot for a cake, price and slot |
| **S7** | 20–24 Dec | AI evals + guardrails + monitoring. **Backend v1 complete** | Bot answers 30 tricky questions correctly |
| **S8** | 27–31 Dec | Three frontend experiments, client picks one | Click through 3 prototypes |
| **S9** | 3–7 Jan | Menu page and cake cards (hover/tap), detail page, cart | Browse the real menu on a phone |
| **S10** | 10–14 Jan | Checkout and chat widget | Order a cake end-to-end on the website |
| **S11** | 17–21 Jan | Admin UI, accessibility, SEO, speed | Owner runs a day of orders |
| **S12** | 24–28 Jan | Production, domain, go-live | 🎉 Launch |
| **S13** | 31 Jan–4 Feb | Final documentation and handover | Handover meeting |
| S14–S18 | Feb–Mar | *(Optional, quoted separately)* Mobile app with Expo | App on TestFlight / Play internal testing |

## Buffers and risks

- Weeks of learning Go are built into S1. If S1–S2 run long, cut "should/could" items, not tests.
- **Peak cake seasons overlap the plan.** Diwali (around 8 Nov, during S1), Christmas and New Year (S7–S8) are the client's busiest weeks. Expect slower replies and shorter demos then; schedule data requests before them.
- Republic Day (26 Jan) falls in the launch sprint S12. Plan go-live for later that week.
- Use the Christmas rush to ask the client what goes wrong with orders today. That is the best input for the AI assistant.
- Biggest external risk: payment gateway onboarding. KYC needs the client's business documents (PAN, bank account, GST if registered). **Start the application in S0.**
- Second risk: photos. Without good photos the frontend phase stalls. Ask in S0, chase in S4.

## Milestones

| Milestone | End of | Exit criteria |
|---|---|---|
| Foundation | S0 | Repo, board, AWS account, spikes answered with ADRs |
| Backend v1 | S7 | All `/v1` endpoints live on dev, payments in test mode, AI eval suite passing |
| Website on dev | S11 | Client signs off the design and flows |
| Go-live | S12 | Real customer order paid and delivered |
| Handover | S13 | Documentation pack delivered and walked through |
