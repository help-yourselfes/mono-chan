package model

type Thread struct {
	GlobalID   int    `json:"-"`
	PostID     int64  `json:"post_id"`
	Caption    string `json:"caption"`
	BoardKey   string `json:"board_key"`
	ReplyCount int64  `json:"reply_count"`
}
