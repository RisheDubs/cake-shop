// Package payments integrates Cashfree Payments (ADR 0008). Built in Sprint 5.
//
// Everything Cashfree-specific stays inside this package, so the rest of the code only
// sees "create a payment for this order" and "this order is paid".
//
//   - The API calls Cashfree Create Order and returns the payment_session_id; the website
//     opens Cashfree's checkout with it (UPI, cards, net banking).
//   - Cashfree calls POST /v1/webhooks/payments. The signature (HMAC-SHA256 of timestamp +
//     raw body, Base64, header x-webhook-signature) is verified before the order is marked
//     paid. The webhook is the source of truth, not the browser redirect.
//   - Webhook handling is idempotent (same event twice = one state change).
//   - Amounts: paise internally, converted to rupees only when talking to Cashfree.
//   - No card or UPI details ever touch our servers.
package payments
