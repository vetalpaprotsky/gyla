package game

import "fmt"

type MatchStats struct {
	Team1   Team
	Team2   Team
	Points1 int
	Points2 int
	WinTeam Team
}

// TODO: When loser team has no tricks, or has one trick, the number of added
// points must be different.
func newMatchStats(m match) MatchStats {
	stats := MatchStats{
		Team1: m.table.Team1,
		Team2: m.table.Team2,
	}

	for _, round := range m.rounds {
		winTeam := round.winTeam()
		if winTeam.IsZero() {
			continue
		}

		pointsToAdd := 6
		if winTeam != round.starterTeam() {
			pointsToAdd = 12
		}

		switch winTeam {
		case stats.Team1:
			stats.Points1 += pointsToAdd
		case stats.Team2:
			stats.Points2 += pointsToAdd
		default:
			panic(fmt.Sprintf("unknown team: %v", winTeam))
		}
	}

	if stats.Points1 >= 60 {
		stats.WinTeam = stats.Team1
	} else if stats.Points2 >= 60 {
		stats.WinTeam = stats.Team2
	}

	return stats
}

func (s MatchStats) isMatchCompleted() bool {
	return !s.WinTeam.IsZero()
}
