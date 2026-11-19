package game

type GameView struct {
	ID              string
	Phase           Phase
	Night           int
	Players         []PlayerView
	TeamComparisons []TeamComparison
}

type PlayerView struct {
	ID    string
	Name  string
	Role  *Role
	Alive bool
	Team  *Team
}

func BuildPlayerView(game *Game, playerID string) (*GameView, error) {
	viewer, err := game.GetPlayer(playerID)
	if err != nil {
		return nil, err
	}
	view := GameView{
		ID:              game.ID,
		Phase:           game.Phase,
		Night:           game.Night,
		TeamComparisons: viewer.Knowledge.TeamComparisons,
		Players:         make([]PlayerView, 0, len(game.Players)),
	}

	for _, player := range game.Players {
		if player == nil {
			continue
		}
		playerView := PlayerView{
			ID:    player.Id,
			Name:  player.Name,
			Alive: player.Alive,
		}

		if player.Id == playerID {
			role := player.Role
			team := player.Team
			playerView.Role = &role
			playerView.Team = &team
		} else if team, ok := game.State.RevealedTeamMembers[player.Id]; ok {
			playerView.Team = &team
		}
		if role, ok := viewer.Knowledge.KnownRoles[player.Id]; ok {
			playerView.Role = &role
			team := GetTeamByRole(role)
			playerView.Team = &team
		}

		if (viewer.Role == Judge || viewer.Role == Police || viewer.Role == Angel) &&
			(player.Role == Judge || player.Role == Police || player.Role == Angel) {

			role := player.Role
			team := player.Team
			playerView.Role = &role
			playerView.Team = &team

		} else if (viewer.Role == Assassin || viewer.Role == Apprentice || viewer.Role == Silencer) &&
			(player.Role == Assassin || player.Role == Apprentice || player.Role == Silencer) {

			role := player.Role
			team := player.Team
			playerView.Role = &role
			playerView.Team = &team
		}
		view.Players = append(view.Players, playerView)
	}
	return &view, nil
}
