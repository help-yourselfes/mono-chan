package service

import (
	"context"
	"time"

	"github.com/helpyourselfes/mono-chan/internal/app"
	"github.com/helpyourselfes/mono-chan/internal/app/post/dto"
	"github.com/helpyourselfes/mono-chan/internal/app/post/model"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
	"github.com/helpyourselfes/mono-chan/internal/pkg/security"

	_ "github.com/helpyourselfes/mono-chan/internal/app/board/repo"
	_ "github.com/helpyourselfes/mono-chan/internal/app/post/repo"
	_ "github.com/helpyourselfes/mono-chan/internal/app/thread/repo"
)

type PostService struct {
	app.Repos
	txManager app.TransactionManager
}

func NewPostService(repos *app.Repos, tx app.TransactionManager) *PostService {
	return &PostService{
		Repos:     *repos,
		txManager: tx,
	}
}

func (s *PostService) Create(ctx context.Context, reqPost *dto.CreatePostRequest) (*dto.PostResponse, error) {
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
		RootPostID:   reqPost.RootPostID,
		BoardKey:     reqPost.BoardKey,
		Text:         reqPost.Text,
		PasswordHash: passwordHash,
		MediaLinks:   reqPost.MediaLinks,
		CreatedAt:    now,
		UpdatedAt:    nil,
	}
	var id int64
	err := s.txManager.WithinTransaction(ctx, func(ctx context.Context) error {
		var err error

		id, err = s.Boards.IncPosts(ctx, post.BoardKey)
		if err != nil {
			return err
		}

		thread, err := s.Threads.GetByPostID(ctx, post.BoardKey, reqPost.RootPostID)
		if err != nil {
			return err
		}

		post.ID = id
		post.ThreadID = thread.GlobalID

		_, err = s.Posts.Create(ctx, post)
		if err != nil {
			return err
		}

		return s.Threads.Reply(ctx, post.BoardKey, thread.PostID)
	})

	post.ID = id
	resPost := dto.ToPostResponse(post)
	return resPost, err
}

func (s *PostService) GetById(ctx context.Context, boardKey string, id int64) (*dto.PostResponse, error) {
	post, err := s.Posts.GetById(ctx, boardKey, id)
	if err != nil {
		return nil, err
	}
	resPost := dto.ToPostResponse(post)

	return resPost, nil
}

func (s *PostService) Update(ctx context.Context, reqPost *dto.UpdatePostRequest) error {
	post, err := s.Posts.GetById(ctx, reqPost.BoardKey, reqPost.ID)
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

	return s.Posts.Update(ctx, reqPost)
}

func (s *PostService) DeleteByUser(ctx context.Context, boardKey string, id int64, password string) error {
	post, err := s.Posts.GetById(ctx, boardKey, id)

	if err != nil {
		return err
	}
	if !post.HasPassword() {
		return customErrors.ErrNoPasswordSet
	}
	if post.ID == post.RootPostID {
		return customErrors.ErrPostIsRoot
	}

	equal, err := security.Verify(password, post.PasswordHash)
	if err != nil {
		return err
	}

	if !equal {
		return customErrors.ErrIncorectPassword
	}

	return s.Posts.Delete(ctx, post.GlobalID)
}

func (s *PostService) DeleteByAdmin(ctx context.Context, boardKey string, id int64) error {
	post, err := s.Posts.GetById(ctx, boardKey, id)
	if err != nil {
		return err
	}

	if post.ID == post.RootPostID {
		return customErrors.ErrPostIsRoot
	}
	return s.Posts.Delete(ctx, post.GlobalID)
}

func (s *PostService) List(ctx context.Context, boardKey string, threadId int64) ([]*dto.PostResponse, error) {
	posts, err := s.Posts.List(ctx, boardKey, threadId)
	if err != nil {
		return nil, err
	}
	resPosts := make([]*dto.PostResponse, len(posts))
	for i := range posts {
		resPosts[i] = dto.ToPostResponse(posts[i])
	}

	return resPosts, nil
}
