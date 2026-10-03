package game

import (
	"fmt"
)

type GameEngine struct{}

func NewEngine() *GameEngine {
	return &GameEngine{}
}

func (e *GameEngine) StartGame(game *Game) (Event, error) {
	if !IsValidPlayerCount(game.PlayerCount()) {
		return Event{}, fmt.Errorf("invalid player count")
	}
	game.MaxNights = MaxNights(game.PlayerCount())
	game.Phase = Night
	game.Night = 1
	DistributeRoles(game)
	return Event{Type: GameStarted, GameID: game.ID}, nil
}
func (e *GameEngine) StartNight(game *Game) Event {
	game.Phase = Night
	game.Night += 1

	return Event{Type: NightStarted, GameID: game.ID}
}
func (e *GameEngine) StartDay(game *Game) ([]Event, error) {
	game.Phase = Day

	events, err := e.ResolveActions(game)
	if err != nil {
		return []Event{{}}, err
	}
	return events, nil
}

func (e *GameEngine) EndNight(game *Game) Event {
	return Event{Type: DayStarted, GameID: game.ID}
}

func (e *GameEngine) SubmitAction(game *Game, action Action) error {
	player, err := game.GetPlayer(action.PlayerID)
	if err != nil {
		return fmt.Errorf("player not found")
	}
	err = e.validateAction(game, player, action)
	if err != nil {
		return err
	}
	game.PendingActions = append(game.PendingActions, action)
	return nil
}

func (e *GameEngine) validateAction(game *Game, player *Player, action Action) error {
	if game.Phase != Night {
		return fmt.Errorf("actions can only be submitted during the night phase")
	}
	if !player.IsAlive() {
		return fmt.Errorf("player is dead")
	}
	if player.IsParalyzed() {
		return fmt.Errorf("player is paralyzed and cannot submit actions")
	}
	if !CanMakeAction(game, action) {
		return fmt.Errorf("player's role cannot make this action")
	}
	if game.HasBeenActioned(action.PlayerID) {
		return fmt.Errorf("player has already submitted an action")
	}
	if !IsActionValid(game, action) {
		return fmt.Errorf("invalid target for action")
	}
	return nil
}
func (e *GameEngine) ResolveActions(game *Game) ([]Event, error) {
	game.SortActionsByPriority()
	var allEvents []Event
	for _, action := range game.PendingActions {
		events, err := e.ResolveAction(game, action)
		if err != nil {
			return nil, err
		}
		allEvents = append(allEvents, events...)
	}
	return allEvents, nil
}
func (e *GameEngine) ResolveAction(game *Game, action Action) ([]Event, error) {
	// refatorar essa função para retornar um erro - o motivo de não ser mais válida também pode ser um evento
	if !IsActionValid(game, action) {
		return []Event{}, fmt.Errorf("action is no longer valid")
	}
	switch action.Type {
	case Kill:
		killEvents, err := e.KillPlayer(game, action.TargetID, DeathByAssassin)
		if err != nil {
			return []Event{}, err
		}
		return killEvents, nil
	}
	return []Event{}, nil
}

func (e *GameEngine) KillPlayer(game *Game, playerID string, cause DeathCause) ([]Event, error) {
	player, err := game.GetPlayer(playerID)
	if err != nil {
		return nil, fmt.Errorf("player not found")
	}
	eventType := GetType(cause)
	player.Alive = false
	switch player.Role {
	case Assassin:
		return []Event{
			{Type: eventType, GameID: game.ID, PlayerID: player.Id},
			{Type: ApprenticeActivated, GameID: game.ID},
		}, nil
	default:
		return []Event{
			{Type: eventType, GameID: game.ID, PlayerID: player.Id},
		}, nil
	}

}
func (e *GameEngine) ApplyEffect()          {}
func (e *GameEngine) RemoveExpiredEffects() {}

func (e *GameEngine) StartVoting(game *Game) Event {
	game.Phase = Voting

	return Event{Type: VotingStarted, GameID: game.ID}
}
func (e *GameEngine) CastVote(game *Game, vote Vote) error {
	player, err := game.GetPlayer(vote.PlayerID)
	if err != nil {
		return fmt.Errorf("player not found")
	}
	if !IsVoteValid(game, vote, player) {
		return fmt.Errorf("invalid vote")
	}
	if player.Role == Villager {
		vote.IsVillager = true
	}
	game.Votes = append(game.Votes, vote)
	return nil
}
func (e *GameEngine) CalculateVotes(game *Game) []string {
	results := make(map[string]int)
	for _, vote := range game.Votes {
		results[vote.TargetID]++
	}
	max := 0
	var votedOut []string
	for targetID, count := range results {
		if count > max {
			max = count
			votedOut = []string{targetID}
		} else if count == max {
			votedOut = append(votedOut, targetID)
		}
	}
	return votedOut
}
func (e *GameEngine) ResolveVoting(game *Game) ([]Event, error) {
	result := e.CalculateVotes(game)
	if len(result) > 1 {
		return []Event{{Type: Draw, GameID: game.ID}}, nil
	}
	events, err := e.KillPlayer(game, result[0], DeathByVote)
	if err != nil {
		return []Event{}, err
	}
	return events, nil
}

func (e *GameEngine) CheckWinCondition() {}

func (e *GameEngine) GetPlayerView() {}
