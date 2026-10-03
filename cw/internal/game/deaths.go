package game

type DeathCause string

const (
	DeathByAssassin  DeathCause = "assassin"
	DeathByVote      DeathCause = "vote"
	DeathByExecution DeathCause = "execution"
	DeathBySpirit    DeathCause = "spirit"
	DeathByPsycopath DeathCause = "psycopath"
	DeathByFairy     DeathCause = "fairy"
)
