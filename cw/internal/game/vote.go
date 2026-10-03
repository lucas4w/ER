package game

type Vote struct {
	PlayerID   string
	TargetID   string
	IsVillager bool
}

func IsVoteValid(game *Game, vote Vote, player *Player) bool {
	target, err := game.GetPlayer(vote.TargetID)
	if err != nil {
		return false
	}
	if !target.IsAlive() || !player.Alive || player.IsSilenced() || target.IsProtected() || target == player {
		return false
	}
	return true
}
