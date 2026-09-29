package game

import "fmt"

type GameEngine struct{}

func (e *GameEngine) StartGame(game *Game) error {
	if !IsValidPlayerCount(len(game.Players)) {
		return fmt.Errorf("invalid player count")
	}
	DistributeRoles(game)
	game.Phase = Night
	game.Night = 1
	return nil
}
func (e *GameEngine) StartNight() {}
func (e *GameEngine) StartDay()   {}

func (e *GameEngine) SubmitAction()   {}
func (e *GameEngine) ValidateAction() {}
func (e *GameEngine) ResolveActions() {}
func (e *GameEngine) ResolveAction()  {}

func (e *GameEngine) KillPlayer()           {}
func (e *GameEngine) ResolveDeath()         {}
func (e *GameEngine) ApplyEffect()          {}
func (e *GameEngine) RemoveExpiredEffects() {}

func (e *GameEngine) StartVoting()    {}
func (e *GameEngine) CastVote()       {}
func (e *GameEngine) CalculateVotes() {}
func (e *GameEngine) ResolveVoting()  {}

func (e *GameEngine) CheckWinCondition() {}

func (e *GameEngine) GetPlayerView() {}
