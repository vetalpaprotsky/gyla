package game

type Game struct {
	rounds  []round
	plrsRel playersRelation
	events  []Event
}

func NewGame(t1, p1, p3, t2, p2, p4 string) Game {
	team1 := Team(t1)
	player1 := Player(p1)
	player3 := Player(p3)
	team2 := Team(t2)
	player2 := Player(p2)
	player4 := Player(p4)

	game := Game{
		plrsRel: playersRelation{
			team1:   team1,
			player1: player1,
			player3: player3,
			team2:   team2,
			player2: player2,
			player4: player4,
		},
	}

	return game
}

func (g *Game) MakeAction(action Action) (ActionRejectedError, bool) {
	if g.expectedNextAction() != action.Name {
		// return ActionRejectedError
	}

	// TODO

	return ActionRejectedError{}, true
}

func (g Game) expectedNextAction() string {
	return ""
}

func (g Game) getEvent() {

}

func (g Game) getState() {

}

func (g Game) currentRound() *round {
	if len(g.rounds) == 0 {
		return nil
	}

	curRound := &g.rounds[0]
	for i := 1; i < len(g.rounds); i++ {
		if g.rounds[i].number > curRound.number {
			curRound = &g.rounds[i]
		}
	}

	return curRound
}

func (g *Game) startNextRound() (*round, error) {
	var round round
	var err error

	if curRound := g.currentRound(); curRound == nil {
		round = newFirstRound(g.plrsRel)
	} else {
		round, err = newRound(*curRound)
	}

	if err != nil {
		return nil, err
	}

	g.rounds = append(g.rounds, round)

	return &g.rounds[len(g.rounds)-1], nil
}

func (g Game) score() score {
	return newScore(g)
}
