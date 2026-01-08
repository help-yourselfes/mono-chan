package repo

import (
	"context"
	"database/sql"
	"errors"

	m "github.com/helpyourselfes/mono-chan/internal/app/board/model"
	"github.com/helpyourselfes/mono-chan/internal/app/board/repo"
	"github.com/helpyourselfes/mono-chan/internal/infrastructure/storage"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
	"github.com/mattn/go-sqlite3"
)

type sqliteBoardRepo struct {
	db *sql.DB
}

var _ repo.BoardRepo = &sqliteBoardRepo{}

func NewSQLiteBoardRepo(db *sql.DB) *sqliteBoardRepo {
	return &sqliteBoardRepo{db: db}
}

func (r *sqliteBoardRepo) Create(ctx context.Context, board *m.Board) error {
	ex := storage.GetExecutor(ctx, r.db)
	query :=
		`
		INSERT INTO boards (
			key, caption, description, last_post_id
		) VALUES (?,?,?,?)
	`

	_, err := ex.ExecContext(ctx, query, board.Key, board.Caption, board.Description, board.LastPostID)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return customErrors.ErrAlreadyExists
		}
		return err
	}

	return nil
}

func (r *sqliteBoardRepo) GetByKey(ctx context.Context, key string) (*m.Board, error) {
	ex := storage.GetExecutor(ctx, r.db)
	query := `
		SELECT * FROM boards
		WHERE key = ?
	`
	row := ex.QueryRowContext(ctx, query, key)

	var board m.Board
	if err := row.Scan(&board.Key, &board.Caption, &board.Description, &board.LastPostID); err != nil {
		if err == sql.ErrNoRows {
			return nil, customErrors.ErrNotFound
		}
		return nil, err
	}
	return &board, nil
}

func (r *sqliteBoardRepo) Update(ctx context.Context, key string, board *m.Board) error {
	ex := storage.GetExecutor(ctx, r.db)
	query := `
		UPDATE boards
		SET caption = ?, description = ?
		WHERE key = ?
	`

	_, err := ex.ExecContext(ctx, query, board.Caption, board.Description, board.Key)
	if err != nil {
		return err
	}

	return nil
}

func (r *sqliteBoardRepo) Delete(ctx context.Context, key string) error {
	ex := storage.GetExecutor(ctx, r.db)
	query := `
	DELETE FROM boards
	WHERE key = ?
	`
	_, err := ex.ExecContext(ctx, query, key)
	return err
}

func (r *sqliteBoardRepo) List(ctx context.Context) ([]*m.Board, error) {
	ex := storage.GetExecutor(ctx, r.db)
	query := `
		SELECT * FROM boards
		ORDER BY key
	`
	rows, err := ex.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	boards := make([]*m.Board, 0)
	for rows.Next() {
		var b m.Board
		if err := rows.Scan(&b.Key, &b.Caption, &b.Description, &b.LastPostID); err != nil {
			return nil, err
		}
		boards = append(boards, &b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return boards, nil
}

func (r *sqliteBoardRepo) IncPosts(ctx context.Context, boardKey string) (int64, error) {
	ex := storage.GetExecutor(ctx, r.db)
	var localId int64
	query := `
	UPDATE boards 
	SET last_post_id = last_post_id + 1 
	WHERE key = ? 
	RETURNING last_post_id;
	`
	err := ex.QueryRowContext(ctx, query, boardKey).Scan(&localId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, customErrors.ErrNotFound
		}
		return -1, err
	}
	return localId, nil
}
