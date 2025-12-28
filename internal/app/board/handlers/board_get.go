package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	valid "github.com/helpyourselfes/mono-chan/internal/app/board/validator"
	resp "github.com/helpyourselfes/mono-chan/internal/pkg/api/response"
)

func (h *BoardHandler) GetBoardByKey(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	if !valid.IsValidKey(key) {
		render.JSON(w, r, resp.Error("key is invalid"))
		return
	}

	board, err := h.service.GetBoardByKey(r.Context(), key)
	if err != nil {
		http.Error(w, "failed to get board", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(board)
}
