// Package payments integrates the payment gateway. Built in Sprint 5.
//
// The gateway is chosen in spike SPIKE-03: an Indian gateway that supports UPI, cards
// and net banking, such as Razorpay, Cashfree or PayU, depending on fees and what the
// client can get approved for (KYC). Whatever the choice:
//
//   - The API creates a payment session for an order and returns the hosted-checkout URL.
//   - The gateway calls POST /v1/webhooks/payments; the signature is verified before
//     the order is marked paid. The webhook is the source of truth, not the browser redirect.
//   - Webhook handling is idempotent (same event twice = one state change).
//   - No card or UPI details ever touch our servers.
package payments
