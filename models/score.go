package models

// TODO: This might be a better approach?
// type Score struct {
// 	team1  Team
// 	score1 int
//
// 	team2  Team
// 	score2 int
// }

type Score map[Team]int

func newScore(game Game) Score {
	score := make(Score)

	score[game.Relation.Team1] = 0
	score[game.Relation.Team2] = 0

	for _, round := range game.Rounds {
		winnerTeam := round.winnerTeam()

		if winnerTeam == Team("") {
			continue
		}

		if winnerTeam == game.Relation.getTeam(round.starter) {
			score[winnerTeam] += 12
		}
	}

	return score
}

func (s Score) isGameCompleted() bool {
	for _, points := range s {
		if points >= 60 {
			return true
		}
	}

	return false
}

func (s Score) winnerTeam() Team {
	var team Team
	var points int

	for k, v := range s {
		if v > points {
			team = k
			points = v
		}
	}

	return team
}
