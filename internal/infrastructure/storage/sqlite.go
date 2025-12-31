package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitSQLiteStorage(dbPath string) (*sql.DB, error) {
	connectionString := dbPath + "?parseTime=true"
	db, err := sql.Open("sqlite3", connectionString)

	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(Query); err != nil {
		return nil, err
	}

	return db, err
}
