package handlers

import (
	postDTO "github.com/helpyourselfes/mono-chan/internal/app/post/dto"
	threadDTO "github.com/helpyourselfes/mono-chan/internal/app/thread/dto"
	"github.com/helpyourselfes/mono-chan/internal/app/thread/service"
)

type ThreadHandler struct {
	service *service.ThreadService
}

func NewThreadHandler(s *service.ThreadService) *ThreadHandler {
	return &ThreadHandler{
		service: s,
	}
}

type createThreadWithPostRequest struct {
	postDTO.CreatePostRequest     `json:"post"`
	threadDTO.CreateThreadRequest `json:"thread"`
}
