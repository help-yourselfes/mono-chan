package repo

import (
	"context"

	m "github.com/helpyourselfes/mono-chan/internal/board/model"
)

type BoardRepo interface {
	Create(ctx context.Context, board *m.Board) error
	GetByKey(ctx context.Context, key m.Key) (*m.Board, error)
	Update(ctx context.Context, key m.Key, board *m.Board) error
	Delete(ctx context.Context, key m.Key) error
	List(ctx context.Context) ([]*m.Board, error)
}
