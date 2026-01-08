package repo

import (
	"context"

	m "github.com/helpyourselfes/mono-chan/internal/app/thread/model"
)

type ThreadRepo interface {
	Create(ctx context.Context, thread *m.Thread) (int64, error)
	GetByGlobalID(ctx context.Context, globalID int64) (*m.Thread, error)
	GetByPostID(ctx context.Context, boardKey string, postID int64) (*m.Thread, error)
	Update(ctx context.Context, thread *m.Thread) error
	Delete(ctx context.Context, globalID int64) error
	List(ctx context.Context, boardKey string) ([]*m.Thread, error)
	ListWithPost(ctx context.Context, boardKey string) ([]*m.ThreadPost, error)
	Reply(ctx context.Context, boardKey string, postID int64) error
}
