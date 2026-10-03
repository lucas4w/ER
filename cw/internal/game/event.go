package game

type EventType string

const (
	PlayerKilled          EventType = "player_killed"
	PlayerSilenced        EventType = "player_silenced"
	PlayerPhotographed    EventType = "player_photographed"
	PlayerParalyzed       EventType = "player_paralyzed"
	PlayerPossessed       EventType = "player_possessed"
	PlayerPsycopathyzed   EventType = "player_psycopathyzed"
	PlayerExecuted        EventType = "player_executed"
	PlayerProtected       EventType = "player_protected"
	PlayerInvestigated    EventType = "player_investigated"
	PlayerRevenged        EventType = "player_revenged"
	PlayerEnchanted       EventType = "player_enchanted"
	PlayerEnchantedKilled EventType = "player_enchanted_killed"
	ActionReflected       EventType = "action_reflected"
	ApprenticeActivated   EventType = "apprentice_activated"
	GameStarted           EventType = "game_started"
	NightStarted          EventType = "night_started"
	DayStarted            EventType = "day_started"
	VotingStarted         EventType = "voting_started"
	PlayerVotedOut        EventType = "player_voted_out"
	Draw                  EventType = "draw"
)

type Event struct {
	Type     EventType
	GameID   string
	PlayerID string
	TargetID string
}

func GetType(cause DeathCause) EventType {
	switch cause {
	case DeathByAssassin:
		return PlayerKilled
	case DeathByExecution:
		return PlayerExecuted
	case DeathBySpirit:
		return PlayerRevenged
	case DeathByPsycopath:
		return PlayerPsycopathyzed
	case DeathByFairy:
		return PlayerEnchantedKilled
	case DeathByVote:
		return PlayerVotedOut
	default:
		return PlayerKilled
	}
}
