package game

import "math/rand/v2"

type Role string

const (
	Assassin   Role = "assassin"
	Apprentice Role = "apprentice"
	Silencer   Role = "silencer"
	Papparazzi Role = "papparazzi"
	Paralyzer  Role = "paralyzer"
	Demon      Role = "demon"
	Wizard     Role = "wizard"
	Psycopath  Role = "psycopath"

	Judge     Role = "judge"
	Police    Role = "police"
	Angel     Role = "angel"
	Detective Role = "detective"
	Villager  Role = "villager"
	Spirit    Role = "spirit"
	Squire    Role = "squire"
	Fairy     Role = "fairy"
)

func DistributeRoles(game *Game) {
	roles := []Role{
		Assassin, Apprentice, Silencer, Papparazzi, Demon, Paralyzer,
		Judge, Police, Angel, Detective, Villager, Spirit,
	}

	if len(game.Players) >= 14 {
		roles = append(roles, Squire, Wizard)
	}
	if len(game.Players) == 16 {
		roles = append(roles, Fairy, Psycopath)
	}

	rand.Shuffle(len(roles), func(i, j int) {
		roles[i], roles[j] = roles[j], roles[i]
	})

	index := 0
	for _, player := range game.Players {
		if player != nil {
			player.Role = roles[index]
			index++
		}
	}
}
