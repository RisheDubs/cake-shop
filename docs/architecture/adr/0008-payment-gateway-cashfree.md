# 0008. Cashfree as the payment gateway

- Status: accepted
- Date: 2026-10-29
- Resolves: SPIKE-03 (research in `docs/research/spike-03-payments.md`)

## Context
Customers in Bangalore mostly pay by UPI. The shop needs UPI, cards and net banking, a hosted checkout so no payment details touch our servers, signed webhooks, refunds, and a sandbox we can build against before the client's KYC is approved. Razorpay, Cashfree and PayU were compared; all charge roughly 2% + GST.

## Decision
Use **Cashfree Payments**. The client agreed on 29 Oct 2026.

## Why
- **Sandbox needs no KYC**, so Sprint 5 can be built and tested while the client's activation is in progress.
- **Official Go SDK** (`github.com/cashfree/cashfree-pg`), plus a plain REST API.
- **T+1 standard settlement** (Razorpay and PayU list T+2), so the shop gets its money a day sooner.
- Rate about 1.95% + GST at the time of the decision.
- Payment Links are available for orders that still come in over WhatsApp.

## How we will use it (Sprint 5)
1. Customer confirms an order → API re-prices it and calls Cashfree **Create Order** with the amount, customer details and our `orderId`. Cashfree returns a `payment_session_id`.
2. The website opens Cashfree's checkout with that session id. The customer pays there by UPI, card or net banking.
3. Cashfree calls `POST /v1/webhooks/payments`. We verify the signature, then mark the order paid.
4. Cancellations use the **Refunds** API.

Implementation rules:
- **Money conversion at the edge only.** We store paise as integers; Cashfree's order amount is in rupees. Convert in one function (`85000` → `850.00`) with a unit test, never with floats in business logic.
- **Webhook verification:** HMAC-SHA256 over the timestamp header plus the **raw** request body, using the secret key, Base64-encoded, compared with `x-webhook-signature`. Read the raw body before any JSON parsing. Prefer the SDK's verify function, and test against sandbox webhooks, since Cashfree's docs show the concatenation in two slightly different ways.
- **Idempotency:** webhooks are delivered at least once. Store each event (`PAYEVT#<id>`, using `x-idempotency-header` or the payment id) with a conditional put; a duplicate does nothing.
- **Webhook is the source of truth**, not the browser coming back from checkout. The return page only shows "checking payment…" and reads the order status from our API.
- **Secrets:** App ID and Secret Key live in SSM Parameter Store (`/cakeshop/<stage>/cashfree/app-id`, `/cakeshop/<stage>/cashfree/secret-key`, SecureString). Sandbox keys locally via `.env`, never committed.
- **Environments:** sandbox (`sandbox.cashfree.com`) for local and dev; production (`api.cashfree.com`, keys starting `PROD_`) only in the prod stack.
- **Pin the API version** for requests and the webhook endpoint, and upgrade deliberately.

## Consequences
- The client must complete Cashfree KYC (business proofs for their entity type, bank verification) before launch. Start in Sprint 0 so approval is not on the critical path.
- Go-live needs production domain whitelisting and Cashfree's go-live checklist; this goes into Sprint 12.
- Confirm the final pricing (including any annual fee; public sources disagree) in the client's Cashfree dashboard during onboarding, and record it in `docs/project/cost-log.md`.
- Switching gateway later is contained: everything Cashfree-specific lives in `api/internal/payments`.
