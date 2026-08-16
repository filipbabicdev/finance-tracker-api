-- +goose Up
CREATE TABLE IF NOT EXISTS categories (
    id         BIGSERIAL   PRIMARY KEY,
    name       TEXT        NOT NULL UNIQUE,
    type       TEXT        NOT NULL,
    bucket     TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT chk_category_type   CHECK (type IN ('income', 'expense')),
    CONSTRAINT chk_category_bucket CHECK (bucket IS NULL OR bucket IN ('needs', 'wants', 'savings')),
    CONSTRAINT chk_income_no_bucket CHECK (
        (type = 'income'  AND bucket IS NULL) OR
        (type = 'expense' AND bucket IS NOT NULL)
    )
);

INSERT INTO categories (name, type, bucket) VALUES
    ('After-tax salary or wages', 'income', NULL),
    ('Any additional income (rental, financial aid, self-employment, child support, pension, etc.)', 'income', NULL),

    ('Rent/mortgage',                              'expense', 'needs'),
    ('Infostan',                                   'expense', 'needs'),
    ('Electricity and natural gas bill',           'expense', 'needs'),
    ('Sanitation/garbage bill',                    'expense', 'needs'),
    ('Phone bill',                                 'expense', 'needs'),
    ('Internet + TV bill',                         'expense', 'needs'),
    ('Groceries, toiletries and other essentials', 'expense', 'needs'),
    ('Out-of-pocket medical costs',                'expense', 'needs'),
    ('Minimum student loan payments',              'expense', 'needs'),
    ('Car payment',                                'expense', 'needs'),
    ('Parking and registration fees',              'expense', 'needs'),
    ('Car maintenance and repairs',                'expense', 'needs'),
    ('Gasoline',                                   'expense', 'needs'),
    ('Public transportation',                      'expense', 'needs'),
    ('Pets',                                       'expense', 'needs'),
    ('Sport / trening',                            'expense', 'needs'),
    ('Other needs',                                'expense', 'needs'),

    ('Clothing, jewelry, etc.',                                      'expense', 'wants'),
    ('Bars, Clubs and Dining out',                                   'expense', 'wants'),
    ('Take-out, food delivery',                                      'expense', 'wants'),
    ('Sports and recreation',                                        'expense', 'wants'),
    ('Gifts and celebrations',                                       'expense', 'wants'),
    ('Movie, concert and event tickets',                             'expense', 'wants'),
    ('Travel expenses (airline tickets, hotels, rental cars, etc.)', 'expense', 'wants'),
    ('Subscriptions',                                                'expense', 'wants'),
    ('Home and equipment',                                           'expense', 'wants'),
    ('Electronics',                                                  'expense', 'wants'),
    ('Other wants',                                                  'expense', 'wants'),

    ('Emergency fund contributions',                'expense', 'savings'),
    ('Savings account contributions',               'expense', 'savings'),
    ('Individual retirement account contributions', 'expense', 'savings'),
    ('Travel sinking fund',                         'expense', 'savings'),
    ('Debt repayment',                              'expense', 'savings'),
    ('Excess payments on student loans',            'expense', 'savings'),
    ('Excess payments on mortgage',                 'expense', 'savings')
ON CONFLICT (name) DO NOTHING;

ALTER TABLE transactions
    ADD COLUMN category_id BIGINT REFERENCES categories(id) ON DELETE RESTRICT;

CREATE INDEX idx_transactions_category_id ON transactions (category_id);

-- +goose Down
DROP INDEX IF EXISTS idx_transactions_category_id;

ALTER TABLE transactions
    DROP COLUMN category_id;

DROP TABLE IF EXISTS categories;