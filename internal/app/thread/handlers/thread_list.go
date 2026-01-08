package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/helpyourselfes/mono-chan/internal/logger"
	"github.com/helpyourselfes/mono-chan/internal/logger/sl"
)

func (h *ThreadHandler) List(w http.ResponseWriter, r *http.Request) {
	const op = "thread.handlers.list"
	log := logger.FromContext(r.Context())
	log = log.With(slog.String("op", op))

	boardKey := chi.URLParam(r, "boardKey")
	if boardKey == "" {
		const msg = "invalid board key"
		log.Error(msg)
		render.JSON(w, r, msg)
		return
	}

	threads, err := h.service.List(r.Context(), boardKey)
	if err != nil {
		const msg = "unknown error"
		log.Error(msg, sl.Err(err))
		render.JSON(w, r, msg)
		return
	}

	render.JSON(w, r, threads)
}
