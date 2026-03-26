package handlers

import (
	"gogobot/internal/api/rest"
	"net/http"
)

func (h *Handler) StartGame(w http.ResponseWriter, r *http.Request) {
	rest.StartGame(h.Session.Users[0].Token, h.Session.Users[1].UserId)
}
