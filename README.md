# Finance Tracker API

A REST API for tracking personal financial transactions. Built in Go as a portfolio project to demonstrate idiomatic backend patterns (repository pattern, explicit dependency injection, proper money handling).

## Tech Stack

- **Language**: Go 1.25
- **Web framework**: [Gin](https://github.com/gin-gonic/gin)
- **Database**: PostgreSQL 16 (via [pgx/v5](https://github.com/jackc/pgx))
- **Monetary types**: [shopspring/decimal](https://github.com/shopspring/decimal)
- **Infrastructure**: Docker + docker-compose

## Features

- ✅ CRUD operations for financial transactions
- ✅ PostgreSQL 16 with connection pooling
- ✅ Repository pattern (handler layer never touches `*sql.DB`)
- ✅ Exact-precision money with `NUMERIC(12,2)` + `shopspring/decimal`
- ✅ Docker-based local dev (single `docker compose up`)
- 🚧 JWT authentication
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

The server starts on `localhost:8090`.

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

## Roadmap

- [ ] JWT authentication (next)
- [ ] Integration tests with [testcontainers-go](https://github.com/testcontainers/testcontainers-go)
- [ ] gRPC service layer
- [ ] Deploy to Render or Fly.io with managed Postgres

## Author

Filip Babić — [GitHub](https://github.com/filipbabicdev) · [LinkedIn](https://linkedin.com/in/rooky)