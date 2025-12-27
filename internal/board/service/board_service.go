package service

import (
	"context"

	m "github.com/helpyourselfes/mono-chan/internal/board/model"
	"github.com/helpyourselfes/mono-chan/internal/board/repo"
	"github.com/helpyourselfes/mono-chan/internal/pkg/errors"
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
		return errors.ErrInvalidInput
	}
	return s.repo.Create(ctx, board)
}

func (s *t) GetBoardByKey(ctx context.Context, key m.Key) (*m.Board, error) {
	if key == "" {
		return nil, errors.ErrInvalidInput
	}
	return s.repo.GetByKey(ctx, key)
}

func (s *t) Update(ctx context.Context, key m.Key, board *m.Board) error {
	if key == "" {
		return errors.ErrInvalidInput
	}
	return s.Update(ctx, key, board)
}

func (s *t) Delete(ctx context.Context, key m.Key) error {
	if key == "" {
		return errors.ErrInvalidInput
	}
	return s.Delete(ctx, key)
}

func (s *t) List(ctx context.Context) ([]*m.Board, error) {
	return s.List(ctx)
}
