package repo

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/helpyourselfes/mono-chan/internal/app/post/dto"
	"github.com/helpyourselfes/mono-chan/internal/app/post/model"
	"github.com/helpyourselfes/mono-chan/internal/app/post/repo"
	"github.com/helpyourselfes/mono-chan/internal/infrastructure/storage"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
	"github.com/mattn/go-sqlite3"
)

type sqlitePostRepo struct {
	db *sql.DB
}
type p = sqlitePostRepo

var _ repo.PostRepo = &p{}

func NewSQLitePostRepo(db *sql.DB) *sqlitePostRepo {
	return &sqlitePostRepo{db: db}
}

func (r *p) Create(ctx context.Context, post *model.Post) (int64, error) {
	mediaJSON, err := json.Marshal(post.MediaLinks)
	if err != nil {
		return -1, err
	}

	ex := storage.GetExecutor(ctx, r.db)

	query := `
	INSERT INTO posts (
		board_key,
		thread_id,
		root_post_id,
		id,
		text,
		media_json,
		password_hash,
		created_at,
		updated_at,
		is_op
	) VALUES (?,?,?,?,?,?,?,?,?,?)
	`

	_, err = ex.ExecContext(ctx, query,
		post.BoardKey,
		post.ThreadID,
		post.RootPostID,
		post.ID,
		post.Text,
		mediaJSON,
		post.PasswordHash,
		post.CreatedAt,
		post.UpdatedAt,
		post.IsOP,
	)

	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return -1, customErrors.ErrAlreadyExists
		}
		return -1, err
	}

	return post.ID, nil
}
func (r *p) GetById(ctx context.Context, boardKey string, id int64) (*model.Post, error) {
	ex := storage.GetExecutor(ctx, r.db)
	query := `
	SELECT
		global_id,
		board_key,
		thread_id,
		root_post_id,
		id,
		text,
		media_json,
		password_hash,
		created_at,
		updated_at,
		is_op
	FROM posts
	WHERE board_key = ? AND id = ?`

	var post model.Post
	var mediaJSON sql.NullString

	err := ex.QueryRowContext(ctx, query, boardKey, id).Scan(
		&post.GlobalID,
		&post.BoardKey,
		&post.ThreadID,
		&post.RootPostID,
		&post.ID,
		&post.Text,
		&mediaJSON,
		&post.PasswordHash,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.IsOP,
	)
	if err == sql.ErrNoRows {
		return nil, customErrors.ErrNotFound
	}
	if err != nil {
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
	ex := storage.GetExecutor(ctx, r.db)
	query := `
	UPDATE posts SET
		text = ?,
		media_json = ?	
	WHERE board_key = ? AND id = ? 
	`
	_, err = ex.ExecContext(ctx, query, post.Text, mediaJSON, post.BoardKey, post.ID)
	return err
}
func (r *p) Delete(ctx context.Context, globalId int64) error {
	ex := storage.GetExecutor(ctx, r.db)
	query := `
	DELETE FROM posts
	WHERE global_id = ?
	`
	_, err := ex.ExecContext(ctx, query, globalId)
	return err
}
func (r *p) List(ctx context.Context, boardKey string, threadId int64) ([]*model.Post, error) {
	ex := storage.GetExecutor(ctx, r.db)
	query := `
	SELECT
		global_id,
		board_key,
		thread_id,
		root_post_id,
		id,
		text,
		media_json,
		password_hash, 
		created_at,
		updated_at,
		is_op
	FROM posts
	WHERE board_key = ? AND root_post_id = ?`

	var mediaJSON sql.NullString

	rows, err := ex.QueryContext(ctx, query, boardKey, threadId)
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
			&post.RootPostID,
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
