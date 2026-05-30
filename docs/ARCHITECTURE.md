# Architecture

## Overview

This repository is a Go backend for TableTap, a restaurant self-ordering flow with:

- admin/dashboard APIs
- shopper/customer APIs
- Postgres-backed persistence
- in-process websocket notifications

The application is organized as a small multi-package codebase rather than a single large `main` package.

Current environment status:

- local development exists
- staging does not exist yet
- production does not exist yet

## Entry Point

The executable starts in `server/main.go`.

Startup flow:

1. Load environment variables with `godotenv.Overload()`.
2. Read `TABLE_TAP_DB` and open Postgres with `sqlx` + `pgx`.
3. Decode `ENCRYPTION_KEY` from base64.
4. Initialize the shared notification hub.
5. Assign shared globals into the `shopper` package.
6. Build the root router.
7. Start `http.ListenAndServe`.

## Routing Topology

`server/router.go` defines the root router:

- `/s` -> shopper-facing routes from `shopper.GetRouter()`
- `/api` -> admin/API routes from `getApiRouter()`

### Admin Routes

Mounted under `/api/v1`, with `verify` middleware applied.

Representative areas:

- `/signin`
- `/healthCheck`
- `/menu_items`
- `/orders`
- `/tables`
- `/notifications`

These handlers live in `server/*.go`.

### Shopper Routes

Mounted under `/s/api/v1` for most shopper endpoints.

Representative areas:

- `/tables/validate`
- `/shops/{id}/menu_items`
- `/orders`
- `/notifications`

These handlers live in `shopper/*.go`.

## Package Responsibilities

### `server/`

- main program bootstrap
- admin-facing HTTP handlers
- admin auth and permission flow
- root router composition

### `shopper/`

- shopper-facing HTTP handlers
- table-token verification
- routes mounted into the main service

### `db/`

- SQL access layer
- domain-specific query files such as `orders.go`, `tables.go`, `menu_items.go`
- lightweight transaction and named-query helpers in `db.go`

The codebase generally expects handlers to call `DBConn` methods rather than issuing raw SQL inline.

### `notifications/`

- websocket connection upgrade
- client registration and subscription
- in-memory pub/sub hub

`NotificationHub` is process-local. It is not Redis-backed in the current implementation, even though the README mentions Redis for pub/sub setup.

### `internal/types/`

- domain models
- shared response payloads
- HTTP-facing structs
- custom types such as ISO 8601 helpers

### `internal/httphelper/`

- JSON decoding with `DisallowUnknownFields()`
- JSON response writing

### `internal/utils/`

- context value accessors
- shared helpers such as S3 upload

## Authentication Model

### Admin/API

Implemented in `server/verify_middleware.go`.

Two modes are supported:

- `X-ADMIN-SECRET`
- `X-HMAC-SIGNATURE`

For HMAC requests, the middleware computes a signature from the request URL and `SECRET_API_KEY`. It also requires `X-USER-EMAIL`, then looks up the current business user and injects these context values:

- user email
- user ID
- business ID
- permissions

### Shopper/API

Implemented in `shopper/verify_middleware.go`.

- Requires `TABLE-TOKEN`
- Resolves the table through the DB layer
- Injects business ID, table ID, and table number into context

## Notification Flow

Order creation in `shopper/orders.go` is the best current example:

1. Create the order in Postgres.
2. Insert a notification row in the database.
3. Build a JSON payload.
4. Publish the payload to the in-memory hub on the `admin` topic.

Because the hub is in-memory, notification delivery is tied to the running process instance.

Current websocket protocol:

- clients send JSON messages
- subscribe shape: `{"type":"subscribe","topic":"admin"}`
- publish shape: `{"type":"message","topic":"...","content":"..."}`
- the only concrete publish topic visible in application code today is `admin`

## Persistence And Migrations

- Postgres connection string comes from `TABLE_TAP_DB`.
- SQL migrations live in `migrations/`.
- The `Makefile` wraps `migrate` commands for up/down/fix/create.
- Redis is not required in the current implementation.
- Storage is currently wired through S3 environment variables, but the intended direction is Cloudflare R2.

When schema changes are needed:

1. add a migration
2. update `internal/types` if necessary
3. update the corresponding `db/` file
4. update handlers if request/response payloads changed

## Design Conventions

- Handlers are intentionally thin.
- Shared state is passed through package globals set at startup:
  - `DBConn`
  - `NotificationHub`
- Request-scoped business and identity data travels through `context.Context`.
- Response payloads usually wrap data in typed success/error envelopes from `internal/types`.

## Risk Areas

- `server/.env` exists in the repo; treat it as sensitive local state.
- Shared globals mean package initialization order matters.
- Notification behavior is easy to break if topic names or hub wiring change.
- Route prefixes are easy to misread because shopper endpoints are mounted under `/s`.
