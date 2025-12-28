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

	db.Exec(`
		PRAGMA foreign_keys = ON;
		PRAGMA journal_mode=WAL;
	`)

	query := `
		CREATE TABLE IF NOT EXISTS boards (
			key TEXT PRIMARY KEY UNIQUE, 
			caption TEXT,
			description TEXT
		);

		CREATE TABLE IF NOT EXISTS threads (
			id INTEGER PRIMARY KEY UNIQUE,
		    board_key TEXT NOT NULL REFERENCES boards(key) ON DELETE CASCADE,
			caption TEXT,
			text TEXT,
			media_links TEXT,
			reply_count INTEGER NOT NULL DEFAULT 0,
			password TEXT,
			created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at text

		);
		
		CREATE INDEX IF NOT EXISTS idx_threads_board_key ON threads(board_key);
		CREATE INDEX IF NOT EXISTS idx_threads_board_key_updated_at ON threads(board_key, updated_at DESC)
	`

	if _, err := db.Exec(query); err != nil {
		return nil, err
	}

	return db, err
}
