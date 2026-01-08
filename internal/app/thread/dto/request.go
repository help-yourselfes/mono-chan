package dto

type CreateThreadRequest struct {
	BoardKey string `json:"board_key"`
	UserHash string `json:"user_hash"`
	Caption  string `json:"caption"`
	Password string `json:"password"`
}

type UpdateThreadRequest struct {
	BoardKey string `json:"board_key"`
	PostID   int64  `json:"post_id"`
	Caption  string `json:"caption"`
	Password string `json:"password"`
}

type DeleteAdminThreadRequest struct {
	BoardKey string `json:"board_key"`
	PostID   int64  `json:"post_id"`
}

type DeleteUserThreadRequest struct {
	DeleteAdminThreadRequest
	Password string `json:"password"`
}

type GetThreadRequest struct {
	BoardKey string `json:"board_key"`
	PostID   int64  `json:"post_id"`
}
