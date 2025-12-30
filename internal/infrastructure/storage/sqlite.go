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

	if _, err := db.Exec(Query); err != nil {
		return nil, err
	}

	return db, err
}
