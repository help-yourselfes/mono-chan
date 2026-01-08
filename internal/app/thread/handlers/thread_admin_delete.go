package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/helpyourselfes/mono-chan/internal/app/thread/dto"
	"github.com/helpyourselfes/mono-chan/internal/logger/sl"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
	"github.com/helpyourselfes/mono-chan/internal/router/middleware"
)

func (h *ThreadHandler) DeleteByAdmin(w http.ResponseWriter, r *http.Request) {
	log := middleware.FromContext(r.Context())
	const op = "thread.handlers.adminDelete"
	log = log.With(slog.String("op", op))

	var req dto.DeleteAdminThreadRequest

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		const msg = "failed to decode request"
		log.Error(msg, sl.Err(err))
		render.JSON(w, r, msg)
		return
	}

	err = h.service.DeleteByAdmin(r.Context(), req.BoardKey, req.PostID)
	if errors.Is(err, customErrors.ErrNotFound) {
		const msg = "not found"
		log.Error(msg)
		render.JSON(w, r, msg)
		return
	}
	if errors.Is(err, customErrors.ErrIncorectPassword) {
		const msg = "incorrect password"
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
		const msg = "unknown error"
		log.Error(msg, sl.Err(err))
		render.JSON(w, r, msg)
		return
	}

	render.JSON(w, r, "ok")
}
