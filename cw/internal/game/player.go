package game

import "github.com/gofiber/contrib/v3/websocket"

type Player struct {
	Id            string
	Conn          *websocket.Conn
	Name          string
	Role          Role
	SecondaryRole Role
	Team          Team
	Alive         bool
	Silenced      bool
	Paralyzed     bool
	Protected     bool
}

func NewPlayer(id string, name string) *Player {
	return &Player{
		Id:            id,
		Conn:          nil,
		Name:          name,
		Role:          "",
		Alive:         true,
		Silenced:      false,
		Paralyzed:     false,
		Protected:     false,
		SecondaryRole: "",
	}
}

func (p *Player) IsAlive() bool {
	return p.Alive
}

func (p *Player) IsParalyzed() bool {
	return p.Paralyzed
}

func (p *Player) IsProtected() bool {
	return p.Protected
}

func (p *Player) IsSilenced() bool {
	return p.Silenced
}
