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
	game.State.ExecutionRightDemon = ExecutionRightDemon
	DistributeRoles(game)
	return Event{Type: GameStarted, GameID: game.ID}, nil
}
func (e *GameEngine) StartNight(game *Game) Event {
	game.Phase = Night
	game.Night += 1

	game.State.ActionsAgainstPolice = []Action{}
	game.State.ActionsAgainstPoliceDemon = []Action{}
	e.RemoveExpiredEffects(game)

	return Event{Type: NightStarted, GameID: game.ID}
}

func (e *GameEngine) EndNight(game *Game) ([]Event, error) {
	game.Phase = Day

	events := []Event{{Type: DayStarted, GameID: game.ID}}

	actionsEvents, err := e.ResolveActions(game)
	if err != nil {
		return []Event{{}}, err
	}
	if len(game.State.ActionsAgainstPolice) > 0 {
		events = append(events, Event{Type: PoliceReport, GameID: game.ID})
	}
	if len(game.State.ActionsAgainstPoliceDemon) > 0 {
		events = append(events, Event{Type: DemonPoliceReport, GameID: game.ID})
	}
	events = append(events, actionsEvents...)

	return events, nil
}

func (e *GameEngine) SubmitAction(game *Game, action Action) (Event, error) {
	player, err := game.GetPlayer(action.PlayerID)
	if err != nil {
		return Event{}, fmt.Errorf("player not found")
	}
	err = e.validateAction(game, player, action)
	if err != nil {
		return Event{}, err
	}
	police, _ := game.GetPlayerByRole(Police)
	if police.Id == action.TargetIDs[0] || (len(action.TargetIDs) > 1 && action.TargetIDs[1] != "" && police.Id == action.TargetIDs[1]) {
		game.State.ActionsAgainstPolice = append(game.State.ActionsAgainstPolice, action)
	}
	if game.State.DemonRole == Police {
		demon, _ := game.GetPlayerByRole(Demon)
		if demon.Id == action.TargetIDs[0] || (len(action.TargetIDs) > 1 && action.TargetIDs[1] != "" && demon.Id == action.TargetIDs[1]) {
			game.State.ActionsAgainstPoliceDemon = append(game.State.ActionsAgainstPoliceDemon, action)
		}
	}
	if action.Type == Execute {
		if player.Role != Demon {
			game.State.ExecutionRight = ExecutionRightUsed
		} else {
			game.State.ExecutionRightDemon = ExecutionRightUsed
		}
	}
	if action.Type == Investigate || action.Type == Connect {
		event, err := e.ResolveAction(game, action)
		if err != nil {
			return Event{}, err
		}
		if len(event) > 0 {
			return event[0], nil
		} else {
			return Event{}, fmt.Errorf("failed to resolve action")
		}
	}
	game.PendingActions[player.Role] = action
	eventType := GetEventTypeByActionType(action.Type)
	return Event{Type: eventType, GameID: game.ID}, nil
}

func (e *GameEngine) validateAction(game *Game, player *Player, action Action) error {
	if game.Phase != Night {
		return fmt.Errorf("actions can only be submitted during the night phase")
	}
	if !player.IsAlive() {
		return fmt.Errorf("player is dead")
	}
	if player.IsParalyzed() && action.Type != Possess {
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
	if action.Type == Possess && game.Night != 1 {
		return fmt.Errorf("demon can only possess on night 1")
	}
	if action.Type == Kill && (player.Role == Apprentice || player.SecondaryRole == Apprentice) && !game.State.ApprenticeActivated {
		return fmt.Errorf("apprentice is not active")
	}
	if action.Type == Execute && player.Role == Police && game.State.ExecutionRight != ExecutionRightPolice {
		return fmt.Errorf("player has no right to execute")
	}
	if action.Type == Execute && (player.Role == Judge || player.Role == Police) && game.State.ExecutionRight == ExecutionRightUsed {
		return fmt.Errorf("only one execution avaiable per game")
	}
	if action.Type == Execute && (player.Role == Demon && (player.SecondaryRole == Judge || player.SecondaryRole == Police) &&
		game.State.ExecutionRightDemon == ExecutionRightUsed) {
		return fmt.Errorf("only one execution avaiable per game")
	}
	return nil
}

func (e *GameEngine) ResolveActions(game *Game) ([]Event, error) {
	actions := game.SortActionsByPriority()
	var allEvents []Event
	for _, action := range actions {
		if action.Type == "" {
			continue
		}
		events, err := e.ResolveAction(game, action)
		if err != nil {
			continue
		}
		allEvents = append(allEvents, events...)
	}
	winEvent, hasWinner := e.CheckWinCondition(game)
	if hasWinner {
		allEvents = append(allEvents, winEvent)
	}
	for k := range game.PendingActions {
		game.PendingActions[k] = Action{}
	}
	return allEvents, nil
}

func (e *GameEngine) ResolveAction(game *Game, action Action) ([]Event, error) {
	if !IsActionValid(game, action) {
		return nil, nil
	}
	target, err := game.GetPlayer(action.TargetIDs[0])
	if err != nil {
		return nil, nil
	}
	player, err := game.GetPlayer(action.PlayerID)
	if err != nil {
		return nil, nil
	}
	var events = []Event{}
	if (target.Role == Squire || target.SecondaryRole == Squire) && (action.Type != Investigate && action.Type != Connect && action.Type != Possess) {
		events = append(events, Event{Type: ActionReflected, GameID: game.ID, ActorID: player.Id})
		var p = player
		player = target
		target = p
	}
	switch action.Type {
	case Kill:
		killEvents := e.KillPlayer(game, target, DeathByAssassin)
		events = append(events, killEvents...)
		return events, nil
	case Execute:
		killEvents := e.KillPlayer(game, target, DeathByExecution)
		events = append(events, killEvents...)
		return events, nil
	case Investigate:
		player.Knowledge.KnownRoles[target.Id] = target.Role
		events = append(events, Event{
			Type:      PlayerInvestigated,
			TargetIDs: []string{target.Id},
			Role:      target.Role,
			GameID:    game.ID,
		})
		return events, nil
	case Photograph:
		events = append(events, Event{
			Type:      PlayerPhotographed,
			TargetIDs: []string{target.Id},
			Team:      target.Team,
			GameID:    game.ID,
		})
		return events, nil
	case Connect:
		target2, err := game.GetPlayer(action.TargetIDs[1])
		if err != nil {
			return nil, nil
		}
		teamComparasion := TeamComparison{PlayerA: target.Id, PlayerB: target2.Id, SameTeam: false}
		if target2.Team == target.Team {
			teamComparasion.SameTeam = true
		}
		player.Knowledge.TeamComparisons = append(player.Knowledge.TeamComparisons, teamComparasion)
		events = append(events, Event{
			Type:            PlayerConnected,
			TargetIDs:       action.TargetIDs,
			GameID:          game.ID,
			TeamComparasion: teamComparasion,
		})
		return events, nil

	case Protect, Paralyze, Silence:
		var effect Effect
		var eventType EventType

		switch action.Type {
		case Protect:
			effect = Protected
			eventType = PlayerProtected
		case Paralyze:
			effect = Paralyzed
			eventType = PlayerParalyzed
		default:
			effect = Silenced
			eventType = PlayerSilenced
		}

		e.ApplyEffect(game, target, effect)
		events = append(events, Event{Type: eventType, TargetIDs: action.TargetIDs, GameID: game.ID})
		return events, nil

	case Possess:
		player.SecondaryRole = target.Role
		player.Knowledge.KnownRoles[target.Id] = target.Role
		game.State.DemonRole = target.Role
		events = append(events, Event{Type: PlayerPossessed, GameID: game.ID, Role: target.Role})
		return events, nil
	}
	return events, nil
}

func (e *GameEngine) KillPlayer(game *Game, player *Player, cause DeathCause) []Event {
	eventType := GetEventTypeByDeathCause(cause)
	player.Alive = false
	playerRole := player.Role
	switch playerRole {
	case Assassin:
		game.State.ApprenticeActivated = true
		return []Event{
			{Type: eventType, GameID: game.ID, TargetIDs: []string{player.Id}},
			{Type: ApprenticeActivated, GameID: game.ID},
		}
	case Judge:
		if game.State.ExecutionRight == ExectutionRightJudge {
			game.State.ExecutionRight = ExecutionRightPolice
		}
		return []Event{
			{Type: eventType, GameID: game.ID, TargetIDs: []string{player.Id}},
		}
	case Spirit:
		return []Event{
			{Type: SpiritKilled, GameID: game.ID, TargetIDs: []string{player.Id}},
			{Type: eventType, GameID: game.ID, TargetIDs: []string{player.Id}},
		}
	default:
		return []Event{
			{Type: eventType, GameID: game.ID, TargetIDs: []string{player.Id}},
		}
	}
}

func (e *GameEngine) SpiritRevenge(game *Game, spiritID string, targetID string) []Event {
	spirit, err := game.GetPlayer(spiritID)
	if err != nil {
		return []Event{}
	}
	target, err := game.GetPlayer(targetID)
	if err != nil {
		return []Event{}
	}
	if spirit.Role != Spirit || (spirit.Role != Demon && spirit.SecondaryRole != Spirit) || spirit.IsAlive() {
		return []Event{}
	}
	killEvents := e.KillPlayer(game, target, DeathBySpirit)
	return append([]Event{{Type: SpiritRevenged, GameID: game.ID, ActorID: spirit.Id, TargetIDs: []string{target.Id}}}, killEvents...)
}

func (e *GameEngine) ApplyEffect(game *Game, player *Player, effect Effect) {
	switch effect {
	case Paralyzed:
		player.Paralyzed = true
	case Silenced:
		player.Silenced = true
	case Protected:
		player.Protected = true
	default:
		return
	}
	game.State.PlayersAffectedByEffects = append(game.State.PlayersAffectedByEffects, player)
}

func (e *GameEngine) RemoveExpiredEffects(game *Game) error {
	for _, player := range game.State.PlayersAffectedByEffects {
		if player.Paralyzed {
			player.Paralyzed = false
		}
		if player.Silenced {
			player.Silenced = false
		}
		if player.Protected {
			player.Protected = false
		}
	}
	game.State.PlayersAffectedByEffects = []*Player{}
	return nil
}

func (e *GameEngine) CastVote(game *Game, vote Vote) error {
	player, err := game.GetPlayer(vote.PlayerID)
	if err != nil {
		return fmt.Errorf("player not found")
	}
	if !IsVoteValid(game, vote, player) || game.Phase != Day {
		return fmt.Errorf("invalid vote")
	}
	vote.IsVillager = false
	if (player.Role == Villager || (player.Role == Demon && player.SecondaryRole == Villager)) && !player.IsParalyzed() {
		vote.IsVillager = true
	}
	game.Votes = append(game.Votes, vote)
	return nil
}

func (e *GameEngine) CalculateVotes(game *Game) []string {
	results := make(map[string]int)
	for _, vote := range game.Votes {
		if vote.IsVillager {
			results[vote.TargetID] += 3
		} else {
			results[vote.TargetID]++
		}
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
	if len(result) == 0 {
		return []Event{{Type: Draw, GameID: game.ID}}, nil
	}
	if len(result) > 1 {
		return []Event{{Type: Draw, GameID: game.ID}}, nil
	}
	player, err := game.GetPlayer(result[0])
	if err != nil {
		return []Event{}, err
	}
	events := e.KillPlayer(game, player, DeathByVote)
	return events, nil
}

func (e *GameEngine) CheckWinCondition(game *Game) (Event, bool) {
	alivePlayers := CountAlivePlayersByTeam(game)
	if game.MaxNights == game.Night {
		return Event{Type: CivilWin, GameID: game.ID}, true
	}
	if alivePlayers[Civil] == 0 {
		return Event{Type: MafiaWin, GameID: game.ID}, true
	}
	return Event{}, false
}

func (e *GameEngine) GetPoliceReport(game *Game, playerID string) []Action {
	police, err := game.GetPlayer(playerID)
	if err != nil {
		return []Action{}
	}
	if !police.IsAlive() || (police.Role != Police || police.SecondaryRole != Police) {
		return []Action{}
	}
	for _, action := range game.State.ActionsAgainstPolice {
		player, err := game.GetPlayer(action.PlayerID)
		if err != nil {
			continue
		}
		if !player.IsAlive() {
			continue
		}
		police.Knowledge.KnownRoles[player.Id] = player.Role
	}
	return game.State.ActionsAgainstPolice
}
