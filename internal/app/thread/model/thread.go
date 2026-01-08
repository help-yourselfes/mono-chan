package model

import "github.com/helpyourselfes/mono-chan/internal/app/post/model"

type Thread struct {
	GlobalID     int64
	BoardKey     string
	PostID       int64
	UserHash     string
	PasswordHash string
	Caption      string
	ReplyCount   int64
}

type ThreadPost struct {
	Thread
	model.Post
}
