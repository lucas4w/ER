package game

func MaxNights(playerCount int) int {
	switch playerCount {
	case 12:
		return 5
	case 14:
		return 6
	case 16:
		return 7
	default:
		return 0
	}
}

func IsValidPlayerCount(playerCount int) bool {
	switch playerCount {
	case 12, 14, 16:
		return true
	default:
		return false
	}
}
