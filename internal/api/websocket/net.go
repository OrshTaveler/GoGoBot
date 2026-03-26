package websocket

import (
	"encoding/json"
	"gogobot/internal/api/shared"

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
	defer conn.Close()

	return conn, nil
}
