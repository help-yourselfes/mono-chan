package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/helpyourselfes/mono-chan/internal/board/model"
	valid "github.com/helpyourselfes/mono-chan/internal/board/validator"
	resp "github.com/helpyourselfes/mono-chan/internal/pkg/api/response"
)

func (h *BoardHandler) UpdateBoard(w http.ResponseWriter, r *http.Request) {
	var req BoardRequest

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		render.JSON(w, r, resp.Error("failed to decode request"))
		return
	}

	key := req.Key
	if !valid.IsValidKey(key) {
		render.JSON(w, r, resp.Error("key is invalid"))
		return
	}

	board := &model.Board{
		Key:         req.Key,
		Caption:     req.Caption,
		Description: req.Description,
	}

	if err := h.service.Update(r.Context(), board.Key, board); err != nil {
		http.Error(w, fmt.Errorf("failed to update board: %w", err).Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(board)
}
