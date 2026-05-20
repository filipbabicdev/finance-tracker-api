package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func New(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
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
		id       INTEGER PRIMARY KEY AUTOINCREMENT,
		amount   REAL    NOT NULL,
		category TEXT    NOT NULL,
		date     DATETIME NOT NULL,
		type     TEXT    NOT NULL
	)`)
	return err
}
