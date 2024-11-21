package database

import (
	"database/sql"
	"log"
)

func NewSqliteDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	_, err = db.Exec(createTables)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	return db
}
