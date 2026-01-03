package models

type TricksPerTeam struct {
	Team1   Team
	Team2   Team
	Tricks1 int
	Tricks2 int
}

func newTricksPerTeam(r Round) TricksPerTeam {
	plrsRel := r.plrsRel
	result := TricksPerTeam{
		Team1: plrsRel.Team1,
		Team2: plrsRel.Team2,
	}

	for _, t := range r.Tricks {
		winner, winnerOk := t.Winner()
		if !winnerOk {
			continue
		}

		if plrsRel.getTeam(winner) == plrsRel.Team1 {
			result.Tricks1 += 1
		} else {
			result.Tricks2 += 1
		}
	}

	return result
}
