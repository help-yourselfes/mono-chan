package service

import (
	"context"

	m "github.com/helpyourselfes/mono-chan/internal/app/board/model"
	"github.com/helpyourselfes/mono-chan/internal/app/board/repo"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
)

type BoardService struct {
	repo repo.BoardRepo
}

type t = BoardService

func NewBoardService(repo repo.BoardRepo) *t {
	return &t{repo: repo}
}

func (s *t) CreateBoard(ctx context.Context, board *m.Board) error {
	if board.Key == "" || board.Caption == "" {
		return customErrors.ErrInvalidInput
	}
	if res, _ := s.repo.GetByKey(ctx, board.Key); res != nil {
		return customErrors.ErrAlreadyExists
	}
	return s.repo.Create(ctx, board)
}

func (s *t) GetBoardByKey(ctx context.Context, key string) (*m.Board, error) {
	if key == "" {
		return nil, customErrors.ErrInvalidInput
	}
	return s.repo.GetByKey(ctx, key)
}

func (s *t) Update(ctx context.Context, key string, board *m.Board) error {
	if key == "" {
		return customErrors.ErrInvalidInput
	}
	return s.repo.Update(ctx, key, board)
}

func (s *t) Delete(ctx context.Context, key string) error {
	if key == "" {
		return customErrors.ErrInvalidInput
	}

	if res, _ := s.repo.GetByKey(ctx, key); res == nil {
		return customErrors.ErrNotFound
	}
	return s.repo.Delete(ctx, key)
}

func (s *t) List(ctx context.Context) ([]*m.Board, error) {
	return s.repo.List(ctx)
}
