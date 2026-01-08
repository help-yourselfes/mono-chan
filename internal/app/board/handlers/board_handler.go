package handlers

import "github.com/helpyourselfes/mono-chan/internal/app/board/service"

type BoardHandler struct {
	service *service.BoardService
}

func NewBoardHandler(s *service.BoardService) *BoardHandler {
	return &BoardHandler{
		service: s,
	}
}
