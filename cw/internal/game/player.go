package game

type Player struct {
	Id            string
	Name          string
	Role          Role
	SecondaryRole Role
	Team          Team
	Knowledge     PlayerKnowledge
	Alive         bool
	Silenced      bool
	Paralyzed     bool
	Protected     bool
}

func NewPlayer(id string, name string) *Player {
	return &Player{
		Id:            id,
		Name:          name,
		Role:          "",
		Alive:         true,
		Silenced:      false,
		Paralyzed:     false,
		Protected:     false,
		SecondaryRole: "",
		Knowledge: PlayerKnowledge{
			KnownRoles:      make(map[string]Role),
			TeamComparisons: []TeamComparison{},
		},
	}
}

func (p *Player) IsAlive() bool {
	return p.Alive
}

func (p *Player) IsParalyzed() bool {
	return p.Paralyzed
}

func (p *Player) IsProtected() bool {
	return p.Protected
}

func (p *Player) IsSilenced() bool {
	return p.Silenced
}
