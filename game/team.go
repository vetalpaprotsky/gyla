package game

type Team int

type TeamInfo struct {
	Team Team
	Name string
}

const (
	Team1 = Team(1)
	Team2 = Team(2)
)

func (t Team) isZero() bool {
	return t == 0
}

func (t Team) opponent() Team {
	switch t {
	case Team1:
		return Team2
	case Team2:
		return Team1
	}

	return Team(0)
}
