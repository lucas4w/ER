package game

import (
	"slices"
)

type ActionType string

const (
	Kill         ActionType = "kill"
	Silence      ActionType = "silence"
	Photograph   ActionType = "photograph"
	Paralyze     ActionType = "paralyze"
	Psycopathyze ActionType = "psycopathyze"
	Connect      ActionType = "connect"
	Possess      ActionType = "possess"
	Execute      ActionType = "execute"
	Protect      ActionType = "protect"
	Investigate  ActionType = "investigate"
	Enchant      ActionType = "enchant"
)

type Action struct {
	PlayerID  string
	TargetIDs []string
	Type      ActionType
}

func actionPriority(action Action) int {
	switch action.Type {
	case Enchant:
		return 50
	case Paralyze:
		return 40
	case Protect:
		return 30
	case Execute:
		return 20
	case Kill:
		return 10
	case Psycopathyze:
		return 5
	case Photograph:
		return 0
	case Possess:
		return 0
	}
	return 0
}

var RolePermissions = map[Role][]ActionType{
	Assassin:   {Kill},
	Apprentice: {Kill},
	Silencer:   {Silence},
	Papparazzi: {Photograph},
	Paralyzer:  {Paralyze},
	Demon:      {Possess},
	Oracle:     {Connect},
	Psycopath:  {Psycopathyze},
	Judge:      {Execute},
	Police:     {Execute},
	Angel:      {Protect},
	Detective:  {Investigate},
	Fairy:      {Enchant},
}

func CanMakeAction(game *Game, action Action) bool {
	player, err := game.GetPlayer(action.PlayerID)
	if err != nil {
		return false
	}
	role := player.Role
	allowedActions := RolePermissions[role]
	allowedSecondaryActions := []ActionType{}
	if player.SecondaryRole != "" {
		allowedSecondaryActions = RolePermissions[player.SecondaryRole]
	}
	return slices.Contains(allowedActions, action.Type) || slices.Contains(allowedSecondaryActions, action.Type)
}

func IsActionValid(game *Game, action Action) bool {
	player, err := game.GetPlayer(action.PlayerID)
	if err != nil {
		return false
	}
	for i := range action.TargetIDs {
		if action.TargetIDs[i] == "" {
			continue
		}
		target, err := game.GetPlayer(action.TargetIDs[i])
		if err != nil {
			return false
		}
		if !player.IsAlive() || !target.IsAlive() || (player.IsParalyzed() && action.Type != Possess) || (target.IsProtected() && action.Type != Possess) || target == player {
			return false
		}
	}
	return true
}
