package storage

import (
	"context"
	"database/sql"
)

type txKey struct{}

type SqlTxManager struct {
	db *sql.DB
}

func NewSqlTxManager(db *sql.DB) *SqlTxManager {
	return &SqlTxManager{db: db}
}

func (m *SqlTxManager) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	ctxWithTx := context.WithValue(ctx, txKey{}, tx)

	err = fn(ctxWithTx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func GetTx(ctx context.Context) (*sql.Tx, bool) {
	tx, ok := ctx.Value(txKey{}).(*sql.Tx)
	return tx, ok
}

type DBExecutor interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

func GetExecutor(ctx context.Context, defaultDB *sql.DB) DBExecutor {
	if tx, ok := GetTx(ctx); ok {
		return tx
	}
	return defaultDB
}
