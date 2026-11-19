package game

type GameState struct {
	ExecutionRight            ExecutionRight
	ExecutionRightDemon       ExecutionRight
	ApprenticeActivated       bool
	PlayersAffectedByEffects  []*Player
	RevealedTeamMembers       map[string]Team
	ActionsAgainstPolice      []Action
	ActionsAgainstPoliceDemon []Action
	DemonRole                 Role
}
