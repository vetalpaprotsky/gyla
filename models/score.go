package models

type Score struct {
	team1 int
	team2 int
}

func NewScore(rounds []Round) Score {
	// TODO: Calculate score based on rounds.
	return Score{}
}
