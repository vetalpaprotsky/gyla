package game

type score struct {
	team1   Team
	points1 int
	team2   Team
	points2 int
}

// TODO: When loser team has no tricks, or has one trick, the number of added
// points must be different.
func newScore(g Game) score {
	score := score{
		team1: g.plrsRel.team1,
		team2: g.plrsRel.team2,
	}

	for _, round := range g.rounds {
		winTeam, winTeamOk := round.winTeam()
		if !winTeamOk {
			continue
		}

		pointsToAdd := 6
		if winTeam != round.starterTeam() {
			pointsToAdd = 12
		}

		if winTeam == score.team1 {
			score.points1 += pointsToAdd
		} else {
			score.points2 += pointsToAdd
		}
	}

	return score
}

func (s score) isGameCompleted() bool {
	return s.points1 >= 60 || s.points2 >= 60
}

func (s score) winTeam() Team {
	if !s.isGameCompleted() {
		return Team("")
	}

	if s.points1 > s.points2 {
		return s.team1
	} else {
		return s.team2
	}
}
