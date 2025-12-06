package models

type TricksPerTeam struct {
	Team1   Team
	Team2   Team
	Tricks1 int
	Tricks2 int
}

func newTricksPerTeam(r Round) TricksPerTeam {
	relation := r.relation
	result := TricksPerTeam{
		Team1: relation.Team1,
		Team2: relation.Team2,
	}

	for _, t := range r.Tricks {
		winner, winnerOk := t.Winner()
		if !winnerOk {
			continue
		}

		if relation.getTeam(winner) == relation.Team1 {
			result.Tricks1 += 1
		} else {
			result.Tricks2 += 1
		}
	}

	return result
}
