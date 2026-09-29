package game

import "github.com/gofiber/contrib/v3/websocket"

type Player struct {
	Id        string
	Conn      *websocket.Conn
	Name      string
	Role      Role
	Alive     bool
	Silenced  bool
	Paralyzed bool
	Immune    bool
}

func NewPlayer(id string, name string) *Player {
	return &Player{
		Id:        id,
		Conn:      nil,
		Name:      name,
		Role:      "",
		Alive:     true,
		Silenced:  false,
		Paralyzed: false,
		Immune:    false,
	}
}
