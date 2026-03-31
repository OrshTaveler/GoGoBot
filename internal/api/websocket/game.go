package websocket

import (
	"fmt"
	"gogobot/internal/api/shared"
	"time"

	"github.com/gorilla/websocket"
)

func ConnectGame(conn *websocket.Conn, player *shared.Player, gameID int) error {
	authData := map[string]interface{}{
		"jwt": player.JWT,
	}

	player.Connect = conn
	player.MoveDone = make(chan bool, 1)
	player.LastMove = 0

	if err := sendMessage(conn, "authenticate", authData); err != nil {
		return err
	}

	if _, _, err := conn.ReadMessage(); err != nil {
		return err
	}

	if err := sendMessage(conn, "game/connect", map[string]interface{}{
		"game_id": gameID,
		"chat":    true,
	}); err != nil {
		return err
	}

	go readLoop(conn, player)

	time.Sleep(2 * time.Second)

	return nil
}

func Move(conn *websocket.Conn, game *shared.Game, move string) error {
	player := &game.Player1
	if conn == game.Player2.Connect {
		player = &game.Player2
	}

	err := sendMessage(conn, "game/move", map[string]interface{}{
		"game_id": game.GameId,
		"move":    move,
	})
	if err != nil {
		return err
	}

	select {
	case ok := <-player.MoveDone:
		if !ok {
			return fmt.Errorf("move rejected")
		}
		return nil
	case <-time.After(5 * time.Second):
		return fmt.Errorf("move timeout")
	}
}

func AcceptScore(conn *websocket.Conn, game *shared.Game) error {
	return sendMessage(conn, "game/removed_stones/accept", map[string]interface{}{
		"game_id":          game.GameId,
		"stones":           "",
		"strict_seki_mode": false,
	})
}
