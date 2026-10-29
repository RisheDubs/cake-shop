# API Contract (v1)

Base URL: `https://api.<domain>/v1` (dev: the `ApiUrl` stack output). JSON everywhere. Auth: `Authorization: Bearer <Cognito ID token>`.

Status: ✅ built · 🔜 planned (sprint)

## Public

| Method | Path | Purpose | Status |
|---|---|---|---|
| GET | `/health` | Liveness | ✅ |
| GET | `/v1/cakes?occasion=&tag=` | Menu, available cakes only | ✅ |
| GET | `/v1/cakes/{id}` | One cake: ingredients, allergens, sizes, prices | ✅ |
| GET | `/v1/cakes/{id}/quote?size=&qty=` | Server-side price | ✅ |
| GET | `/v1/slots?date=&cake=` | Delivery slots, lead-time aware | 🔜 S3 |
| POST | `/v1/chat` | AI assistant turn `{sessionId, message}` | 🔜 S6 |
| POST | `/v1/webhooks/payments` | Gateway callback (signature verified) | 🔜 S5 |

## Customer (token required)

| Method | Path | Purpose | Status |
|---|---|---|---|
| POST | `/v1/orders` | Create order → returns checkout URL | 🔜 S3/S5 |
| GET | `/v1/orders/{id}` | Order status (own orders only) | 🔜 S3 |
| GET | `/v1/me/orders` | Order history | 🔜 S3 |
| PUT | `/v1/me/profile` | Name, phone, address | 🔜 S4 |

## Admin (token + `admin` group)

| Method | Path | Purpose | Status |
|---|---|---|---|
| POST / PUT / DELETE | `/v1/admin/cakes[/{id}]` | Manage menu | 🔜 S4 |
| POST | `/v1/admin/cakes/{id}/image-upload-url` | Pre-signed S3 upload | 🔜 S4 |
| GET | `/v1/admin/orders?date=` | Orders for a day | 🔜 S4 |
| PATCH | `/v1/admin/orders/{id}` | Change status (baking, out for delivery …) | 🔜 S4 |
| PUT | `/v1/admin/slots/{date}` | Capacity / block a day | 🔜 S4 |

## Error shape (all endpoints)

```json
{ "error": { "code": "not_found", "message": "cake not found" } }
```

Codes: `bad_request` 400, `unauthorized` 401, `forbidden` 403, `not_found` 404, `unavailable` 409, `rate_limited` 429, `internal` 500.

## Example: cake

```json
{
  "id": "pistachio-rose",
  "name": "Pistachio Rose",
  "ingredients": ["flour", "pistachios", "rose water", "condensed milk", "cream", "sugar"],
  "tags": ["eggless"],
  "allergens": ["gluten", "dairy", "nuts"],
  "sizes": [
    { "id": "6in", "label": "6 inch", "serves": 8, "price": { "amount": 110000, "currency": "INR" } }
  ],
  "leadTimeHours": 48,
  "available": true
}
```

When the contract stabilises (end of Sprint 5), generate an `openapi.yaml` from it so the frontend and mobile app can generate typed clients.
