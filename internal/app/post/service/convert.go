package service

import (
	"github.com/helpyourselfes/mono-chan/internal/app/post/dto"
	"github.com/helpyourselfes/mono-chan/internal/app/post/model"
)

func postToResponse(post *model.Post) *dto.PostResponse {
	return &dto.PostResponse{
		ID:         post.ID,
		ThreadID:   post.ThreadID,
		Text:       post.Text,
		MediaLinks: post.MediaLinks,
		CreatedAt:  post.CreatedAt,
		UpdatedAt:  post.UpdatedAt,
		IsOP:       post.IsOP,
	}
}
