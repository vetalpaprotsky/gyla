package game

type tricksPerTeam struct {
	team1   Team
	team2   Team
	tricks1 int
	tricks2 int
}

func newTricksPerTeam(r round) tricksPerTeam {
	plrsRel := r.plrsRel
	result := tricksPerTeam{
		team1: plrsRel.team1,
		team2: plrsRel.team2,
	}

	for _, t := range r.tricks {
		winner, winnerOk := t.winner()
		if !winnerOk {
			continue
		}

		if plrsRel.getTeam(winner) == plrsRel.team1 {
			result.tricks1 += 1
		} else {
			result.tricks2 += 1
		}
	}

	return result
}
