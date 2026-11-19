package game

type DeathCause string

const (
	DeathByAssassin  DeathCause = "assassin"
	DeathByPsycopath DeathCause = "psycopath"
	DeathByVote      DeathCause = "vote"
	DeathByExecution DeathCause = "execution"
	DeathBySpirit    DeathCause = "spirit"
	DeathByFairy     DeathCause = "fairy"
)
