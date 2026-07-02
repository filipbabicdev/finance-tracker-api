-- +goose Up
CREATE TABLE IF NOT EXISTS transactions (
    id       BIGSERIAL       PRIMARY KEY,
    amount   NUMERIC(12,2)   NOT NULL,
    category TEXT            NOT NULL,
    date     TIMESTAMPTZ     NOT NULL,
    type     TEXT            NOT NULL
);
-- +goose Down
DROP TABLE IF EXISTS transactions;