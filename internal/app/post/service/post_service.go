package service

import (
	"context"
	"time"

	"github.com/helpyourselfes/mono-chan/internal/app/post/dto"
	"github.com/helpyourselfes/mono-chan/internal/app/post/model"
	"github.com/helpyourselfes/mono-chan/internal/app/post/repo"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
	"github.com/helpyourselfes/mono-chan/internal/pkg/security"
)

type PostService struct {
	repo repo.PostRepo
}

type p = PostService

func NewPostService(repo repo.PostRepo) *p {
	return &p{repo: repo}
}

func (s *p) Create(ctx context.Context, reqPost *dto.CreatePostRequest) (*dto.PostResponse, error) {
	if !(reqPost.Text != "" || len(reqPost.MediaLinks) != 0) {
		return nil, customErrors.ErrInvalidInput
	}

	passwordHash, err := security.Hash(reqPost.Password)
	if err != nil {
		return nil, err
	}

	post := &model.Post{
		ThreadID:     reqPost.ThreadID,
		Text:         reqPost.Text,
		PasswordHash: passwordHash,
		MediaLinks:   reqPost.MediaLinks,
		CreatedAt:    time.Now(),
	}

	id, err := s.repo.Create(ctx, post)
	if err != nil {
		return nil, err
	}

	post.ID = id
	resPost := dto.ToPostResponse(post)
	return resPost, nil
}

func (s *p) GetById(ctx context.Context, id int64) (*dto.PostResponse, error) {
	post, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	resPost := dto.ToPostResponse(post)

	return resPost, nil
}
func (s *p) Update(ctx context.Context, id int64, reqPost *dto.UpdatePostRequest) error {
	post, err := s.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	if post.PasswordHash == "" {
		return customErrors.ErrNoPasswordSet
	}

	equal, err := security.Verify(reqPost.Password, post.PasswordHash)
	if err != nil {
		return err
	}
	if !equal {
		return customErrors.ErrIncorectPassword
	}

	return s.repo.Update(ctx, id, reqPost)
}
func (s *p) Delete(ctx context.Context, id int64, password string) error {
	post, err := s.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	equal, err := security.Verify(password, post.PasswordHash)
	if err != nil {
		return err
	}

	if !equal {
		return customErrors.ErrIncorectPassword
	}

	return s.repo.Delete(ctx, id)
}
func (s *p) List(ctx context.Context, threadId int64) ([]*dto.PostResponse, error) {
	posts, err := s.repo.List(ctx, threadId)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
