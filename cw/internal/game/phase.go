package game

type Phase string

const (
	Waiting Phase = "waiting"
	Night   Phase = "night"
	Day     Phase = "day"
	Voting  Phase = "voting"
)
