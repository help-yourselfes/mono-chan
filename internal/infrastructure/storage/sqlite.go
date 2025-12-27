package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitSQLiteStorage(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		return nil, err
	}

	query := `
		CREATE TABLE IF NOT EXISTS boards (
			key TEXT PRIMARY KEY UNIQUE, 
			caption TEXT,
			description TEXT
		);
	`

	if _, err := db.Exec(query); err != nil {
		return nil, err
	}

	return db, err
}
