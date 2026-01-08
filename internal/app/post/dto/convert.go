package dto

import "github.com/helpyourselfes/mono-chan/internal/app/post/model"

func ToPostResponse(post *model.Post) *PostResponse {
	return &PostResponse{
		ID:         post.ID,
		RootPostID: post.RootPostID,
		Text:       post.Text,
		MediaLinks: post.MediaLinks,
		CreatedAt:  post.CreatedAt,
		UpdatedAt:  post.UpdatedAt,
		IsOP:       post.IsOP,
	}
}
