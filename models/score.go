package models

type Score struct {
	Team1   Team
	Points1 int
	Team2   Team
	Points2 int
}

// TODO: When loser team has no tricks, or has one trick, the number of added
// points must be different.
func newScore(g Game) Score {
	score := Score{
		Team1: g.Relation.Team1,
		Team2: g.Relation.Team2,
	}

	for _, round := range g.Rounds {
		winnerTeam := round.winnerTeam()

		if winnerTeam == Team("") {
			continue
		}

		pointsToAdd := 6
		if winnerTeam != round.starterTeam() {
			pointsToAdd = 12
		}

		if winnerTeam == score.Team1 {
			score.Points1 += pointsToAdd
		} else {
			score.Points2 += pointsToAdd
		}
	}

	return score
}

func (s Score) isGameCompleted() bool {
	return s.Points1 >= 60 || s.Points2 >= 60
}

func (s Score) winnerTeam() Team {
	if !s.isGameCompleted() {
		return Team("")
	}

	if s.Points1 > s.Points2 {
		return s.Team1
	} else {
		return s.Team2
	}
}
