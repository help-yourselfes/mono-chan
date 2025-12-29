package service

import (
	"context"

	"github.com/helpyourselfes/mono-chan/internal/app/post/dto"
	"github.com/helpyourselfes/mono-chan/internal/app/post/repo"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
)

type PostService struct {
	repo repo.PostRepo
}

type p = PostService

func NewPostService(repo repo.PostRepo) *p {
	return &p{repo: repo}
}

func (s *p) Create(ctx context.Context, post *dto.CreatePostRequest) (int64, error) {
	if !(post.Text != "" || len(post.MediaLinks) != 0) {
		return -1, customErrors.ErrInvalidInput
	}

	res, err := s.repo.Create(ctx, post)
	if err != nil {
		return -1, err
	}

	return res, nil
}
func (s *p) GetById(ctx context.Context, id int64) (*dto.PostResponse, error) {
	post, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	resPost := postToResponse(post)

	return resPost, nil
}
func (s *p) Update(ctx context.Context, id int64, reqPost *dto.UpdatePostRequest) error {
	post, err := s.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	if post.Password == "" {
		return customErrors.ErrNoPasswordSet
	}

	return nil
}
func (s *p) Delete(ctx context.Context, id int64, password string) error {
	if password == "" {
		return customErrors.ErrNoPasswordSet
	}

	return nil
}
func (s *p) List(ctx context.Context, threadId int64) error {
	return nil
}
