package websocket

import (
	"encoding/json"
	"gogobot/internal/api/shared"
	"time"

	"github.com/gorilla/websocket"
)

func sendMessage(conn *websocket.Conn, event string, data map[string]interface{}) error {
	msg := []interface{}{event, data}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	finalMsg := shared.SOCKET_IO_PREFIX + string(msgBytes)

	err = conn.WriteMessage(websocket.TextMessage, []byte(finalMsg))
	if err != nil {
		return err
	}

	return nil
}

func MakeConnection() (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial(shared.OGS_WEBSOCKET_URL, nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func readLoop(conn *websocket.Conn, player *shared.Player) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// ping → pong
		if string(msg) == "2" {
			conn.WriteMessage(websocket.TextMessage, []byte("3"))
			continue
		}

		// pong (можно игнорить)
		if string(msg) == "3" {
			continue
		}

		if len(msg) < 2 || string(msg[:2]) != shared.SOCKET_IO_PREFIX {
			continue
		}

		var data []interface{}
		if err := json.Unmarshal(msg[2:], &data); err != nil {
			continue
		}

		if len(data) < 2 {
			continue
		}

		event, _ := data[0].(string)

		switch event {
		case "game/update":
			select {
			case player.MoveDone <- true:
			default:
			}
		case "error", "game/error":
			select {
			case player.MoveDone <- false:
			default:
			}
		}
	}
}

func StartPing(conn *websocket.Conn, done <-chan struct{}) {
	ticker := time.NewTicker(25 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			conn.WriteMessage(websocket.TextMessage, []byte("2"))
		case <-done:
			return
		}
	}
}
