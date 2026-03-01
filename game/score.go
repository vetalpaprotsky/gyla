package game

// TODO: We could store every round history here.
// And most likely we want to rename this to matchStats.
// Let's first image how "Match Stats" table should look like on UI,
// and store everything we need here.
//
// I think we should just store match field hear, and provide any kind of
// needed methods to get some info.
type score struct {
	team1   Team
	team2   Team
	points1 int
	points2 int
}

// TODO: When loser team has no tricks, or has one trick, the number of added
// points must be different.
func newScore(m match) score {
	score := score{
		team1: m.table.team1,
		team2: m.table.team2,
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

		switch winTeam {
		case score.team1:
			score.points1 += pointsToAdd
		case score.team2:
			score.points2 += pointsToAdd
		default:
			panic("unknown team: " + winTeam)
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
