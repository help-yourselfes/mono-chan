package model

import "time"

type Post struct {
	GlobalID     int64
	ID           int64
	BoardKey     string
	ThreadID     int64
	Text         string
	MediaLinks   []string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	IsOP         bool
}

func (p Post) HasPassword() bool {
	return p.PasswordHash != ""
}
