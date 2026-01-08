package dto

type CreateBoardRequest struct {
	Caption     string `json:"caption"`
	Description string `json:"description"`
}

type UpdateBoardRequest struct {
	CreateBoardRequest
}
