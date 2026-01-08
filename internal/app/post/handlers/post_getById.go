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

func (h *PostHandler) GetById(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context())
	const op = "post.handlers.getById"
	log.With(slog.String("op", op))
	boardKey := chi.URLParam(r, "boardKey")
	if boardKey == "" {
		msg := "no board key provided"
		log.Error(msg)
		render.JSON(w, r, msg)
	}

	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		msg := "no id provided"
		log.Error(msg)
		render.JSON(w, r, msg)
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		msg := "invalid id"
		log.Error(msg, sl.Err(err))
		render.JSON(w, r, msg)
	}

	post, err := h.service.GetById(r.Context(), boardKey, id)
	if errors.Is(err, customErrors.ErrNotFound) {
		msg := "not found"
		log.Error(msg)
		render.JSON(w, r, msg)
		return
	}
	if err != nil {
		msg := "internal err"
		log.Error(msg, sl.Err(err))
		render.JSON(w, r, msg)
		return
	}

	render.JSON(w, r, post)

}
