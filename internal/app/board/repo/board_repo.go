package repo

import (
	"context"

	m "github.com/helpyourselfes/mono-chan/internal/app/board/model"
)

type BoardRepo interface {
	Create(ctx context.Context, board *m.Board) error
	GetByKey(ctx context.Context, key string) (*m.Board, error)
	Update(ctx context.Context, key string, board *m.Board) error
	Delete(ctx context.Context, key string) error
	List(ctx context.Context) ([]*m.Board, error)
	IncPosts(ctx context.Context, boardKey string) (int64, error)
}
