package dto

import "time"

type PostResponse struct {
	ID         int64      `json:"id"`
	BoardKey   string     `json:"board_key"`
	RootPostID int64      `json:"thread_id"`
	Text       string     `json:"text"`
	MediaLinks []string   `json:"media_links"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	IsOP       bool       `json:"is_op"`
}
