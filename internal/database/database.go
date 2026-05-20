package database

import (
	"database/sql"

	_"github.com/jackc/pgx/v5/stdlib"
)

func New(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if err = createTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS transactions (
		id       BIGSERIAL 		PRIMARY KEY,
		amount   NUMERIC(12, 2) NOT NULL,
		category TEXT    		NOT NULL,
		date     TIMESTAMPTZ 	NOT NULL,
		type     TEXT    		NOT NULL
	)`)
	return err
}
