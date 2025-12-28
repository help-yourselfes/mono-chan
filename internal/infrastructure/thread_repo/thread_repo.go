package thread_repo

import (
	"context"
	"database/sql"
	"encoding/json"

	m "github.com/helpyourselfes/mono-chan/internal/app/thread/model"
	"github.com/helpyourselfes/mono-chan/internal/app/thread/repo"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
	"github.com/mattn/go-sqlite3"
)

type sqliteBoardRepo struct {
	db *sql.DB
}
type t = sqliteBoardRepo

var _ repo.ThreadRepo = &t{}

func NewSQLiteThreadRepo(db *sql.DB) *t {
	return &t{db: db}
}

/*

	internal/infrastructure/thread_repo/get_by_id.go

	internal/infrastructure/thread_repo/update.go

*/

func (r *t) Create(ctx context.Context, thread *m.Thread) (int64, error) {
	boardKey := thread.BoardKey
	caption := thread.Caption
	text := thread.Text
	mediaJSON, err := json.Marshal(thread.MediaLinks)
	password := thread.Password

	if err != nil {
		return -1, nil
	}
	mediaLinks := string(mediaJSON)

	query := `
	INSERT INTO threads (
		board_key, caption, text, media_links, password
	) VALUES (?,?,?,?,?)
	`

	res, err := r.db.ExecContext(ctx, query,
		boardKey, caption, text, mediaLinks, password,
	)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return -1, customErrors.ErrAlreadyExists
		}
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *t) Delete(ctx context.Context, id int64, password string) error {
	query := `
		DELETE FROM threads
		WHERE id = ?
	`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *t) Reply(ctx context.Context, id int64) error {

	query := `
	UPDATE threads
	SET reply_count = reply_count + 1
	WHERE id = ?
	`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return nil
	}

	return nil
}
