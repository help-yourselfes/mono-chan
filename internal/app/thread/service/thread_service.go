package service

import (
	"context"
	"time"

	"github.com/helpyourselfes/mono-chan/internal/app"
	postDTO "github.com/helpyourselfes/mono-chan/internal/app/post/dto"
	postModel "github.com/helpyourselfes/mono-chan/internal/app/post/model"
	threadDTO "github.com/helpyourselfes/mono-chan/internal/app/thread/dto"
	threadModel "github.com/helpyourselfes/mono-chan/internal/app/thread/model"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
	"github.com/helpyourselfes/mono-chan/internal/pkg/security"
)

type ThreadService struct {
	app.Repos
	txManager app.TransactionManager
}

func NewThreadService(repos *app.Repos, tx app.TransactionManager) *ThreadService {
	return &ThreadService{
		Repos:     *repos,
		txManager: tx,
	}
}

func (s *ThreadService) Create(ctx context.Context,
	reqPost *postDTO.CreatePostRequest,
	reqThread *threadDTO.CreateThreadRequest) (*threadDTO.ThreadResponse, error) {

	var passwordHash string

	if reqPost.Password != "" {
		var err error
		passwordHash, err = security.Hash(reqPost.Password)
		if err != nil {
			return nil, err
		}
	}

	thread := &threadModel.Thread{
		PostID:       -1,
		UserHash:     reqThread.UserHash,
		Caption:      reqThread.Caption,
		BoardKey:     reqThread.BoardKey,
		ReplyCount:   0,
		PasswordHash: passwordHash,
	}
	err := s.txManager.WithinTransaction(ctx, func(ctx context.Context) error {
		postId, err := s.Boards.IncPosts(ctx, reqPost.BoardKey)
		if err != nil {
			return err
		}

		thread.PostID = postId

		threadId, err := s.Threads.Create(ctx, thread)
		thread.GlobalID = threadId
		if err != nil {
			return err
		}

		now := time.Now()
		post := &postModel.Post{
			BoardKey:     reqPost.BoardKey,
			ID:           postId,
			ThreadID:     threadId,
			RootPostID:   postId,
			Text:         reqPost.Text,
			MediaLinks:   reqPost.MediaLinks,
			PasswordHash: passwordHash,
			CreatedAt:    now,
			IsOP:         true,
		}

		_, err = s.Posts.Create(ctx, post)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	resThread := threadDTO.ToThreadResponse(thread)
	return resThread, nil
}

func (s *ThreadService) GetByPostID(ctx context.Context, boardKey string, id int64) (*threadDTO.ThreadResponse, error) {
	thread, err := s.Threads.GetByPostID(ctx, boardKey, id)
	if err != nil {
		return nil, err
	}
	resThread := threadDTO.ToThreadResponse(thread)
	return resThread, nil
}

func (s *ThreadService) Update(ctx context.Context, boardKey string, reqThread *threadDTO.UpdateThreadRequest) error {
	thread, err := s.Threads.GetByPostID(ctx, boardKey, reqThread.PostID)
	if err != nil {
		return err
	}
	if thread.PasswordHash == "" {
		return customErrors.ErrNoPasswordSet
	}

	eq, err := security.Verify(reqThread.Password, thread.PasswordHash)
	if !eq {
		return customErrors.ErrIncorectPassword
	}
	if err != nil {
		return err
	}

	newThread := &threadModel.Thread{
		BoardKey: boardKey,
		PostID:   reqThread.PostID,
		Caption:  reqThread.Caption,
	}

	return s.Threads.Update(ctx, newThread)
}

func (s *ThreadService) GetWithPost(ctx context.Context, boardKey string, postID int64) (*threadDTO.ThreadPostResponse, error) {
	thread, err := s.Threads.GetByPostID(ctx, boardKey, postID)
	if err != nil {
		return nil, err
	}

	post, err := s.Posts.GetById(ctx, boardKey, postID)
	if err != nil {
		return nil, err
	}

	var res = &threadDTO.ThreadPostResponse{
		Post:   *postDTO.ToPostResponse(post),
		Thread: *threadDTO.ToThreadResponse(thread),
	}
	return res, err
}

func (s *ThreadService) DeleteByAdmin(ctx context.Context, boardKey string, id int64) error {
	thread, err := s.Threads.GetByPostID(ctx, boardKey, id)
	if err != nil {
		return err
	}

	return s.Threads.Delete(ctx, thread.GlobalID)
}

func (s *ThreadService) DeleteByUser(ctx context.Context, boardKey string, id int64, password string) error {
	thread, err := s.Threads.GetByPostID(ctx, boardKey, id)
	if err != nil {
		return err
	}
	if thread.PasswordHash == "" {
		return customErrors.ErrNoPasswordSet
	}
	eq, err := security.Verify(password, thread.PasswordHash)
	if err != nil {
		return err
	}
	if !eq {
		return customErrors.ErrIncorectPassword
	}

	return s.Threads.Delete(ctx, thread.GlobalID)
}

func (s *ThreadService) List(ctx context.Context, boardKey string) ([]*threadDTO.ThreadResponse, error) {
	threads, err := s.Threads.List(ctx, boardKey)
	if err != nil {
		return nil, err
	}

	resThreads := make([]*threadDTO.ThreadResponse, len(threads))
	for i := range threads {
		resThreads[i] = threadDTO.ToThreadResponse(threads[i])
	}

	return resThreads, nil
}

func (s *ThreadService) ListWithPost(ctx context.Context, boardKey string) ([]*threadDTO.ThreadPostResponse, error) {
	threadsWithPost, err := s.Threads.ListWithPost(ctx, boardKey)
	if err != nil {
		return nil, err
	}

	resThreadsWithPost := make([]*threadDTO.ThreadPostResponse, len(threadsWithPost))
	for i := range threadsWithPost {
		resThreadsWithPost[i] = &threadDTO.ThreadPostResponse{
			Thread: *threadDTO.ToThreadResponse(&threadsWithPost[i].Thread),
			Post:   *postDTO.ToPostResponse(&threadsWithPost[i].Post),
		}
	}

	return resThreadsWithPost, nil
}
