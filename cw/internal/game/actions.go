package game

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
