package game

import "fmt"

type Game struct {
	ID             string
	Night          int
	MaxNights      int
	Players        [16]*Player
	PendingActions []Action
	Votes          []Vote
	Phase          Phase
	State          GameState
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

func (g *Game) GetPlayer(id string) (*Player, error) {
	for i := range len(g.Players) {
		if g.Players[i] != nil {
			if g.Players[i].Id == id {
				return g.Players[i], nil
			}
		}
	}
	return nil, fmt.Errorf("player not found")
}

func (g *Game) HasBeenActioned(playerID string) bool {
	for _, action := range g.PendingActions {
		if action.PlayerID == playerID {
			return true
		}
	}
	return false
}

func (g *Game) SortActionsByPriority() {
	for i := 0; i < len(g.PendingActions); i++ {
		for j := i + 1; j < len(g.PendingActions); j++ {
			if actionPriority(g.PendingActions[j]) > actionPriority(g.PendingActions[i]) {
				g.PendingActions[i], g.PendingActions[j] = g.PendingActions[j], g.PendingActions[i]
			}
		}
	}
}
