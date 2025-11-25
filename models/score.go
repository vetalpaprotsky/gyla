package models

// TODO: Use PlayersRelation in Game to set team1 and team2 and calculate scores.
// type Score struct {
// 	team1  Team
// 	score1 int

// 	team2  Team
// 	score2 int
// }

type Score map[string]int

func newScore(game Game) Score {
	score := make(Score)

	score[game.Player1.Team.Name] = 0
	score[game.Player1.leftOpponent.Team.Name] = 0

	for _, round := range game.Rounds {
		winnerTeam := round.winnerTeam()

		if winnerTeam == nil {
			continue
		}

		if winnerTeam.Name == round.starter.Team.Name {
			score[winnerTeam.Name] += 12
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

func (s Score) winnerTeam() string {
	// var team Team
	var team string
	var points int

	for k, v := range s {
		if v > points {
			team = k
			points = v
		}
	}

	return team
}
