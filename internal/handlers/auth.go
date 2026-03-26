package handlers

import (
	"gogobot/internal/api/rest"
	"gogobot/internal/api/shared"
	"gogobot/internal/session"
	"net/http"
)

type Handler struct {
	Session *session.Session
}

func (h *Handler) AuthHandler(w http.ResponseWriter, r *http.Request) {
	token, err := rest.GetAccessToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username, userid, err := rest.GetUserInfo(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jwt, err := rest.GetJWT(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	player := shared.Player{
		Token:    token,
		Username: username,
		JWT:      jwt,
		UserId:   userid,
	}

	for i, user := range h.Session.Users {
		if user.Username == player.Username {
			h.Session.Users[i].Token = player.Token
			h.Session.Users[i].JWT = player.JWT

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("updated"))
			return
		}
	}

	h.Session.Users = append(h.Session.Users, player)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
