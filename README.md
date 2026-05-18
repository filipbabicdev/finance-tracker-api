# Finance Tracker API

A REST API for tracking personal financial transactions. Built in Go as a learning project to practice idiomatic backend patterns.

## Tech Stack

- **Language**: Go 1.22+
- **Web framework**: [Gin](https://github.com/gin-gonic/gin)
- **Database**: SQLite (migration to PostgreSQL in progress)
- **Authentication**: JWT (in progress)

## Features

- ✅ CRUD operations for financial transactions
- ✅ Persistent storage with SQLite
- ✅ Modular file structure (handlers / models / database)
- 🚧 JWT authentication
- 🚧 PostgreSQL support via Docker Compose
- 🚧 Integration tests with testify
- 🚧 Structured logging

## Getting Started

### Prerequisites

- Go 1.22 or higher
- SQLite (or use the included database file structure)

### Run locally

```bash
git clone https://github.com/filipbabicdev/finance-tracker-api.git
cd finance-tracker-api
go mod download
go run .
```

The server starts on `localhost:8080`.

### Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/transactions` | List all transactions |
| GET | `/transactions/:id` | Get a transaction by ID |
| POST | `/transactions` | Create a new transaction |
| PUT | `/transactions/:id` | Update a transaction |
| DELETE | `/transactions/:id` | Delete a transaction |

Example:

```bash
curl -X POST http://localhost:8080/transactions \
  -H "Content-Type: application/json" \
  -d '{"amount": 42.50, "category": "groceries", "note": "weekly shop"}'
```

## Roadmap

- [ ] Migrate storage layer to PostgreSQL
- [ ] Add JWT-based authentication and per-user transactions
- [ ] Containerize with Docker + docker-compose
- [ ] Add integration tests
- [ ] Add structured logging with `slog`
- [ ] Add request validation
- [ ] Deploy to Render with managed Postgres

## Author

Filip Babić — [GitHub](https://github.com/filipbabicdev) · [LinkedIn](https://linkedin.com/in/rooky)