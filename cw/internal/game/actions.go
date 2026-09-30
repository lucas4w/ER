package game

import "slices"

type ActionType string

const (
	Kill         ActionType = "kill"
	Silence      ActionType = "silence"
	Photograph   ActionType = "photograph"
	Paralyze     ActionType = "paralyze"
	Psycopathyze ActionType = "psycopathyze"
	Possess      ActionType = "possess"
	Execute      ActionType = "execute"
	Protect      ActionType = "protect"
	Investigate  ActionType = "investigate"
	Enchant      ActionType = "enchant"
)

type Action struct {
	PlayerID string
	TargetID string
	Type     ActionType
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
	Demon:      {Possess, Kill, Silence, Photograph, Paralyze, Execute, Protect, Investigate, Enchant, Psycopathyze},
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
	return slices.Contains(allowedActions, action.Type)
}

func IsValid(game *Game, action Action) bool {
	player, err := game.GetPlayer(action.PlayerID)
	if err != nil {
		return false
	}
	target, err := game.GetPlayer(action.TargetID)
	if err != nil {
		return false
	}
	if !target.IsAlive() || !player.IsAlive() || player.IsParalyzed() || (target.IsProtected() && action.Type != Possess) {
		return false
	}
	return true
}
