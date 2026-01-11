package game

// TODO: We could store every round history here.
// We might even want to rename this to matchStats?
type score struct {
	team1   Team
	points1 int
	team2   Team
	points2 int
}

// TODO: When loser team has no tricks, or has one trick, the number of added
// points must be different.
func newScore(m match) score {
	score := score{
		team1: m.plrsRel.team1,
		team2: m.plrsRel.team2,
	}

	for _, round := range m.rounds {
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

func (s score) isMatchCompleted() bool {
	return s.points1 >= 60 || s.points2 >= 60
}

func (s score) winTeam() Team {
	if !s.isMatchCompleted() {
		return Team("")
	}

	if s.points1 > s.points2 {
		return s.team1
	} else {
		return s.team2
	}
}
