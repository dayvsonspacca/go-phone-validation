package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const DB_URL = "db.sqlite"

func Initialize() {
	db := GetConnection()
	createSchema(db)
}

func createSchema(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS valitaded_phones (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		status TEXT CHECK(status IN ('MATCHED', 'UNMATCHED')) NOT NULL,
		phone TEXT,
		national_identy_number TEXT
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create tables schema: %v", err)
	}
}

func GetConnection() *sql.DB {
	db, err := sql.Open("sqlite3", DB_URL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}
