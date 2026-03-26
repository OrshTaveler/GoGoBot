package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(h.Session.Users); err != nil {
		http.Error(w, "failed to encode users", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(h.Session.Games); err != nil {
		http.Error(w, "failed to encode games", http.StatusInternalServerError)
		return
	}
}
