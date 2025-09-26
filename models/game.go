package models

type Game struct {
	rounds []Round
	team1  *Team
	team2  *Team
}

func (g Game) player1() Player {
	return *g.team1.Player1
}

func (g Game) player2() Player {
	return *g.team2.Player1
}

func (g Game) player3() Player {
	return *g.team1.Player2
}

func (g Game) player4() Player {
	return *g.team2.Player2
}

func (g Game) score() Score {
	return NewScore(g.rounds)
}
