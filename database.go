package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./finance.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		amount REAL NOT NULL,
		category TEXT NOT NULL,
		date DATETIME NOT NULL,
		type TEXT NOT NULL
	);`
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}
