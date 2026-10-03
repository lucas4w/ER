package main

import (
	"cw/internal/game"
)

func main() {
	// app := fiber.New()
	// log.Panic(app.Listen(":300"))
	g := game.NewGame("game1")
	players := []*game.Player{
		game.NewPlayer("player1", "Alice"),
		game.NewPlayer("player2", "Bob"),
		game.NewPlayer("player3", "Lucas"),
		game.NewPlayer("player4", "Joao"),
		game.NewPlayer("player5", "Jose"),
		game.NewPlayer("player6", "Maria"),
		game.NewPlayer("player7", "Ana"),
		game.NewPlayer("player8", "Carlos"),
		game.NewPlayer("player9", "Pedro"),
		game.NewPlayer("player10", "Matheus"),
		game.NewPlayer("player11", "Gabriel"),
		game.NewPlayer("player12", "Rafael"),
	}

	for _, player := range players {
		err := g.AddPlayer(g, player)
		if err != nil {
			panic(err)
		}
	}

	game.DistributeRoles(g)

	for _, player := range g.Players {
		if player != nil {
			println(player.Name, " é "+player.Role)
		}
	}
}
