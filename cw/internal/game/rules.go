package game

func MaxNights(playerCount int) int {
	switch playerCount {
	case 12:
		return 9
	case 14:
		return 10
	case 16:
		return 11
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
