# Website (Phase 3, starts Sprint 8)

Next.js + TypeScript + Tailwind, hosted on Amplify. Empty until the backend is complete.

## Frontend experiments (Sprint 8)

Build three quick prototypes against the **real dev API**, each on its own branch (`ui/<name>`), each deployed as an Amplify preview so the client can click through them on their phone.

| Branch | Idea | What it tests |
|---|---|---|
| `ui/gallery` | Gallery-first: large photo grid, hover/tap reveals ingredients + sizes + price | Classic, fast to browse |
| `ui/story` | Story-scroll: one full-screen cake at a time, details slide in | Most "premium", cake as hero |
| `ui/chat-first` | Assistant is the homepage ("What are you celebrating?"), menu below | How much customers lean on the AI |

Optional fourth if time allows: `ui/builder`, a custom-cake builder.

### Scorecard (fill with the client)

| Criteria (1–5) | Gallery | Story | Chat-first |
|---|---|---|---|
| Cakes look irresistible | | | |
| Find a cake in under 30 s on a phone | | | |
| Ingredients/allergens easy to see | | | |
| Fits the brand | | | |
| Effort to finish | | | |
| **Total** | | | |

Record the winner as an ADR. The losing branches stay in git as history.
