package model

import "time"

type Post struct {
	ID         int64
	ThreadID   int64
	Text       string
	MediaLinks []string
	Password   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	IsOP       bool
}
