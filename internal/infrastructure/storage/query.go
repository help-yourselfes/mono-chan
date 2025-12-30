package storage

var Query = `
PRAGMA FOREIGN_KEYS = ON;


CREATE TABLE IF NOT EXISTS boards (
	key TEXT PRIMARY KEY UNIQUE, 
	caption TEXT,
	description TEXT,
	last_post_id INTEGER
);

CREATE TABLE IF NOT EXISTS threads (
	global_id INTEGER PRIMARY KEY UNIQUE,
	post_id INTEGER UNIQUE,
	FOREIGN KEY(board_key) REFERENCES boards(key) ON DELETE CASCADE,
	caption TEXT,
	reply_count INTEGER
)

CREATE TABLE IF NOT EXISTS posts (
	global_id INTEGER PRIMARY KEY UNIQUE,
	FOREIGN KEY(board_key) REFERENCES boards(key),
	FOREIGN KEY(thread_id) REFERENCES threads(id) ON DELETE CASCADE,
	id INTEGER UNIQUE,
	text TEXT,
	media_json TEXT,
	password_hash TEXT,
	created_at TEXT DEFAULT CURRENT_TIMESTAMP,
	updated_at TEXT DEFAULT NULL
)

CREATE TRIGGER IF NOT EXISTS update_post_timestamp 
AFTER UPDATE ON posts
BEGIN
	UPDATE posts SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
`
