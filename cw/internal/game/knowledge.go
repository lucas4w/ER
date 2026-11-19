package game

type PlayerKnowledge struct {
	KnownRoles      map[string]Role
	TeamComparisons []TeamComparison
}

type TeamComparison struct {
	PlayerA  string
	PlayerB  string
	SameTeam bool
}
