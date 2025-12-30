package repo

import (
	"context"

	"github.com/helpyourselfes/mono-chan/internal/app/post/dto"
	"github.com/helpyourselfes/mono-chan/internal/app/post/model"
)

type PostRepo interface {
	Create(ctx context.Context, post *model.Post) (int64, error)
	GetById(ctx context.Context, boardKey string, id int64) (*model.Post, error)
	Update(ctx context.Context, post *dto.UpdatePostRequest) error
	Delete(ctx context.Context, globalId int64) error
	List(ctx context.Context, boardKey string, threadId int64) ([]*model.Post, error)
}
