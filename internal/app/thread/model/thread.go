package model

import "time"

type Thread struct {
	BoardKey   string     `json:"-"`
	Id         int64      `json:"id"`
	Caption    string     `json:"caption"`
	Text       string     `json:"text"`
	MediaLinks []string   `json:"media-links"`
	ReplyCount int64      `json:"reply-count"`
	CreatedAt  *time.Time `json:"created-at"`
	UpdatedAt  *time.Time `json:"updated-at"`
	Password   string     `json:"-"`

	// DeletedAt *time.Time `json:"-"`
}
