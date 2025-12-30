package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/helpyourselfes/mono-chan/internal/logger/sl"
)

func (h *BoardHandler) GetBoardsList(log *slog.Logger) http.HandlerFunc {
	const op = "boards.handlers.getList"
	log = log.With(slog.String("op", op))

	return func(w http.ResponseWriter, r *http.Request) {

		boards, err := h.service.List(r.Context())
		if err != nil {
			log.Error("failed to get boards list: ", sl.Err(err))
			http.Error(w, "failed to get boards list", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(boards)
	}
}
