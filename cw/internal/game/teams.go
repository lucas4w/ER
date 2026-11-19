package game

type Team string

const (
	Mafia Team = "mafia"
	Civil Team = "civil"
)

var RoleToTeam = map[Role]Team{
	Assassin:   Mafia,
	Apprentice: Mafia,
	Silencer:   Mafia,
	Paralyzer:  Mafia,
	Papparazzi: Mafia,
	Demon:      Mafia,
	Psycopath:  Mafia,
	Oracle:     Mafia,
	Judge:      Civil,
	Police:     Civil,
	Angel:      Civil,
	Detective:  Civil,
	Villager:   Civil,
	Spirit:     Civil,
	Squire:     Civil,
	Fairy:      Civil,
}

func GetTeamByRole(role Role) Team {
	if team, ok := RoleToTeam[role]; ok {
		return team
	}
	return ""
}

func CountAlivePlayersByTeam(game *Game) map[Team]int {
	counts := make(map[Team]int)
	for _, player := range game.Players {
		if player != nil && player.Alive {
			counts[player.Team]++
		}
	}
	return counts
}
