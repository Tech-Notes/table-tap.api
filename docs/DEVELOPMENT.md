# Development

## Prerequisites

You will typically need:

- Go `1.23.5` or compatible
- PostgreSQL
- `migrate` CLI
- a populated `server/.env`

Optional, depending on the area being exercised:

- S3-compatible object storage credentials
- a client that can open websocket connections

## Local Setup

1. Create or update `server/.env` from `server/.env.sample`.
2. Set `TABLE_TAP_DB` to a working Postgres connection string.
3. Set `ENCRYPTION_KEY` to a base64 value that decodes to 32 bytes.
4. Fill in auth and AWS variables if your change touches those paths.

## Useful Commands

Run from repository root unless noted otherwise.

Start the API:

```powershell
go run ./server
```

Alternative from the `server/` directory:

```powershell
go run .
```

Compile-check everything:

```powershell
go test ./...
```

Apply migrations:

```powershell
make migration_up
```

Roll back the last migration batch:

```powershell
make migration_down
```

Fix migration version manually:

```powershell
make migration_fix VERSION=<version>
```

Create a new migration:

```powershell
make create_migration SEQ_NAME=<name>
```

## Where To Change Things

Add or modify admin endpoints:

- router registration in `server/router.go`
- handler code in the relevant `server/*.go` file
- DB access in the matching `db/*.go` file

Add or modify shopper endpoints:

- router registration in `shopper/router.go`
- handler code in the relevant `shopper/*.go` file
- DB access in the matching `db/*.go` file

Add fields to payloads or responses:

- update `internal/types/`
- update handlers that read/write those payloads
- update SQL scanning/select lists if the field is persisted

Change auth behavior:

- admin auth lives in `server/verify_middleware.go` and `server/auth.go`
- shopper auth lives in `shopper/verify_middleware.go`

Change notifications:

- websocket plumbing is in `notifications/`
- notification persistence is in `db/notifications.go`
- publishing examples exist in `shopper/orders.go`

## Working Conventions

- Prefer `readJSON` and `writeJSON` over custom JSON handling.
- Preserve `DisallowUnknownFields()` behavior unless there is a deliberate API compatibility decision.
- Keep SQL in the `db/` package rather than embedding it in handlers.
- Use context helpers from `internal/utils/ctx.go` to read business, user, and table IDs.
- Match the existing response envelope style from `internal/types`.

## Validation Expectations

Before finishing a change, prefer:

1. `go test ./...`
2. targeted manual route verification for the area you changed
3. migration validation if you touched `migrations/`

There are currently no visible `_test.go` files, so compile success is the main automated safety net unless new tests are added.

## Agent Notes

- Check for uncommitted user work before editing.
- Avoid touching `server/.env` unless the task is explicitly about local configuration.
- Be careful with route prefixes. Shopper endpoints are not rooted directly at `/api`.
- If a change affects both admin and shopper flows, verify the shared `DBConn` or `NotificationHub` assumptions still hold.
- Do not add Redis-dependent setup unless the task explicitly introduces Redis; it is not part of the current system.
- If touching file upload/storage code, document whether the implementation still targets AWS S3 directly or has moved to Cloudflare R2 compatibility.
