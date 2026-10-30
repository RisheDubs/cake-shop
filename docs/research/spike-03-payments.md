# SPIKE-03: Payment gateway for India

- Issue: SPIKE-03
- Time box: 4 hours (Day 2: research, Day 5: decision)
- Output: ADR 0008 + the client's merchant application started

## Question

Which Indian payment gateway should the cake shop use for online orders?

## Must have

- UPI (most Bangalore customers will pay this way), plus cards and net banking
- Hosted checkout page, so no card or UPI details touch our servers
- Webhooks with a signature we can verify in Go
- Test mode that works before the client's KYC is approved
- Refunds through the API (cancelled orders)
- Settlement to the client's bank account within 1–2 working days

## Shortlist

Published rates from mid-2026 comparison articles. **Confirm on each gateway's own pricing page on the day you check**, since rates and promos change.

| | Razorpay | Cashfree | PayU |
|---|---|---|---|
| Rate (UPI and domestic cards) | about 2% + GST | about 1.95% + GST (new-signup promos exist) | about 2% + GST |
| Setup fee | ₹0 | check (sources disagree) | not listed publicly |
| Annual fee | ₹0 | check (one source says ₹4,999/year) | not listed publicly |
| Standard settlement | T+2 | T+1 | T+2 to T+3 |
| Typical onboarding | hours to 5 days | 1–7 days | about 1 day |
| Test mode before KYC | ☐ check | ☐ check | ☐ check |
| Webhook signature docs | ☐ check | ☐ check | ☐ check |
| Go SDK or plain REST API | ☐ check | ☐ check | ☐ check |
| Refund API | ☐ check | ☐ check | ☐ check |
| Payment links (useful for orders taken on WhatsApp) | ☐ check | ☐ check | ☐ check |
| KYC documents for the client's business type | ☐ check | ☐ check | ☐ check |
| Website requirements before activation | ☐ check | ☐ check | ☐ check |

### Cost at the client's volume

Fill in once the client answers the "Orders today" tab.

| | Value |
|---|---|
| Online orders per month (estimate) | |
| Average order value (₹1200) | |
| Monthly online sales (₹10000) | |
| Gateway fee at 2% (₹200) | |
| GST on the fee at 18% (₹500) | |
| **Total payment cost per month (₹1000)** | |

Example: 100 orders × ₹1,200 = ₹1,20,000 sales → 2% = ₹2,400 + 18% GST = **₹2,832 a month**. This is usually the biggest running cost, bigger than AWS.

## Things to watch

- **Activation needs a website.** Gateways usually review the live site before going live, and many ask for visible policy pages: terms, privacy, refund/cancellation, shipping/delivery and contact details. Plan these pages early (Sprint 9–11) and apply to switch from test to live mode before launch.
- **Business type matters.** A registered business (GST, Udyam or shop licence) onboards more smoothly than an unregistered home baker. Check what the client has.
- **Cash on delivery and UPI QR at the door** may still be wanted. Ask the client; the order model can mark these as "pay on delivery".

## Findings

(Write what you learn here: links, screenshots of sign-up steps, answers from support chat.)

## Recommendation

(Day 5. Then write `docs/architecture/adr/0008-payment-gateway.md`.)

## Sources

- Mid-2026 comparisons: [SureCart: Razorpay vs Cashfree vs PayU](https://surecart.com/blog/razorpay-vs-cashfree-vs-payu-wordpress-india/), [ParityDeals: payment gateways in India 2026](https://www.paritydeals.com/blog/best-payment-gateways-india-2026/)
