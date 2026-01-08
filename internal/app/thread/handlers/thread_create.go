package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/helpyourselfes/mono-chan/internal/logger/sl"
	"github.com/helpyourselfes/mono-chan/internal/pkg/api/response"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
	"github.com/helpyourselfes/mono-chan/internal/router/middleware"
)

func (h *ThreadHandler) Create(w http.ResponseWriter, r *http.Request) {
	const op = "threads.handlers.create"
	log := middleware.FromContext(r.Context())
	log = log.With(slog.String("op", op))

	var req createThreadWithPostRequest
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		const msg = "failed to decode request"
		log.Error(msg, sl.Err(err))
		render.JSON(w, r, response.Error(msg))
		return
	}

	thread, err := h.service.Create(r.Context(), &req.CreatePostRequest, &req.CreateThreadRequest)
	if errors.Is(err, customErrors.ErrAlreadyExists) {
		const msg = "thread already exists"
		log.Error(msg)
		render.JSON(w, r, response.Error(msg))
		return
	}
	if err != nil {
		const msg = "internal error"
		log.Error(msg, sl.Err(err))
		render.JSON(w, r, response.Error(msg))
		return
	}

	render.DefaultResponder(w, r, thread)
}
