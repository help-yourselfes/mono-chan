package dto

type CreatePostRequest struct {
	ThreadID   int64    `json:"thread_id"`
	Text       string   `json:"text"`
	MediaLinks []string `json:"media_links"`
	IsOP       bool     `json:"is_op"`
}

type UpdatePostRequest struct {
	Text       string   `json:"text"`
	MediaLinks []string `json:"media_links"`
	Password   string   `json:"password"`
}
