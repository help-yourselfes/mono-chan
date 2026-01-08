package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/helpyourselfes/mono-chan/internal/app/board/model"
	valid "github.com/helpyourselfes/mono-chan/internal/app/board/validator"
	"github.com/helpyourselfes/mono-chan/internal/logger"
	"github.com/helpyourselfes/mono-chan/internal/logger/sl"
	resp "github.com/helpyourselfes/mono-chan/internal/pkg/api/response"
)

func (h *BoardHandler) UpdateBoard(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context())
	const op = "boards.handlers.update"
	log = log.With(slog.String("op", op))
	var req BoardRequest

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		log.Error("failed to decode request", sl.Err(err))
		render.JSON(w, r, resp.Error("failed to decode request"))
		return
	}

	key := req.Key
	if !valid.IsValidKey(key) {
		log.Error("key is invalid")
		render.JSON(w, r, resp.Error("key is invalid"))
		return
	}

	board := &model.Board{
		Key:         req.Key,
		Caption:     req.Caption,
		Description: req.Description,
	}

	if err := h.service.Update(r.Context(), board.Key, board); err != nil {
		log.Error("failed to update board", sl.Err(err))
		http.Error(w, fmt.Errorf("failed to update board: %w", err).Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(board)
}
