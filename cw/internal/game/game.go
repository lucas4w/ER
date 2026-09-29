package game

import "fmt"

type Game struct {
	ID      string
	Night   int
	Players [16]*Player
	Phase   Phase
}

func NewGame(id string) *Game {
	return &Game{
		ID:      id,
		Night:   0,
		Players: [16]*Player{},
		Phase:   Waiting,
	}
}

func (g *Game) AddPlayer(game *Game, player *Player) error {
	for i := range len(game.Players) {
		if game.Players[i] == nil {
			game.Players[i] = player
			return nil
		}
	}
	return fmt.Errorf("game is full")
}

func (g *Game) RemovePlayer(game *Game, playerID string) error {
	for i := range len(game.Players) {
		if game.Players[i] != nil && game.Players[i].Id == playerID {
			game.Players[i] = nil
			return nil
		}
	}
	return fmt.Errorf("player not found")
}
