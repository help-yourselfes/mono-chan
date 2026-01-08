package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/helpyourselfes/mono-chan/internal/app/board/validator"
	"github.com/helpyourselfes/mono-chan/internal/logger"
	"github.com/helpyourselfes/mono-chan/internal/logger/sl"
	"github.com/helpyourselfes/mono-chan/internal/pkg/api/response"
)

func (h *BoardHandler) DeleteBoard(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context())
	const op = "boards.handlers.delete"
	log = log.With(slog.String("op", op))

	key := chi.URLParam(r, "key")
	if !validator.IsValidKey(key) {
		log.Error("key is invalid")
		render.JSON(w, r, response.Error("key is invalid"))
		return
	}

	err := h.service.Delete(r.Context(), key)
	if err != nil {
		log.Error("failed to delete board:", sl.Err(err))
		http.Error(
			w,
			"failed to delete board",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
