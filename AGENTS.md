# AGENTS.md

This file is the quick-start guide for AI agents working in `table-tap.api`.

## Project Summary

TableTap is a Go backend for a restaurant self-ordering system.

- `server/` is the main executable and mounts both admin APIs and shopper APIs.
- `shopper/` contains customer-facing handlers and middleware, mounted under `/s`.
- `db/` contains SQL-backed data access methods grouped by domain.
- `notifications/` contains the in-memory websocket hub and client handling.
- `internal/types/` defines shared request/response and domain types.
- `internal/httphelper/` contains JSON read/write helpers.
- `internal/utils/` contains context accessors and shared utilities such as S3 upload.
- `migrations/` contains SQL migrations managed with `migrate`.

Current deployment status:

- local development only
- no staging environment yet
- no production deployment yet

## Runtime Shape

- The main binary is started from `server/main.go`.
- `server/main.go` loads env vars with `godotenv.Overload()`, opens Postgres, initializes the notification hub, and starts the HTTP server.
- `server/router.go` mounts:
  - shopper routes at `/s`
  - admin/API routes at `/api`
- Shopper routes are defined in `shopper/router.go`.
- Admin routes are defined in `server/router.go`.

## Auth And Request Context

- Admin/API requests use `server/verify_middleware.go`.
- Admin auth supports either:
  - `X-ADMIN-SECRET`
  - `X-HMAC-SIGNATURE` plus `X-USER-EMAIL`
- Shopper requests use `shopper/verify_middleware.go` and require `TABLE-TOKEN`.
- Business, user, and table identity are passed via `context.Context`, then read through helpers in `internal/utils/ctx.go`.

When changing handlers or DB access, preserve the current context-driven pattern instead of passing IDs around manually.

## Important Environment Variables

Defined in `server/.env.sample`:

- `TABLE_TAP_DB`
- `PORT`
- `ENCRYPTION_KEY`
- `SECRET_API_KEY`
- `TOKEN_SECRET_KEY`
- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `AWS_BUCKET_NAME`

Notes:

- `ENCRYPTION_KEY` is expected to be base64-encoded and decode to 32 bytes.
- The current code uses S3-compatible upload variables, but the intended storage direction is Cloudflare R2 rather than long-term AWS S3 usage.
- Websocket notifications are currently in-process only.
- Local startup expects a `.env` file in `server/`.

## Common Commands

From repo root:

```powershell
go test ./...
```

```powershell
go run ./server
```

```powershell
make migration_up
```

```powershell
make migration_down
```

If `make` is unavailable, use the `migrate` CLI directly against `./migrations`.

## Editing Guidance

- Keep handlers thin; business logic should stay in `db/`, `notifications/`, or shared helpers unless the project is already using a different pattern in that area.
- Follow the existing package split instead of introducing cross-package shortcuts.
- Reuse `readJSON`/`writeJSON` and existing response types from `internal/types`.
- Keep route registration inside the existing router files.
- Prefer extending the matching domain file, for example:
  - order queries in `db/orders.go`
  - table handlers in `server/tables.go` or `shopper/tables.go`
- Be careful with websocket changes because `NotificationHub` is shared across admin and shopper packages.

## Known Repo Quirks

- The repo currently tracks `server/.env`; do not overwrite secrets while working.
- `README.md` still contains user-facing setup guidance but not enough implementation detail for code agents.
- There are no visible test files today, so `go test ./...` is mainly a compile check unless tests are added later.
- The README mentions Redis, but Redis is not part of the current implementation and is not required for local development.
- `X-ADMIN-SECRET` exists as an internal developer/testing shortcut so engineers can call admin routes without generating an HMAC signature.

## More Detailed Docs

- Architecture reference: `docs/ARCHITECTURE.md`
- Development workflow: `docs/DEVELOPMENT.md`
- API and integration notes: `docs/API.md`
- Domain rules and state notes: `docs/DOMAIN.md`
- Concrete request examples: `docs/EXAMPLES.md`
