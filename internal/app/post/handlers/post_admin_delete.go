package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/helpyourselfes/mono-chan/internal/app/post/dto"
	"github.com/helpyourselfes/mono-chan/internal/logger"
	"github.com/helpyourselfes/mono-chan/internal/logger/sl"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
)

func (h *PostHandler) AdminDelete(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context())
	const op = "post.handlers.adminDelete"
	log.With(slog.String("op", op))
	var req dto.AdminDeletePostRequest

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		const msg = "failed to decode request"
		log.Error(msg, sl.Err(err))
		render.JSON(w, r, msg)
		return
	}

	err = h.service.DeleteByAdmin(r.Context(), req.BoardKey, req.ID)
	if errors.Is(err, customErrors.ErrNotFound) {
		const msg = "not found"
		log.Error(msg)
		render.JSON(w, r, msg)
		return
	}

	if errors.Is(err, customErrors.ErrPostIsRoot) {
		const msg = "make thread deletion request, not post"
		log.Error(msg)
		render.JSON(w, r, msg)
		return
	}
	if err != nil {
		const msg = "failed to decode request"
		log.Error(msg, sl.Err(err))
		render.JSON(w, r, msg)
		return
	}

}
