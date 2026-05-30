# API

## Base URLs

Current intended route prefixes:

- admin API: `https://example.com/api/v1`
- shopper API: `https://example.com/s/api/v1`

For local development, these same route shapes apply to whatever host and port the server is running on.

## Auth Model

### Admin API

Admin requests are protected by `server/verify_middleware.go`.

Supported request modes:

1. HMAC mode for real clients
2. `X-ADMIN-SECRET` mode for internal developer tooling

#### HMAC mode

Expected headers:

- `X-HMAC-SIGNATURE`
- `X-USER-EMAIL`

How it works:

- the client and server share the same secret
- the server uses `SECRET_API_KEY` from env
- the request signature is derived from the URL string
- the client must compute the signature the same way as the server

Current server-side algorithm:

1. build a string in the form `URL:<request-url>`
2. for notification URLs, the code uses `URL:<path>` instead of the full URL string
3. compute HMAC-SHA256 using `SECRET_API_KEY`
4. hex-encode the digest
5. send the result in `X-HMAC-SIGNATURE`

Important note for agents:

- keep client and server implementations exactly aligned
- the notification-path special case is easy to miss and can break integrations

#### Internal tooling mode

Expected headers:

- `X-ADMIN-SECRET`
- `X-USER-EMAIL`

This mode exists for developers to test admin routes without generating `X-HMAC-SIGNATURE`.

### Shopper API

Shopper requests are protected by `shopper/verify_middleware.go`.

Expected header:

- `TABLE-TOKEN`

The server resolves the table by token, then injects business and table identity into request context.

## Route Overview

### Admin routes

Under `/api/v1`:

- `POST /signin`
- `GET /healthCheck/`
- `GET /menu_items/`
- `POST /menu_items/`
- `GET /orders/`
- `GET /orders/{table_id}`
- `GET /orders/{order_id}`
- `PATCH /orders/{order_id}/status`
- `POST /tables/`
- `GET /tables/`
- `GET /tables/{id}`
- `POST /tables/{id}/paid`
- `GET /notifications/`
- `PATCH /notifications/{id}`
- `GET /notifications/ws`

### Shopper routes

Under `/s/api/v1`:

- `POST /tables/validate`
- `GET /shops/{id}/menu_items`
- `POST /orders/`
- `GET /notifications/`

## Websocket Notes

Current websocket upgrade handlers exist at:

- `/api/v1/notifications/ws`
- `/s/api/v1/notifications/`

The websocket protocol in `notifications/client.go` currently supports:

- subscribing with `{"type":"subscribe","topic":"admin"}`
- publishing with `{"type":"message","topic":"...","content":"..."}`

Current known topic from application code:

- `admin`

Current known published event:

- new order notifications from `shopper/orders.go`

Example payload shape:

```json
{
  "code": "new_order",
  "message": "A new order has been created from table - #12",
  "order_id": "101",
  "table_id": "12"
}
```

## Happy Path Integration

This is the clearest current end-to-end flow for a shopper order.

1. Admin creates a table through `POST /api/v1/tables/`.
2. The server generates a table token and QR code URL.
3. Shopper validates the token through `POST /s/api/v1/tables/validate`.
4. Shopper sends `TABLE-TOKEN` on subsequent shopper requests.
5. Shopper creates an order through `POST /s/api/v1/orders/`.
6. The server stores the order.
7. The server stores a notification row.
8. The server publishes a websocket message on the `admin` topic.
9. Admin clients subscribed to `admin` can react to the new order event.

## Storage Notes

- The current implementation uploads generated table QR codes through `internal/utils/upload.go`.
- The code currently uses AWS SDK S3 configuration variables.
- The intended direction is Cloudflare R2, so future storage changes should prefer S3-compatible abstractions rather than hard-coding AWS-specific assumptions deeper into the app.

## Deployment Notes

The project has not been deployed yet.

Current documented environments:

- local only

There is currently no staging or production runbook to preserve.
