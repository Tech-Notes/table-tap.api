# Examples

## Notes

These examples are based on the current server code and are meant as practical starting points for integrations, local testing, and future AI agents.

Replace:

- `https://example.com` with your actual host
- placeholder IDs and tokens with real values
- placeholder secrets with values from your local environment

## Admin HMAC Signing

Current signing behavior in `server/verify_middleware.go`:

1. build a string in the form `URL:<request-url>`
2. compute HMAC-SHA256 using `SECRET_API_KEY`
3. hex-encode the result
4. send it as `X-HMAC-SIGNATURE`

Special case:

- for notification URLs, the current code signs `URL:<path>` instead of the full request URL

Pseudo-code:

```text
data = "URL:" + request_url
if path contains "notifications":
  data = "URL:" + request_path

signature = hex(HMAC_SHA256(secret_api_key, data))
```

## Admin Request Example

Get menu items for the authenticated business:

```bash
curl --request GET "https://example.com/api/v1/menu_items/" \
  --header "X-USER-EMAIL: dev@example.com" \
  --header "X-HMAC-SIGNATURE: <hex-signature>"
```

Developer shortcut using internal tooling auth:

```bash
curl --request GET "https://example.com/api/v1/menu_items/" \
  --header "X-USER-EMAIL: dev@example.com" \
  --header "X-ADMIN-SECRET: <internal-secret>"
```

## Shopper Request Example

Validate a table token:

```bash
curl --request POST "https://example.com/s/api/v1/tables/validate" \
  --header "Content-Type: application/json" \
  --data "{\"tableToken\":\"<table-token>\"}"
```

Expected success shape:

```json
{
  "valid": true,
  "shopId": 1,
  "tableId": 12,
  "tableToken": "f5f5f9e4-0000-0000-0000-111111111111"
}
```

Possible invalid-token shape:

```json
{
  "valid": false
}
```

## Create Table

Admin creates a table:

```bash
curl --request POST "https://example.com/api/v1/tables/" \
  --header "Content-Type: application/json" \
  --header "X-USER-EMAIL: dev@example.com" \
  --header "X-HMAC-SIGNATURE: <hex-signature>" \
  --data "{\"table_no\":12,\"description\":\"Window seat\"}"
```

Current behavior:

- generates a UUID token
- generates a QR PNG
- uploads the QR code through the storage helper
- stores the new table in Postgres

Expected response shape:

```json
{
  "success": true,
  "message": "success",
  "data": {
    "id": 12
  }
}
```

## Create Shopper Order

Create an order for a validated table:

```bash
curl --request POST "https://example.com/s/api/v1/orders/" \
  --header "TABLE-TOKEN: <table-token>"
```

Current behavior:

- resolves the table from `TABLE-TOKEN`
- creates an order with status `pending`
- stores a `new_order` notification row
- publishes a websocket message on topic `admin`

Expected response shape:

```json
{
  "success": true,
  "message": "success",
  "data": {
    "id": 101
  }
}
```

## Update Order Status

Advance an order to `preparing`:

```bash
curl --request PATCH "https://example.com/api/v1/orders/101/status" \
  --header "Content-Type: application/json" \
  --header "X-USER-EMAIL: dev@example.com" \
  --header "X-HMAC-SIGNATURE: <hex-signature>" \
  --data "{\"status\":\"preparing\"}"
```

Important current rule:

- `paid` is not accepted by this endpoint

## Mark Table As Paid

Complete checkout for a table:

```bash
curl --request POST "https://example.com/api/v1/tables/12/paid" \
  --header "X-USER-EMAIL: dev@example.com" \
  --header "X-HMAC-SIGNATURE: <hex-signature>"
```

Current behavior:

- sets the table status to `available`
- marks all non-paid orders for the table as `paid`

## Websocket Example

Connect to the admin websocket endpoint:

```text
wss://example.com/api/v1/notifications/ws
```

After connecting, subscribe to the current known topic:

```json
{
  "type": "subscribe",
  "topic": "admin"
}
```

Current known new-order event shape:

```json
{
  "code": "new_order",
  "message": "A new order has been created from table - #12",
  "order_id": "101",
  "table_id": "12"
}
```

## End-To-End Happy Path

1. Admin creates a table.
2. Shopper validates the table token.
3. Shopper sends `TABLE-TOKEN` and creates an order.
4. Admin receives a stored notification and, if subscribed, a websocket event on topic `admin`.
5. Admin updates the order through `pending -> preparing -> ready`.
6. Admin marks the table as paid, which resets the table to `available` and marks outstanding orders as `paid`.
