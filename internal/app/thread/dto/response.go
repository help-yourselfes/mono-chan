package dto

import (
	"github.com/helpyourselfes/mono-chan/internal/app/post/dto"
)

type ThreadResponse struct {
	PostID     int64  `json:"post_id"`
	Caption    string `json:"caption"`
	BoardKey   string `json:"board_key"`
	ReplyCount int64  `json:"reply_count"`
}

type ThreadPostResponse struct {
	Thread ThreadResponse   `json:"thread"`
	Post   dto.PostResponse `json:"post"`
}
