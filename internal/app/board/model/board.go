package model

type Board struct {
	Key         string `json:"id"`
	Caption     string `json:"caption"`
	Description string `json:"description"`
	LastPostID  int64  `json:"-"`
}
