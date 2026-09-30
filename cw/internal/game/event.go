package game

type EventType string

const (
	PlayerKilled        EventType = "player_killed"
	PlayerSilenced      EventType = "player_silenced"
	PlayerPhotographed  EventType = "player_photographed"
	PlayerParalyzed     EventType = "player_paralyzed"
	PlayerPossessed     EventType = "player_possessed"
	PlayerPsycopathyzed EventType = "player_psycopathyzed"
	PlayerExecuted      EventType = "player_executed"
	PlayerProtected     EventType = "player_protected"
	PlayerInvestigated  EventType = "player_investigated"
	PlayerRevenged      EventType = "player_revenged"
	ActionReflected     EventType = "action_reflected"
	ApprenticeActivated EventType = "apprentice_activated"
	NightStarted        EventType = "night_started"
	DayStarted          EventType = "day_starteds"
	VotingStarted       EventType = "voting_started"
	PlayerVotedOut      EventType = "player_voted_out"
)

type Event struct {
	Type     EventType
	GameID   string
	PlayerID string
	TargetID string
}
