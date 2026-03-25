package session

import "gogobot/internal/api/shared"

type Session struct {
	Users []shared.Player
}

func AddUser(session *Session, user shared.Player) {
	session.Users = append(session.Users, user)
}
