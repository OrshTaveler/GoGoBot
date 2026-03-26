package api

import (
	"gogobot/internal/api/shared"
	"gogobot/internal/handlers"
	"gogobot/internal/session"
	"net/http"
)

func Router(session session.Session) {

	h := &handlers.Handler{
		Session: &session,
	}

	http.HandleFunc(shared.AUTH_ENDPOINT, h.AuthHandler)
	http.HandleFunc("/session", h.GetAllUsers)
	http.HandleFunc("/start", h.StartGame)
	http.HandleFunc("/games", h.GetAllGames)

	http.ListenAndServe(shared.PORT, nil)
}
