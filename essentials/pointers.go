package essentials

type Team struct {
	name        string
	yearFounded int
	stadium     string
	owner       string
}

type Player struct {
	name     string
	height   int
	weight   int
	position string
	team     Team
}

func DefaultInit() Player {
	return Player{
		name:     "default",
		height:   180,
		weight:   60,
		position: "PG",
		team: Team{
			name:        "ABC",
			yearFounded: 1800,
			stadium:     "default",
			owner:       "default",
		},
	}
}

// one gotcha with golang is that the * in the receiver fn indicates address
// the above is called type description
// so to access the real object, another * is needed in line 36
func (p *Player) UpdatePosition(pos string) {
	(*p).position = pos
}

// Value types - int, float, bool, string, struct
// Ref types - slices, maps, channels, pointers, functions
