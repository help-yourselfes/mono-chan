package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/helpyourselfes/mono-chan/internal/app/post/dto"
	"github.com/helpyourselfes/mono-chan/internal/logger/sl"
	"github.com/helpyourselfes/mono-chan/internal/pkg/api/response"
)

func (h *PostHandler) CreatePost(log *slog.Logger) http.HandlerFunc {
	const op = "post.handlers.create"
	log.With(slog.String("op", op))

	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.CreatePostRequest

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request", sl.Err(err))
			render.JSON(w, r, response.Error("failed to decode request"))
			return
		}

		post, err := h.service.Create(r.Context(), &req)
		if err != nil {
			log.Error("failed to create post", sl.Err(err))
			http.Error(
				w,
				fmt.Errorf("failed to create post: %w", err).Error(),
				http.StatusConflict,
			)
			return
		}

		render.DefaultResponder(w, r, post)
	}
}
