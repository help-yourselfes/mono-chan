package dto

import "github.com/helpyourselfes/mono-chan/internal/app/thread/model"

func ToThreadResponse(thread *model.Thread) *ThreadResponse {
	return &ThreadResponse{
		PostID:     thread.PostID,
		Caption:    thread.Caption,
		BoardKey:   thread.BoardKey,
		ReplyCount: thread.ReplyCount,
	}
}
