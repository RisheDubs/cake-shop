// Package orders handles carts, delivery slots and order lifecycle. Built in Sprint 3.
//
// Planned pieces:
//
//   - Slot rules: daily capacity per slot, blocked dates set by the owner, and each
//     cake's LeadTimeHours (a 48h cake ordered today cannot be delivered tomorrow).
//   - Order states: draft -> pending_payment -> paid -> baking -> out_for_delivery -> delivered
//     (plus cancelled / refunded). Transitions are validated in one place.
//   - Prices are re-quoted from cakes.Service at checkout; the client never sends a price.
//   - Each order stores a snapshot of cake name, size and price so later menu edits
//     do not change historical orders.
package orders
