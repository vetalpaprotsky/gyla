package game

import "fmt"

// TODO: Store info about every round:
// points per team, trumper, tricks per player, tricks per team.
type GameStats struct {
	Team1Points int
	Team2Points int
	WinTeam     Team
}

// TODO: When loser team has no tricks, or has one trick, the number of added
// points must be different.
func newGameStats(g Game) GameStats {
	stats := GameStats{}

	for _, round := range g.rounds {
		winTeam := round.winTeam()
		if winTeam.isZero() {
			continue
		}

		pointsToAdd := 6
		if winTeam != round.starterTeam() {
			pointsToAdd = 12
		}

		switch winTeam {
		case Team1:
			stats.Team2Points += pointsToAdd
		case Team2:
			stats.Team2Points += pointsToAdd
		default:
			panic(fmt.Sprintf("unknown team: %v", winTeam))
		}
	}

	if stats.Team1Points >= 60 {
		stats.WinTeam = Team1
	} else if stats.Team2Points >= 60 {
		stats.WinTeam = Team2
	}

	return stats
}

func (gs GameStats) deepCopy() GameStats {
	// TODO: When we have an array of stats for each round.
	return gs
}
