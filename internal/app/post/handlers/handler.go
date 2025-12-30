package handlers

import "github.com/helpyourselfes/mono-chan/internal/app/post/service"

type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(s *service.PostService) *PostHandler {
	return &PostHandler{
		service: s,
	}
}
