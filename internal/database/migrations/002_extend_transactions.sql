-- +goose Up
ALTER TABLE transactions
    ADD COLUMN description TEXT NOT NULL DEFAULT '',
    ADD COLUMN merchant TEXT,
    ADD COLUMN source TEXT NOT NULL DEFAULT 'manual',
    ADD COLUMN currency TEXT NOT NULL DEFAULT 'RSD';

ALTER TABLE transactions
    ADD CONSTRAINT chk_type CHECK (type IN ('income', 'expense')),
    ADD CONSTRAINT chk_source CHECK (source IN ('manual', 'import', 'notification', 'recurring')),
    ADD CONSTRAINT chk_currency CHECK (currency ~ '^[A-Z]{3}$');

-- +goose Down
ALTER TABLE transactions
    DROP CONSTRAINT chk_type,
    DROP CONSTRAINT chk_source,
    DROP CONSTRAINT chk_currency;

ALTER TABLE transactions
    DROP COLUMN description,
    DROP COLUMN merchant,
    DROP COLUMN source,
    DROP COLUMN currency;