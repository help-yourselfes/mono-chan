package repo

import (
	"context"
	"database/sql"

	m "github.com/helpyourselfes/mono-chan/internal/app/board/model"
	"github.com/helpyourselfes/mono-chan/internal/app/board/repo"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
	"github.com/mattn/go-sqlite3"
)

type sqliteBoardRepo struct {
	db *sql.DB
}
type t = sqliteBoardRepo

var _ repo.BoardRepo = &t{}

func NewSQLiteBoardRepo(db *sql.DB) *t {
	return &t{db: db}
}

func (r *t) Create(ctx context.Context, board *m.Board) error {
	query :=
		`
		INSERT INTO boards (
			key, caption, description
		) VALUES (?,?,?)
	`

	_, err := r.db.ExecContext(ctx, query, board.Key, board.Caption, board.Description)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return customErrors.ErrAlreadyExists
		}
		return err
	}

	return nil
}

func (r *t) GetByKey(ctx context.Context, key string) (*m.Board, error) {
	query := `
		SELECT * FROM boards
		WHERE key = ?
	`
	row := r.db.QueryRowContext(ctx, query, key)

	var board m.Board
	if err := row.Scan(&board.Key, &board.Caption, &board.Description); err != nil {
		if err == sql.ErrNoRows {
			return nil, customErrors.ErrNotFound
		}
		return nil, err
	}
	return &board, nil
}

func (r *t) Update(ctx context.Context, key string, board *m.Board) error {
	query := `
		UPDATE boards
		SET caption = ?, description = ?
		WHERE key = ?
	`

	_, err := r.db.Exec(query, board.Caption, board.Description, board.Key)
	if err != nil {
		return err
	}

	return nil
}

func (r *t) Delete(ctx context.Context, key string) error {
	query := `
		DELETE FROM boards
		WHERE key = ?
	`
	_, err := r.db.Exec(query, key)
	return err
}

func (r *t) List(ctx context.Context) ([]*m.Board, error) {
	query := `
		SELECT * FROM boards
		ORDER BY key
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	boards := make([]*m.Board, 0)
	for rows.Next() {
		var b m.Board
		if err := rows.Scan(&b.Key, &b.Caption, &b.Description); err != nil {
			return nil, err
		}
		boards = append(boards, &b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return boards, nil
}
