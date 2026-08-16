# Finance Tracker API

A REST API for tracking personal financial transactions. Built in Go as a portfolio project to demonstrate idiomatic backend patterns (repository pattern, explicit dependency injection, proper money handling).

**Live demo:** https://finance-tracker-api-jsvw.onrender.com/transactions

> Deployed on Render (Docker) with a managed PostgreSQL database on Neon.
> Hosted on a free tier, so the first request after idle may take 30–50s to wake (cold start).

## Tech Stack

- **Language**: Go 1.25
- **Web framework**: [Gin](https://github.com/gin-gonic/gin)
- **Database**: PostgreSQL 16 (via [pgx/v5](https://github.com/jackc/pgx))
- **Migrations**: [goose](https://github.com/pressly/goose) (embedded, applied at startup)
- **Monetary types**: [shopspring/decimal](https://github.com/shopspring/decimal)
- **Infrastructure**: Docker + docker-compose; deployed on Render with managed Postgres (Neon)

## Features

- ✅ CRUD operations for financial transactions
- ✅ PostgreSQL 16 with connection pooling
- ✅ Repository pattern (handler layer never touches `*sql.DB`)
- ✅ Exact-precision money with `NUMERIC(12,2)` + `shopspring/decimal`
- ✅ Embedded goose migrations, applied automatically on startup
- ✅ Request validation at both the DTO and database layers
- ✅ Category domain with 50/30/20 budget buckets (`needs` / `wants` / `savings`)
- ✅ Docker-based local dev (single `docker compose up`)
- 🚧 Category HTTP endpoints
- 🚧 Integration tests with testcontainers-go
- 🚧 Structured logging

## Getting Started

### Prerequisites

- Docker + docker-compose

### Run locally

```bash
git clone https://github.com/filipbabicdev/finance-tracker-api.git
cd finance-tracker-api
cp .env.example .env
docker compose up --build
```

The server starts on `localhost:8090`. Migrations run automatically on startup.

### Endpoints

| Method | Endpoint | Request body | Response |
|--------|----------|--------------|----------|
| `GET` | `/transactions` | — | Array of transactions |
| `POST` | `/transactions` | `{"amount": "49.99", "category": "groceries", "date": "2026-05-20T00:00:00Z", "type": "expense"}` | Created transaction |
| `PUT` | `/transactions/:id` | `{"amount": "55.00", "category": "groceries", "date": "2026-05-20T00:00:00Z", "type": "expense"}` | Updated transaction |
| `DELETE` | `/transactions/:id` | — | `204 No Content` |

Example:

```bash
curl -X POST http://localhost:8090/transactions \
  -H "Content-Type: application/json" \
  -d '{"amount": "49.99", "category": "groceries", "date": "2026-05-20T00:00:00Z", "type": "expense"}'
```

## Design Decisions

**Repository pattern** — handlers never touch `*sql.DB` directly. Swapping the storage backend only requires a new struct satisfying the same interface, not a handler rewrite.

**`NUMERIC(12,2)` for amounts** — `FLOAT`/`REAL` use binary floating-point and can't represent decimal fractions exactly (`0.10` may round-trip as `0.09999...`). `NUMERIC` is exact-precision; `shopspring/decimal` carries that guarantee into Go.

**Connection pool limits** — capped via `DB_MAX_OPEN_CONNS` (default 25) and `DB_MAX_IDLE_CONNS` (default 5). Without a cap, a traffic spike can exhaust PostgreSQL's `max_connections`. The idle limit prevents holding unused connections while still amortizing connection setup cost.

**Two-step migration on a live database** — introducing the `categories` table added `transactions.category_id` as a **nullable** foreign key rather than `NOT NULL`. Migrations run at application startup, so a migration that fails against existing rows doesn't just fail — the service never comes up. Nullable first means existing rows migrate untouched and production keeps serving; the backfill and the `NOT NULL` constraint follow as a separate, independently verifiable step.

**Two dimensions, two columns** — a category carries both a direction (`type`: `income` / `expense`) and a budgeting purpose (`bucket`: `needs` / `wants` / `savings`). These are independent, so they are separate columns rather than one overloaded enum. A cross-column `CHECK` enforces the relationship between them: income categories must have no bucket, expense categories must have one.

**Sentinel errors at the repository boundary** — `GetByID` translates `sql.ErrNoRows` into `ErrCategoryNotFound`. Without it, the HTTP layer would have to import `database/sql` to recognise a 404, leaking the persistence layer into transport code.

## Roadmap

- [x] Deploy to Render with managed Postgres (Neon)
- [x] Embedded goose migrations (replacing hand-rolled table creation)
- [x] Input validation: constrain `type` to `income`/`expense` (DB `CHECK` + app-layer guard)
- [x] Category domain: table, seed data, repository layer
- [ ] Integration tests with [testcontainers-go](https://github.com/testcontainers/testcontainers-go)
- [ ] Category HTTP endpoints (`GET /categories`)
- [ ] Link transactions to categories: backfill `category_id`, then enforce `NOT NULL`
- [ ] CSV / spreadsheet import
- [ ] Structured logging
- [ ] gRPC service layer

## Author

Filip Babić — [GitHub](https://github.com/filipbabicdev) · [LinkedIn](https://linkedin.com/in/rooky)