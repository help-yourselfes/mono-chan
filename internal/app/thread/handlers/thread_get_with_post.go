package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/helpyourselfes/mono-chan/internal/app/thread/dto"
	"github.com/helpyourselfes/mono-chan/internal/logger"
	"github.com/helpyourselfes/mono-chan/internal/logger/sl"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
)

func (h *ThreadHandler) GetWithPost(w http.ResponseWriter, r *http.Request) {
	const op = "thread.handlers.getByPostId"
	log := logger.FromContext(r.Context())
	log = log.With(slog.String("op", op))

	var req dto.GetThreadRequest

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		const msg = "failed to decode request"
		log.Error(msg, sl.Err(err))
		render.JSON(w, r, msg)
		return
	}

	threadWithPost, err := h.service.GetWithPost(r.Context(), req.BoardKey, req.PostID)
	if errors.Is(err, customErrors.ErrNotFound) {
		const msg = "not found"
		log.Error(msg)
		render.JSON(w, r, msg)
		return
	}
	if err != nil {
		const msg = "unknown error"
		log.Error(msg, sl.Err(err))
		render.JSON(w, r, msg)
		return
	}

	render.JSON(w, r, threadWithPost)
}
