package game

import "fmt"

type Game struct {
	ID        string
	Night     int
	MaxNights int
	Players   [16]*Player
	Phase     Phase
}

func NewGame(id string) *Game {
	return &Game{
		ID:      id,
		Night:   0,
		Players: [16]*Player{},
		Phase:   Waiting,
	}
}

func (g *Game) AddPlayer(player *Player) error {
	for i := range len(g.Players) {
		if g.Players[i] == nil {
			g.Players[i] = player
			return nil
		}
	}
	return fmt.Errorf("game is full")
}

func (g *Game) RemovePlayer(playerID string) error {
	for i := range len(g.Players) {
		if g.Players[i] != nil && g.Players[i].Id == playerID {
			g.Players[i] = nil
			return nil
		}
	}
	return fmt.Errorf("player not found")
}

func (g *Game) PlayerCount() int {
	count := 0
	for i := range len(g.Players) {
		if g.Players[i] != nil {
			count++
		}
	}
	return count
}
