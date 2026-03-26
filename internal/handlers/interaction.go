package handlers

import (
	"encoding/json"
	"gogobot/internal/api/rest"
	"gogobot/internal/api/shared"
	"gogobot/internal/session"
	"gogobot/internal/utils"
	"net/http"
)

func (h *Handler) StartGame(w http.ResponseWriter, r *http.Request) {
	white := r.URL.Query().Get("white")
	black := r.URL.Query().Get("black")

	if white == "" || black == "" {
		http.Error(w, "missing players", http.StatusBadRequest)
		return
	}

	var player1, player2 *shared.Player

	player1, found1 := session.GetUserByUsername(h.Session, white)
	player2, found2 := session.GetUserByUsername(h.Session, black)

	if !found1 || !found2 {
		http.Error(w, "players not found", http.StatusNotFound)
		return
	}

	challenge, err := rest.SendChallenge(
		player1.Token,
		player2.UserId,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	gameid, err := rest.AcceptChallenge(
		player2.Token,
		int(challenge),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	game := shared.Game{
		Player1: *player1,
		Player2: *player2,
		GameId:  int(gameid),
		Moves:   []string{},
	}

	h.Session.Games = append(h.Session.Games, game)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "game started",
		"game_id": gameid,
	})
}

func (h *Handler) Play(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	moves, _ := utils.ParseSGFMoves("/Users/ubica/GoGoBot/81499963-267-Studenton-regina25.sgf")
	json.NewEncoder(w).Encode(moves)
}
