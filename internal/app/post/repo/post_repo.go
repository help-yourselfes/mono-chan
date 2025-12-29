package repo

import (
	"context"

	"github.com/helpyourselfes/mono-chan/internal/app/post/dto"
	"github.com/helpyourselfes/mono-chan/internal/app/post/model"
)

type PostRepo interface {
	Create(ctx context.Context, post *dto.CreatePostRequest) (int64, error)
	GetById(ctx context.Context, id int64) (*model.Post, error)
	Update(ctx context.Context, id int64, post *dto.UpdatePostRequest) error
	Delete(ctx context.Context, id int64, password string) error
	List(ctx context.Context, threadId int64) error
}
