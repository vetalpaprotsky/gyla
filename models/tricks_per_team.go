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
		if !t.IsCompleted() {
			continue
		}

		winTeam := relation.getTeam(t.Winner())

		if winTeam == relation.Team1 {
			result.Tricks1 += 1
		} else {
			result.Tricks2 += 1
		}
	}

	return result
}
