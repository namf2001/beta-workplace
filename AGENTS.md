# AGENTS.md

## Project Snapshot
- Go monolith API (`github.com/namf2001/beta-workplace`) using Gin + Postgres.
- Runtime wiring is in `cmd/server/main.go`: config -> DB pool -> OAuth init -> repository registry -> controllers -> HTTP handlers.
- Route registration lives in `cmd/server/router.go` and is the fastest source of truth for what is actually exposed.

## Architecture You Should Follow
- Keep the existing flow: **handler -> controller -> repository -> DB**.
- Handlers in `internal/handler/rest/v1/*` do HTTP concerns only (bind/validate/status code/response envelope).
- Controllers in `internal/controller/*` hold business logic and cross-repo orchestration.
- Repositories in `internal/repository/*` own SQL and persistence errors.
- Shared response envelope is `internal/handler/response.Response` via `response.NewResponse(code, message, data)`.

## Canonical Patterns (Use These as Templates)
- New user flow: `internal/handler/rest/v1/users/create_user.go` -> `internal/controller/users/create_user.go` -> `internal/repository/users/create_user.go`.
- Transactional multi-step auth flow uses `repo.DoInTx(...)` in `internal/controller/auth/register.go`.
- Middleware auth gate is `internal/handler/middleware/auth.go` and sets `c.Set("userID", claims.UserID)`.
- Repo constructors follow `New(db pg.ContextExecutor)` (example: `internal/repository/users/new.go`).

## Critical Caveats
- Do not trust feature docs as implementation truth. Current live routes are mostly auth + users from `cmd/server/router.go`.
- `Logout()` currently reads `user_claims` (`internal/handler/rest/v1/auth/logout.go`) but middleware sets `userID`; this mismatch is an existing bug risk.
- `docs/feature /API_DOCUMENTATION.md` includes many endpoints/entities not currently wired in router.

## Local Workflows
- Environment loading: `config.Init(env)` loads `.env` then `.env.<env>` from project root (`config/config.go`).
- Typical local run flow from `README.md` and `Makefile`:
```bash
cp .env.example .env.dev
docker-compose up -d
make migrate-up
make dev
```
- Helpful targets: `make test`, `make lint`, `make swagger`, `make run-dev`, `make run-product`.
- Hot reload uses `.air.toml`; excludes `vendor`, `tmp`, `testdata`, and `*_test.go`.

## Database + Migrations
- Migrations are plain SQL files applied in filename order by `make migrate-up`.
- Core tables start in `migrations/001_create_users_table.up.sql`, `003_auth_schema.up.sql`, `004_core_schema.up.sql`, `005_projects_tasks.up.sql`.
- Repository tests assume a real Postgres and wrap each test in rollback tx via `internal/pkg/testdb/testdb.go`.

## Testing Conventions
- Most existing tests are repository-level (`internal/repository/users/*_test.go`).
- Pattern: `testdb.WithTx(...)` + `testdb.LoadTestSQLFile(...)` + repository call + `require.*` assertions.
- If you add repository behavior, add/extend tests in the same package with SQL fixtures under `testdata/`.

## Integrations and External Dependencies
- Auth token: `internal/pkg/jwt/jwt.go` (HMAC JWT using `JWT_SECRET` + `JWT_ACCESS_DURATION`).
- Google OAuth: `internal/pkg/oauth/google.go` and auth handlers in `internal/handler/rest/v1/auth/google.go`.
- Mail OTP flow: `internal/pkg/mail/mail.go` with HTML templates under `internal/pkg/mail/templates/`.
- Observability/public ops endpoints already mounted: `/health`, `/metrics`, `/swagger/*any` (`cmd/server/router.go`).

## When Adding New Features
- Add route in `cmd/server/router.go` first, then implement handler/controller/repository chain.
- Reuse constants from `constants/message_constants.go` for API `code`/`message` responses.
- Keep controller logic context-first (`func(ctx context.Context, ...)`) and pass request context down to repo calls.
- Prefer extending existing repository interfaces and registry (`internal/repository/registry.go`) over direct DB access from controllers.

