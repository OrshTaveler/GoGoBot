package websocket

import (
	"gogobot/internal/api/shared"

	"github.com/gorilla/websocket"
)

func ConnectGame(conn *websocket.Conn, player *shared.Player) {
	authData := map[string]interface{}{
		"jwt": player.JWT,
	}
	sendMessage(conn, "authenticate", authData)
}

func Move(conn *websocket.Conn, game *shared.Game, move string) {
	sendMessage(conn, "game/move", map[string]interface{}{
		"game_id": game.GameId,
		"move":    move,
	})
}
