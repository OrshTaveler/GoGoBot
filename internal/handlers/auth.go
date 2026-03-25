package handlers

import (
	"fmt"
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
		fmt.Println(err)
		return
	}

	username, err := rest.GetUserInfo(token)
	if err != nil {
		fmt.Println(err)
		return
	}

	player := shared.Player{
		Token:    token,
		Username: username,
	}

	for i, user := range h.Session.Users {
		if user.Username == player.Username {
			h.Session.Users[i].Token = player.Token
			return
		}
	}

	h.Session.Users = append(h.Session.Users, player)
}
