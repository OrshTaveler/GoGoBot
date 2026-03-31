package session

import "gogobot/internal/api/shared"

type Session struct {
	Users []shared.Player
	Games []shared.Game
}

func AddUser(session *Session, user shared.Player) {
	session.Users = append(session.Users, user)
}

func GetUserById(session *Session, id int) (*shared.Player, bool) {
	for i := range session.Users {
		if session.Users[i].UserId == float64(id) {
			return &session.Users[i], true
		}
	}
	return nil, false
}

func GetUserByUsername(session *Session, username string) (*shared.Player, bool) {
	for i := range session.Users {
		if session.Users[i].Username == username {
			return &session.Users[i], true
		}
	}
	return nil, false
}

func GetGameByID(session *Session, id int) (*shared.Game, bool) {
	for i := range session.Games {
		if session.Games[i].GameId == id {
			return &session.Games[i], true
		}
	}
	return nil, false
}
