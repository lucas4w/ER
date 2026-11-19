package game

import (
	"fmt"
	"slices"
)

type Game struct {
	ID             string
	Night          int
	MaxNights      int
	Players        [16]*Player
	PendingActions map[Role]Action
	Votes          []Vote
	Phase          Phase
	State          GameState
}

func NewGame(id string) *Game {
	var pendingActions = map[Role]Action{
		Angel:      {},
		Assassin:   {},
		Detective:  {},
		Demon:      {},
		Judge:      {},
		Oracle:     {},
		Papparazzi: {},
		Paralyzer:  {},
		Silencer:   {},
	}
	return &Game{
		ID:             id,
		Night:          0,
		Players:        [16]*Player{},
		Phase:          Waiting,
		PendingActions: pendingActions,
		State: GameState{
			ExecutionRight:      ExectutionRightJudge,
			ApprenticeActivated: false,
		},
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
	if id == "" {
		return nil, fmt.Errorf("invalid id")
	}
	for i := range len(g.Players) {
		if g.Players[i] != nil {
			if g.Players[i].Id == id {
				return g.Players[i], nil
			}
		}
	}
	return nil, fmt.Errorf("player not found")
}

func (g *Game) GetPlayerByRole(role Role) (*Player, error) {
	for i := range len(g.Players) {
		if g.Players[i] != nil {
			if g.Players[i].Role == role {
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

func (g *Game) SortActionsByPriority() []Action {
	actions := make([]Action, 0, len(g.PendingActions))
	for _, action := range g.PendingActions {
		actions = append(actions, action)
	}
	slices.SortFunc(actions, func(a, b Action) int {
		diff := actionPriority(b) - actionPriority(a)
		if diff != 0 {
			return diff
		}

		demon, _ := g.GetPlayerByRole(Demon)
		if a.PlayerID == demon.Id { // demon never has priority over the original role
			return 1
		}

		return -1
	})
	return actions
}
