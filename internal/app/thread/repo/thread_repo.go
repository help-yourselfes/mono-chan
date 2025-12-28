package repo

import (
	"context"

	m "github.com/helpyourselfes/mono-chan/internal/app/thread/model"
)

type ThreadRepo interface {
	Create(ctx context.Context, thread *m.Thread) (int64, error)
	GetByID(ctx context.Context, id int64) (*m.Thread, error)
	Update(ctx context.Context, thread *m.Thread) error
	Delete(ctx context.Context, id int64, password string) error
	List(ctx context.Context, boardKey string) ([]*m.Thread, error)
	Reply(ctx context.Context, id int64) error
}
