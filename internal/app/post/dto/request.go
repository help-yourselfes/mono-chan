package dto

type CreatePostRequest struct {
	ThreadID   int64    `json:"thread_id"`
	Text       string   `json:"text"`
	MediaLinks []string `json:"media_links"`
	IsOP       bool     `json:"is_op"`
	Password   string   `json:"password"`
}

type UpdatePostRequest struct {
	Text       string   `json:"text"`
	MediaLinks []string `json:"media_links"`
	Password   string   `json:"password"`
}
