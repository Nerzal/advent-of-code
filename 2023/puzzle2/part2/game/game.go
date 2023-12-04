package game

const (
	MaxRed   = 12
	MaxGreen = 13
	MaxBlue  = 14
)

type Game struct {
	ID     int
	Valid  bool
	Rounds []*Round
}

func (g *Game) Validate() {
	for _, round := range g.Rounds {
		if !round.Valid {
			g.Valid = false
			return
		}
	}

	g.Valid = true
}

func (g *Game) GetMax() (red, green, blue int) {
	for _, round := range g.Rounds {
		if round.Red > red {
			red = round.Red
		}

		if round.Green > green {
			green = round.Green
		}

		if round.Blue > blue {
			blue = round.Blue
		}
	}

	return
}

type Round struct {
	Valid bool
	Red   int
	Blue  int
	Green int
}

func (r *Round) Validate() {
	if r.Red > MaxRed {
		r.Valid = false
		return
	}

	if r.Blue > MaxBlue {
		r.Valid = false
		return
	}

	if r.Green > MaxGreen {
		r.Valid = false
		return
	}

	r.Valid = true
}
