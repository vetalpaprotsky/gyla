package game

type RoundStats struct {
	Trumper Player
	WinTeam Team
	Tricks  map[Team]int
	Points  map[Team]int
}

func newRoundStats(r round) RoundStats {
	winTeam := r.winTeam()
	if winTeam.isZero() {
		return RoundStats{}
	}

	tricks := map[Team]int{Team1: 0, Team2: 0}
	for _, t := range r.tricks {
		tricks[t.winner().team()] += 1
	}

	points := map[Team]int{Team1: 0, Team2: 0}

	switch tricks[winTeam] {
	case tricksPerRoundCount:
		points[winTeam] += 24
	case tricksPerRoundCount - 1:
		points[winTeam] += 18
	default:
		points[winTeam] += 6
	}

	if winTeam != r.starterTeam() {
		points[winTeam] += 6
	}

	return RoundStats{
		Trumper: r.starter,
		WinTeam: winTeam,
		Tricks:  tricks,
		Points:  points,
	}
}
