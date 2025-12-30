package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/helpyourselfes/mono-chan/internal/app/post/dto"
	"github.com/helpyourselfes/mono-chan/internal/app/post/model"
	"github.com/helpyourselfes/mono-chan/internal/app/post/repo"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
	"github.com/mattn/go-sqlite3"
)

type sqlitePostRepo struct {
	db *sql.DB
}
type p = sqlitePostRepo

var _ repo.PostRepo = &p{}

func NewSQLitePostRepo(db *sql.DB) *t {
	return &t{db: db}
}

func (r *p) Create(ctx context.Context, post *model.Post) (int64, error) {
	mediaJSON, err := json.Marshal(post.MediaLinks)
	if err != nil {
		return -1, err
	}
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	if err != nil {
		return -1, err
	}

	defer tx.Rollback()

	var localId int64
	queryUpdate := `
	UPDATE boards GET last_post_id = last_post_id + 1
	WHERE key = ? RETURNING last_post_id
	`
	err = tx.QueryRowContext(ctx, queryUpdate, post.BoardKey).Scan(&localId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, customErrors.ErrNotFound
		}
		return -1, err
	}

	queryInsert := `
	INSERT INTO posts (
		board_key, thread_id, id, text, media_json, password_hash, created_at, updated_at, is_op
	) VALUES (?,?,?,?,?,?,?,?,?)
	`

	_, err = tx.ExecContext(ctx, queryInsert,
		post.BoardKey,
		post.ThreadID,
		localId,
		post.Text,
		mediaJSON,
		post.PasswordHash,
		post.CreatedAt,
		*post.UpdatedAt,
		post.IsOP,
	)

	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return -1, customErrors.ErrAlreadyExists
		}
		return -1, err
	}

	if err := tx.Commit(); err != nil {
		return -1, nil
	}

	return localId, nil
}
func (r *p) GetById(ctx context.Context, boardKey string, id int64) (*model.Post, error) {
	query := `
	SELECT
		global_id, board_key, thread_id, id, text, media_json,
		password_hash, created_at, updated_at, is_op
	FROM posts
	WHERE board_key = ? AND id = ?`

	var post model.Post
	var mediaJSON sql.NullString

	err := r.db.QueryRowContext(ctx, query, boardKey, id).Scan(
		&post.GlobalID,
		&post.BoardKey,
		&post.ThreadID,
		&post.ID,
		&post.Text,
		&mediaJSON,
		&post.PasswordHash,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.IsOP,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customErrors.ErrNotFound
		}
		return nil, err
	}

	post.MediaLinks, err = unJsonMedia(mediaJSON.String)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
func (r *p) Update(ctx context.Context, post *dto.UpdatePostRequest) error {
	mediaJSON, err := jsonMedia(post.MediaLinks)
	if err != nil {
		return err
	}
	query := `
	UPDATE posts SET
	text = ?,
	media_json = ?	
	WHERE board_key = ? AND id = ? 
	`
	_, err = r.db.ExecContext(ctx, query, post.Text, mediaJSON, post.BoardKey, post.ID)
	return err
}
func (r *p) Delete(ctx context.Context, globalId int64) error {
	query := `--sql
	DELETE FROM posts
	WHERE global_id = ?`
	_, err := r.db.ExecContext(ctx, query, globalId)
	return err
}
func (r *p) List(ctx context.Context, boardKey string, threadId int64) ([]*model.Post, error) {
	query := `
	SELECT
		global_id, board_key, thread_id, id, text, media_json,
		password_hash, created_at, updated_at, is_op
	FROM posts
	WHERE board_key = ? AND thread_id = ?`

	var mediaJSON sql.NullString

	rows, err := r.db.QueryContext(ctx, query, boardKey, threadId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts = make([]*model.Post, 0)
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(
			&post.GlobalID,
			&post.BoardKey,
			&post.ThreadID,
			&post.ID,
			&post.Text,
			&mediaJSON,
			&post.PasswordHash,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.IsOP,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, customErrors.ErrNotFound
			}
			return nil, err
		}
		post.MediaLinks, err = unJsonMedia(mediaJSON.String)
		if err != nil {
			return nil, err
		}

		posts = append(posts, &post)
	}

	return posts, nil
}
