package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/helpyourselfes/mono-chan/internal/app/board/validator"
	"github.com/helpyourselfes/mono-chan/internal/pkg/api/response"
)

func (h *BoardHandler) DeleteBoard(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	if !validator.IsValidKey(key) {
		render.JSON(w, r, response.Error("key is invalid"))
		return
	}

	err := h.service.Delete(r.Context(), key)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf("failed to delete board: %w", err).Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
