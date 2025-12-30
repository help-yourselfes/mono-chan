package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/helpyourselfes/mono-chan/internal/app/post/dto"
	"github.com/helpyourselfes/mono-chan/internal/pkg/api/response"
)

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePostRequest

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		render.JSON(w, r, response.Error("failed to decode request"))
		return
	}

	post, err := h.service.Create(r.Context(), &req)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf("failed to create post: %w", err).Error(),
			http.StatusConflict,
		)
		return
	}

	render.DefaultResponder(w, r, post)
}
