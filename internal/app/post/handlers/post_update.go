package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/helpyourselfes/mono-chan/internal/app/post/dto"
	"github.com/helpyourselfes/mono-chan/internal/logger"
	"github.com/helpyourselfes/mono-chan/internal/logger/sl"
	"github.com/helpyourselfes/mono-chan/internal/pkg/api/response"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
)

func (h *PostHandler) Update(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context())
	const op = "post.handlers.update"
	log.With(slog.String("op", op))
	var req dto.UpdatePostRequest

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		const msg = "failed to decode request"
		log.Error(msg, sl.Err(err))
		render.JSON(w, r, msg)
		return
	}

	err = h.service.Update(r.Context(), &req)
	if errors.Is(err, customErrors.ErrNotFound) {
		const msg = "not found"
		log.Error(msg)
		render.JSON(w, r, msg)
		return
	}
	if errors.Is(err, customErrors.ErrIncorectPassword) {
		const msg = "password is incorect"
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
	if err != nil {
		const msg = "failed to decode request"
		log.Error(msg, sl.Err(err))
		render.JSON(w, r, msg)
		return
	}

	render.JSON(w, r, response.OK())
}
