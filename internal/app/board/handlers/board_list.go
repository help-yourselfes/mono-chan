package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *BoardHandler) GetBoardsList(w http.ResponseWriter, r *http.Request) {

	boards, err := h.service.List(r.Context())
	if err != nil {
		http.Error(w, "failed to get boards list", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(boards)
}
