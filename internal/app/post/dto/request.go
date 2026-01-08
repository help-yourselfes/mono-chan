package dto

type CreatePostRequest struct {
	RootPostID int64    `json:"thread_id"`
	BoardKey   string   `json:"board_key"`
	Text       string   `json:"text"`
	MediaLinks []string `json:"media_links"`
	IsOP       bool     `json:"is_op"`
	Password   string   `json:"password"`
}

type UpdatePostRequest struct {
	ID         int64    `json:"id"`
	BoardKey   string   `json:"board_key"`
	Text       string   `json:"text"`
	MediaLinks []string `json:"media_links"`
	Password   string   `json:"password"`
}

type AdminDeletePostRequest struct {
	BoardKey string `json:"board_key"`
	ID       int64  `json:"id"`
}

type UserDeletePostRequest struct {
	AdminDeletePostRequest
	Password string `json:"password"`
}
