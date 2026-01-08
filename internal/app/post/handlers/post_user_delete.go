package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/helpyourselfes/mono-chan/internal/app/post/dto"
	"github.com/helpyourselfes/mono-chan/internal/logger/sl"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
)

func (h *PostHandler) UserDelete(log *slog.Logger) http.HandlerFunc {
	const op = "post.handlers.userDelete"
	log.With(slog.String("op", op))

	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.UserDeletePostRequest
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			const msg = "failed to decode request"
			log.Error(msg, sl.Err(err))
			render.JSON(w, r, msg)
			return
		}

		err = h.service.DeleteByUser(r.Context(), req.BoardKey, req.ID, req.Password)

		if errors.Is(err, customErrors.ErrIncorectPassword) {
			const msg = "password is incorect"
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
		if errors.Is(err, customErrors.ErrNoPasswordSet) {
			const msg = "no password set"
			log.Error(msg)
			render.JSON(w, r, msg)
			return
		}
		if errors.Is(err, customErrors.ErrNotFound) {
			const msg = "not found"
			log.Error(msg)
			render.JSON(w, r, msg)
			return
		}
		if err != nil {
			const msg = "failed to delete post"
			log.Error(msg, sl.Err(err))
			render.JSON(w, r, msg)
			return
		}
		render.JSON(w, r, "ok")
	}
}
