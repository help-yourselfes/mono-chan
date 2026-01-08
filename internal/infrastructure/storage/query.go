package storage

var Query = `
PRAGMA FOREIGN_KEYS = ON;
PRAGMA journal_mode=WAL;


CREATE TABLE IF NOT EXISTS boards (
	"key" TEXT PRIMARY KEY UNIQUE, 
	caption TEXT,
	description TEXT,
	last_post_id INTEGER
);

CREATE TABLE IF NOT EXISTS threads (
	global_id INTEGER PRIMARY KEY UNIQUE,
	board_key TEXT NOT NULL,
	post_id INTEGER NOT NULL,
	user_hash TEXT,
	password_hash TEXT,
	caption TEXT,
	reply_count INTEGER DEFAULT 0,
	FOREIGN KEY(board_key) REFERENCES boards("key") ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS posts (
	global_id INTEGER PRIMARY KEY UNIQUE,
	board_key TEXT NOT NULL,
	thread_id INTEGER NOT NULL,
	root_post_id INTEGER,		
	id INTEGER UNIQUE,
	text TEXT,
	media_json TEXT,
	password_hash TEXT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMP DEFAULT NULL,
	is_op INTEGER,
	FOREIGN KEY(board_key) REFERENCES boards("key"),
	FOREIGN KEY(thread_id) REFERENCES threads(global_id) ON DELETE CASCADE
);

CREATE TRIGGER IF NOT EXISTS update_post_timestamp 
AFTER UPDATE ON posts
BEGIN
	UPDATE posts SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;

CREATE UNIQUE INDEX IF NOT EXISTS idx_posts_board_local 
ON posts (board_key, id);

CREATE INDEX IF NOT EXISTS idx_posts_threads 
ON posts (root_post_id) 
WHERE root_post_id IS NOT NULL;
`
