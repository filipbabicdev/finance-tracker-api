# ---- Build stage ----
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Static binary: CGO off (pure-Go pgx driver), stripped for smaller size.
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w" \
    -o server \
    ./cmd/server

# ---- Run stage ----
FROM alpine:3.20

WORKDIR /app

# Required for TLS connections to managed Postgres (Neon sslmode=require).
RUN apk add --no-cache ca-certificates

# Non-root runtime user.
RUN adduser -D -u 10001 appuser
USER appuser

COPY --from=builder /app/server .

EXPOSE 8090

CMD ["./server"]