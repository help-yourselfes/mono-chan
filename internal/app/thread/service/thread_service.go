package service

import (
	"context"
	"time"

	m "github.com/helpyourselfes/mono-chan/internal/app/thread/model"
	"github.com/helpyourselfes/mono-chan/internal/app/thread/repo"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
	"github.com/helpyourselfes/mono-chan/internal/pkg/security"
)

type ThreadService struct {
	repo repo.ThreadRepo
}

type t = ThreadService

func NewThreadService(repo repo.ThreadRepo) *t {
	return &t{repo: repo}
}

func (s *t) CreateThread(ctx context.Context, thread *m.Thread) (int64, error) {
	if /* at least one field must be filled*/
	!(thread.Caption != "" ||
		thread.Text != "" ||
		len(thread.MediaLinks) != 0) {
		return -1, customErrors.ErrInvalidInput
	}

	if thread.Password != "" {
		hash, err := security.Hash(thread.Password)
		if err != nil {
			return -1, err
		}
		thread.Password = hash
	}

	return s.repo.Create(ctx, thread)
}

func (s *t) GetThreadByID(ctx context.Context, boardKey string, id int64) (*m.Thread, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *t) UpdateThread(ctx context.Context, inputThread *m.Thread) error {
	existingThread, err := s.repo.GetByID(ctx, inputThread.Id)
	if err != nil {
		return customErrors.ErrNotFound
	}

	if existingThread.Password == "" {
		return customErrors.ErrUpdateToUpdate
	}

	eq, err := security.Verify(inputThread.Password, existingThread.Password)
	if err != nil {
		return err
	}
	if !eq {
		return customErrors.ErrInvalidPassword
	}

	existingThread.Caption = inputThread.Caption
	existingThread.Text = inputThread.Text
	existingThread.MediaLinks = inputThread.MediaLinks

	now := time.Now()
	existingThread.UpdatedAt = &now

	return s.repo.Update(ctx, existingThread)
}

func (s *t) DeleteThread(ctx context.Context, boardKey string, id int64, password string) error {
	res, _ := s.repo.GetByID(ctx, id)
	if res == nil {
		return customErrors.ErrNotFound
	}
	if res.Password == "" {
		return customErrors.ErrUnableToDelete
	}

	eq, err := security.Verify(password, res.Password)
	if err != nil {
		return err
	}
	if !eq {
		return customErrors.ErrInvalidPassword
	}
	return s.repo.Delete(ctx, id, password)
}
func (s *t) ListThread(ctx context.Context, boardKey string) ([]*m.Thread, error) {
	return s.repo.List(ctx, boardKey)
}
func (s *t) ReplyThread(ctx context.Context, boardKey string, id int64) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Reply(ctx, id)
}
