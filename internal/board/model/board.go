package model

type Key string

type Board struct {
	Key         `json:"id"`
	Caption     string `json:"caption"`
	Description string `json:"description"`
}
