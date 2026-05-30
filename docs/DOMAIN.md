# Domain

## Scope

This document captures business rules that are explicit in code today, plus a small set of implementation notes that future agents should treat carefully.

If a rule is not described here or enforced in code, treat it as unknown rather than assumed.

## Core Entities

### Business

- A business is the top-level tenant boundary used across admin and shopper flows.
- Most admin queries are scoped by `business_id` from authenticated request context.

### Table

Tables represent physical restaurant tables and are identified by both:

- numeric table ID
- opaque token used by shopper clients

Current table fields include:

- `id`
- `business_id`
- `token`
- `status`
- `qr_code_url`
- `table_no`
- `description`

### Order

Orders belong to:

- one business
- one table

Current order fields include:

- `id`
- `business_id`
- `table_id`
- `table_no`
- `status`
- `total`

### Notification

Notifications are stored in Postgres and may also be broadcast over websocket.

Current notification fields include:

- `id`
- `message`
- `type`
- `is_read`
- `meta_data`
- `business_id`

## Table Statuses

Defined in `internal/types/tables.go`:

- `available`
- `reserved`
- `occupied`
- `cleaning`
- `out_of_order`

Important note:

- These statuses exist as types today, but not all transitions are enforced in handlers yet.
- The most concrete implemented transition is in `MarkTableOrdersAsPaid`, which sets the table back to `available`.

## Order Statuses

Defined in `internal/types/orders.go`:

- `pending`
- `preparing`
- `ready`
- `paid`

Current enforced behavior:

- new shopper orders are created as `pending`
- admin order-status updates accept only valid statuses except `paid`
- `paid` is applied through the table payment flow, not through the generic order status patch route

Practical interpretation for agents:

- use `PATCH /api/v1/orders/{order_id}/status` for kitchen/dashboard progression such as `preparing` or `ready`
- use `POST /api/v1/tables/{id}/paid` to settle a table and mark its active orders as `paid`

## Table Payment Flow

The most explicit end-of-meal rule in code today is:

1. mark a table as paid through `POST /api/v1/tables/{id}/paid`
2. set that table's status to `available`
3. mark all non-paid orders for that table as `paid`

This is the clearest current lifecycle rule in the repository.

## Shopper Access Model

Shopper requests do not authenticate as users.

Instead, they authenticate through `TABLE-TOKEN`, which currently maps to:

- `business_id`
- `table_id`
- `table_no`

That means shopper actions are table-scoped, not user-scoped.

## Notification Rules

What is explicit in code today:

- creating a shopper order inserts a notification row
- that notification uses type `new_order`
- the server also publishes a websocket event on topic `admin`

Notification types currently defined:

- `new_order`
- `update_order_status`

Important caution:

- `update_order_status` exists as a type constant, but this repository snapshot does not show a matching publish flow for it
- do not assume all notification types are actively emitted unless you confirm in code

## Storage Rules

Current behavior:

- table creation generates a QR code
- the QR PNG is uploaded through `internal/utils/upload.go`
- the code currently uses AWS S3 SDK configuration and stores the resulting URL

Planned direction from project owner:

- object storage is expected to move toward Cloudflare R2

Guidance for future changes:

- prefer S3-compatible abstraction boundaries
- avoid spreading AWS-specific assumptions beyond the storage helper unless necessary

## Rules That Are Still Unclear

These areas are not fully documented or enforced in the current codebase:

- exact table-status transition rules beyond payment resetting to `available`
- whether `reserved`, `occupied`, `cleaning`, and `out_of_order` are already used by the frontend
- whether order totals are recalculated elsewhere or only persisted as DB values
- whether websocket topics beyond `admin` are planned or already used externally

Future agents should treat these as open questions, not stable facts.
