package game

type EventType string

const (
	AssassinActed         EventType = "assassin_acted"
	SilencerActed         EventType = "silencer_acted"
	PapparazziActed       EventType = "papparazzi_acted"
	ParalyzerActed        EventType = "paralyzer_acted"
	DemonActed            EventType = "demon_acted"
	OracleActed           EventType = "oracle_acted"
	PsycopathActed        EventType = "psycopath_acted"
	JudgeActed            EventType = "judge_acted"
	PoliceActed           EventType = "police_acted"
	AngelActed            EventType = "angel_acted"
	DetectiveActed        EventType = "detective_acted"
	FairyActed            EventType = "fairy_acted"
	PlayerKilled          EventType = "player_killed"
	PlayerSilenced        EventType = "player_silenced"
	PlayerPhotographed    EventType = "player_photographed"
	PlayerParalyzed       EventType = "player_paralyzed"
	PlayerPossessed       EventType = "player_possessed"
	PlayerPsycopathyzed   EventType = "player_psycopathyzed"
	PlayerConnected       EventType = "player_connected"
	PlayerExecuted        EventType = "player_executed"
	PlayerProtected       EventType = "player_protected"
	PlayerInvestigated    EventType = "player_investigated"
	PlayerEnchanted       EventType = "player_enchanted"
	PlayerEnchantedKilled EventType = "player_enchanted_killed"
	SpiritRevenged        EventType = "spirit_revenged"
	SpiritKilled          EventType = "spirit_killed"
	PoliceReport          EventType = "police_report"
	DemonPoliceReport     EventType = "demon_police_report"
	ActionReflected       EventType = "action_reflected"
	ApprenticeActivated   EventType = "apprentice_activated"
	GameStarted           EventType = "game_started"
	NightStarted          EventType = "night_started"
	DayStarted            EventType = "day_started"
	VotingStarted         EventType = "voting_started"
	PlayerVotedOut        EventType = "player_voted_out"
	Draw                  EventType = "draw"
	MafiaWin              EventType = "mafia_win"
	CivilWin              EventType = "civil_win"
)

type Event struct {
	Type            EventType
	GameID          string
	ActorID         string
	TargetIDs       []string
	Role            Role
	Team            Team
	TeamComparasion TeamComparison
}

func GetEventTypeByDeathCause(cause DeathCause) EventType {
	switch cause {
	case DeathByAssassin:
		return PlayerKilled
	case DeathByExecution:
		return PlayerExecuted
	case DeathBySpirit:
		return SpiritRevenged
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

func GetEventTypeByActionType(actionType ActionType) EventType {
	switch actionType {
	case Kill:
		return AssassinActed
	case Execute:
		return JudgeActed
	case Investigate:
		return DetectiveActed
	case Photograph:
		return PapparazziActed
	case Connect:
		return OracleActed
	case Protect:
		return AngelActed
	case Paralyze:
		return ParalyzerActed
	case Silence:
		return SilencerActed
	case Possess:
		return DemonActed
	case Psycopathyze:
		return PsycopathActed
	case Enchant:
		return FairyActed
	default:
		return ""
	}
}
