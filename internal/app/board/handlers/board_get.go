package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	valid "github.com/helpyourselfes/mono-chan/internal/app/board/validator"
	"github.com/helpyourselfes/mono-chan/internal/logger/sl"
	resp "github.com/helpyourselfes/mono-chan/internal/pkg/api/response"
)

func (h *BoardHandler) GetBoardByKey(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "boards.handlers.getByKey"
		log = log.With(slog.String("op", op))

		key := chi.URLParam(r, "key")
		if !valid.IsValidKey(key) {
			log.Error("invalid key")
			render.JSON(w, r, resp.Error("key is invalid"))
			return
		}

		board, err := h.service.GetBoardByKey(r.Context(), key)
		if err != nil {
			log.Error("failed to get board", sl.Err(err))
			http.Error(w, "failed to get board", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(board)
	}
}
