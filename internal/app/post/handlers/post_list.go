package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/helpyourselfes/mono-chan/internal/logger"
	"github.com/helpyourselfes/mono-chan/internal/logger/sl"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
)

func (h *PostHandler) List(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context())
	const op = "post.handlers.getById"
	log.With(slog.String("op", op))
	boardKey := chi.URLParam(r, "boardKey")
	if boardKey == "" {
		const msg = "no board key provided"
		log.Error(msg)
		render.JSON(w, r, msg)
		return
	}

	threadStr := chi.URLParam(r, "threadId")
	if threadStr == "" {
		const msg = "no thread id provided"
		log.Error(msg)
		render.JSON(w, r, msg)
		return
	}

	threadId, err := strconv.ParseInt(threadStr, 10, 64)
	if err != nil {
		const msg = "invalid thread id"
		log.Error(msg, sl.Err(err))
		render.JSON(w, r, msg)
		return
	}

	posts, err := h.service.List(r.Context(), boardKey, threadId)
	if errors.Is(err, customErrors.ErrNotFound) {
		const msg = "not found"
		log.Error(msg)
		render.JSON(w, r, msg)
		return
	}
	if err != nil {
		const msg = "internal error"
		log.Error(msg, sl.Err(err))
		render.JSON(w, r, msg)
		return
	}

	render.JSON(w, r, posts)
}
