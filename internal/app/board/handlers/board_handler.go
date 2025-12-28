package handlers

import "github.com/helpyourselfes/mono-chan/internal/app/board/service"

type BoardHandler struct {
	service *service.BoardService
}

type BoardRequest struct {
	Key         string `json:"key" validate:"required"`
	Caption     string `json:"caption"`
	Description string `json:"description"`
}

func NewBoardHandler(s *service.BoardService) *BoardHandler {
	return &BoardHandler{
		service: s,
	}
}
