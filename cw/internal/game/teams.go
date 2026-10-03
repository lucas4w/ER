package game

type Team string

const (
	Mafia Team = "mafia"
	Civil Team = "civil"
)

var TeamRole = map[Team][]Role{
	Mafia: {Assassin, Apprentice, Silencer, Paralyzer, Papparazzi, Demon, Psycopath, Wizard},
	Civil: {Judge, Police, Angel, Detective, Villager, Spirit, Squire, Fairy},
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
