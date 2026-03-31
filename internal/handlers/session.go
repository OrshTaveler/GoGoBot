package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type UserDTO struct {
		Username string `json:"username"`
	}

	users := make([]UserDTO, 0, len(h.Session.Users))
	for _, u := range h.Session.Users {
		users = append(users, UserDTO{
			Username: u.Username,
		})
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type GameDTO struct {
		GameId int      `json:"game_id"`
		P1     string   `json:"player1"`
		P2     string   `json:"player2"`
		Moves  []string `json:"moves"`
	}

	games := make([]GameDTO, 0, len(h.Session.Games))
	for _, g := range h.Session.Games {
		games = append(games, GameDTO{
			GameId: g.GameId,
			P1:     g.Player1.Username,
			P2:     g.Player2.Username,
			Moves:  g.Moves,
		})
	}

	if err := json.NewEncoder(w).Encode(games); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
