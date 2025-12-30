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
	var passwordHash string

	if reqPost.Password != "" {
		var err error
		passwordHash, err = security.Hash(reqPost.Password)
		if err != nil {
			return nil, err
		}
	}
	now := time.Now()
	post := &model.Post{
		ThreadID:     reqPost.ThreadID,
		Text:         reqPost.Text,
		PasswordHash: passwordHash,
		MediaLinks:   reqPost.MediaLinks,
		CreatedAt:    now,
	}

	id, err := s.repo.Create(ctx, post)
	if err != nil {
		return nil, err
	}

	post.ID = id
	resPost := dto.ToPostResponse(post)
	return resPost, nil
}

func (s *p) GetById(ctx context.Context, boardKey string, id int64) (*dto.PostResponse, error) {
	post, err := s.repo.GetById(ctx, boardKey, id)
	if err != nil {
		return nil, err
	}
	resPost := dto.ToPostResponse(post)

	return resPost, nil
}

func (s *p) Update(ctx context.Context, reqPost *dto.UpdatePostRequest) error {
	post, err := s.repo.GetById(ctx, reqPost.BoardKey, reqPost.ID)
	if err != nil {
		return err
	}

	if !post.HasPassword() {
		return customErrors.ErrNoPasswordSet
	}

	equal, err := security.Verify(reqPost.Password, post.PasswordHash)
	if err != nil {
		return err
	}
	if !equal {
		return customErrors.ErrIncorectPassword
	}

	return s.repo.Update(ctx, reqPost)
}

func (s *p) DeleteByUser(ctx context.Context, boardKey string, id int64, password string) error {
	post, err := s.repo.GetById(ctx, boardKey, id)
	if err != nil {
		return err
	}

	if !post.HasPassword() {
		return customErrors.ErrNoPasswordSet
	}

	equal, err := security.Verify(password, post.PasswordHash)
	if err != nil {
		return err
	}

	if !equal {
		return customErrors.ErrIncorectPassword
	}

	return s.repo.Delete(ctx, post.GlobalID)
}

func (s *p) DeleteByAdmin(ctx context.Context, boardKey string, id int64) error {
	post, err := s.repo.GetById(ctx, boardKey, id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, post.GlobalID)
}

func (s *p) List(ctx context.Context, boardKey string, threadId int64) ([]*dto.PostResponse, error) {
	posts, err := s.repo.List(ctx, boardKey, threadId)
	if err != nil {
		return nil, err
	}
	resPosts := make([]*dto.PostResponse, len(posts))
	for i := range posts {
		resPosts[i] = dto.ToPostResponse(posts[i])
	}

	return resPosts, nil
}
