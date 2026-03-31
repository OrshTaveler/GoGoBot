package handlers

import (
	"encoding/json"
	"gogobot/internal/api/rest"
	"gogobot/internal/api/shared"
	"gogobot/internal/api/websocket"
	"gogobot/internal/session"
	"gogobot/internal/utils"
	"net/http"
	"time"
)

func (h *Handler) StartGame(w http.ResponseWriter, r *http.Request) {
	white := r.URL.Query().Get("white")
	black := r.URL.Query().Get("black")
	gamefile := r.URL.Query().Get("game")

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

	h.Play(gamefile, int(gameid))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "game started",
		"game_id": gameid,
	})
}

func (h *Handler) Play(gamefile string, gameid int) {

	moves, err := utils.ParseSGFMoves("./games/" + gamefile)
	if err != nil {
		return
	}

	game, _ := session.GetGameByID(h.Session, gameid)

	game.Moves = moves

	go h.playGame(game, moves)
}

func (h *Handler) playGame(game *shared.Game, moves []string) {
	moves = append(moves, "..")
	moves = append(moves, "..")

	conn1, err := websocket.MakeConnection()
	if err != nil {
		return
	}
	defer conn1.Close()

	conn2, err := websocket.MakeConnection()
	if err != nil {
		return
	}
	defer conn2.Close()

	done := make(chan struct{})
	defer close(done)

	go websocket.StartPing(conn1, done)
	go websocket.StartPing(conn2, done)

	websocket.ConnectGame(conn1, &game.Player1, game.GameId)
	websocket.ConnectGame(conn2, &game.Player2, game.GameId)

	time.Sleep(2 * time.Second)

	for i := range moves {
		if i%2 == 0 {
			websocket.Move(conn1, game, moves[i])
		} else {
			websocket.Move(conn2, game, moves[i])
		}

		time.Sleep(100 * time.Millisecond)
	}

	websocket.AcceptScore(conn1, game)
	websocket.AcceptScore(conn2, game)

	filtered := h.Session.Games[:0]
	for _, g := range h.Session.Games {
		if g.GameId != game.GameId {
			filtered = append(filtered, g)
		}
	}
	h.Session.Games = filtered
}
