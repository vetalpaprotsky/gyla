package game

type GameStats struct {
	Rounds  []RoundStats
	Points  map[Team]int
	WinTeam Team
}

func newGameStats(g Game) GameStats {
	stats := GameStats{
		Rounds: make([]RoundStats, 0, len(g.rounds)),
		Points: map[Team]int{Team1: 0, Team2: 0},
	}

	for _, round := range g.rounds {
		rs := newRoundStats(round)
		if rs.WinTeam.isZero() {
			continue
		}

		stats.Points[rs.WinTeam] += rs.Points[rs.WinTeam]
		stats.Rounds = append(stats.Rounds, rs)
	}

	if stats.Points[Team1] >= 60 {
		stats.WinTeam = Team1
	} else if stats.Points[Team2] >= 60 {
		stats.WinTeam = Team2
	}

	return stats
}
