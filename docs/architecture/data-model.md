# Data Model (DynamoDB single table)

Table `cakeshop-<stage>`, on-demand billing, keys `PK` (partition) and `SK` (sort), one index `GSI1` (`GSI1PK`, `GSI1SK`), TTL attribute `ttl`.

Why one table: every screen needs one fast query, and at 100 customers everything fits easily. Learn the access patterns first, then the keys follow.

## Items

| Entity | PK | SK | GSI1PK | GSI1SK | Notes |
|---|---|---|---|---|---|
| Cake | `CAKE#<cakeId>` | `META` | `CATALOG` | `<sortOrder>#<name>` | name, description, ingredients, allergens, tags, occasions, imageKey, leadTimeHours, available |
| Cake size | `CAKE#<cakeId>` | `SIZE#<sizeId>` | | | label, serves, priceAmount, currency |
| Customer | `CUST#<cognitoSub>` | `PROFILE` | | | name, phone, default address |
| Order | `ORDER#<orderId>` | `META` | `CUST#<cognitoSub>` | `ORDER#<createdAt>` | status, slot, totals, address, payment ref |
| Order line | `ORDER#<orderId>` | `LINE#<n>` | | | snapshot: cake name, size label, unit price, qty, message on cake |
| Orders by day (owner) | | | `DAY#<yyyy-mm-dd>` | `<slotStart>#<orderId>` | set on Order META, powers the owner's "today" view |
| Slot day | `SLOTDAY#<yyyy-mm-dd>` | `SLOT#<start>` | | | capacity, booked (atomic counter), blocked flag |
| Settings | `CONFIG` | `SHOP` | | | opening hours, default capacity, delivery zones and fees |
| Chat message | `CHAT#<sessionId>` | `MSG#<timestamp>` | | | role, content, tool calls; `ttl` = +30 days |
| Payment event | `PAYEVT#<gatewayEventId>` | `META` | | | stored first → makes webhooks idempotent; `ttl` = +90 days |

## Access patterns

| # | Need | How |
|---|---|---|
| 1 | Menu (all cakes) | Query `GSI1PK = CATALOG` |
| 2 | One cake with sizes | Query `PK = CAKE#id` (META + SIZE# items) |
| 3 | Customer's orders | Query `GSI1PK = CUST#sub`, newest first |
| 4 | One order with lines | Query `PK = ORDER#id` |
| 5 | Owner: orders for a day | Query `GSI1PK = DAY#date` |
| 6 | Slots for a day | Query `PK = SLOTDAY#date` |
| 7 | Book a slot safely | `UpdateItem` with condition `booked < capacity` (no double booking) |
| 8 | Chat history | Query `PK = CHAT#session`, last N |
| 9 | Webhook seen before? | `PutItem` with `attribute_not_exists(PK)` |

## Money

All amounts are integers in minor units (`18500` = 185.00 QAR). Never floats.
