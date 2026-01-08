package repo

import (
	"context"
	"database/sql"

	"github.com/helpyourselfes/mono-chan/internal/app/thread/model"
	"github.com/helpyourselfes/mono-chan/internal/app/thread/repo"
	"github.com/helpyourselfes/mono-chan/internal/infrastructure/storage"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
	"github.com/mattn/go-sqlite3"
)

type sqliteThreadRepo struct {
	db *sql.DB
}
type t = sqliteThreadRepo

var _ repo.ThreadRepo = &t{}

func NewSQLiteThreadRepo(db *sql.DB) *sqliteThreadRepo {
	return &sqliteThreadRepo{db: db}
}

func (r sqliteThreadRepo) Create(ctx context.Context, thread *model.Thread) (int64, error) {
	ex := storage.GetExecutor(ctx, r.db)
	query := `
	INSERT INTO threads (
		board_key,
		post_id,
		user_hash,
		password_hash,	
		caption
	) VALUES (?,?,?,?,?)
	`

	res, err := ex.ExecContext(ctx, query,
		thread.BoardKey,
		thread.PostID,
		thread.UserHash,
		thread.PasswordHash,
		thread.Caption,
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
func (r sqliteThreadRepo) GetByGlobalID(ctx context.Context, globalID int64) (*model.Thread, error) {
	ex := storage.GetExecutor(ctx, r.db)
	query := `
	SELECT 
		global_id,
		board_key,
		post_id,
		user_hash,
		password_hash,
		caption,
		reply_count
	FROM threads
	WHERE global_id = ?
	`
	var thread model.Thread
	err := ex.QueryRowContext(ctx, query, globalID).Scan(
		&thread.GlobalID,
		&thread.BoardKey,
		&thread.PostID,
		&thread.UserHash,
		&thread.PasswordHash,
		&thread.Caption,
		&thread.ReplyCount,
	)
	if err == sql.ErrNoRows {
		return nil, customErrors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &thread, nil
}
func (r sqliteThreadRepo) GetByPostID(ctx context.Context, boardKey string, postID int64) (*model.Thread, error) {
	ex := storage.GetExecutor(ctx, r.db)
	query := `
	SELECT 
		global_id,
		board_key,
		post_id,
		user_hash,
		password_hash,
		caption,
		reply_count
	FROM threads
	WHERE board_key = ? AND post_id = ?
	`
	var thread model.Thread
	err := ex.QueryRowContext(ctx, query, boardKey, postID).Scan(
		&thread.GlobalID,
		&thread.BoardKey,
		&thread.PostID,
		&thread.UserHash,
		&thread.PasswordHash,
		&thread.Caption,
		&thread.ReplyCount,
	)
	if err == sql.ErrNoRows {
		return nil, customErrors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &thread, nil
}
func (r sqliteThreadRepo) Update(ctx context.Context, thread *model.Thread) error {
	ex := storage.GetExecutor(ctx, r.db)
	query := `
	UPDATE threads SET
	caption = ?
	WHERE global_id = ?
	`
	_, err := ex.ExecContext(ctx, query, thread.Caption, thread.GlobalID)
	return err
}
func (r sqliteThreadRepo) Delete(ctx context.Context, globalID int64) error {
	ex := storage.GetExecutor(ctx, r.db)
	query := `--sql
	DELETE FROM threads
	WHERE global_id = ?
	`

	_, err := ex.ExecContext(ctx, query, globalID)
	return err
}
func (r sqliteThreadRepo) List(ctx context.Context, boardKey string) ([]*model.Thread, error) {
	ex := storage.GetExecutor(ctx, r.db)
	query := `
	SELECT
		global_id,
		board_key,
		post_id,
		user_hash,
		password_hash,
		caption,
		reply_count
	FROM threads
	WHERE board_key = ?`

	rows, err := ex.QueryContext(ctx, query, boardKey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var threads = make([]*model.Thread, 0)
	for rows.Next() {
		var thread model.Thread
		if err := rows.Scan(
			&thread.GlobalID,
			&thread.BoardKey,
			&thread.PostID,
			&thread.UserHash,
			&thread.PasswordHash,
			&thread.Caption,
			thread.ReplyCount,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, customErrors.ErrNotFound
			}
			return nil, err
		}

		threads = append(threads, &thread)
	}

	return threads, nil
}
func (r sqliteThreadRepo) ListWithPost(ctx context.Context, boardKey string) ([]*model.ThreadPost, error) {
	return nil, customErrors.ErrNotImplemented
}
func (r sqliteThreadRepo) Reply(ctx context.Context, boardKey string, postID int64) error {
	ex := storage.GetExecutor(ctx, r.db)
	query := `
	UPDATE threads SET
	reply_count = reply_count + 1
	WHERE board_key = ? AND post_id = ?
	`
	_, err := ex.ExecContext(ctx, query, boardKey, postID)
	return err
}
