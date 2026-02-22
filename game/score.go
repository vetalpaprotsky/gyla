package game

// TODO: We could store every round history here.
// And most likely we want to rename this to matchStats.
// Let's first image how "Match Stats" table should look like on UI,
// and store everything we need here.
//
// I think we should just store match field hear, and provide any kind of
// needed methods to get some info.
type score struct {
	points map[Team]int
}

// TODO: When loser team has no tricks, or has one trick, the number of added
// points must be different.
func newScore(m match) score {
	teams := m.table.getTeams()
	score := score{
		points: map[Team]int{
			teams[0]: 0,
			teams[1]: 0,
		},
	}

	for _, round := range m.rounds {
		winTeam := round.winTeam()
		if winTeam == Team("") {
			continue
		}

		pointsToAdd := 6
		if winTeam != round.starterTeam() {
			pointsToAdd = 12
		}

		score.points[winTeam] += pointsToAdd
	}

	return score
}

func (s score) isMatchCompleted() bool {
	for _, pts := range s.points {
		if pts >= 60 {
			return true
		}
	}
	return false
}

func (s score) winTeam() Team {
	if !s.isMatchCompleted() {
		return Team("")
	}

	var winner Team
	maxPoints := 0
	for team, pts := range s.points {
		if pts > maxPoints {
			maxPoints = pts
			winner = team
		}
	}
	return winner
}
