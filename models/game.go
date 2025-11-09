package models

type Game struct {
	Rounds  []Round
	Player1 Player
	Player2 Player
	Player3 Player
	Player4 Player
}

func NewGame(p1, p2, p3, p4 string) *Game {
	Player1 := Player{Name: p1}
	Player2 := Player{Name: p2}
	Player3 := Player{Name: p3}
	Player4 := Player{Name: p4}

	Player1.leftOpponent = &Player2
	Player1.teammate = &Player3
	Player1.rightOpponent = &Player4

	Player2.leftOpponent = &Player3
	Player2.teammate = &Player4
	Player2.rightOpponent = &Player1

	Player3.leftOpponent = &Player4
	Player3.teammate = &Player1
	Player3.rightOpponent = &Player2

	Player4.leftOpponent = &Player1
	Player4.teammate = &Player2
	Player4.rightOpponent = &Player3

	team1 := Team{Name: "team1", Player1: &Player1, Player2: &Player3}
	team2 := Team{Name: "team2", Player1: &Player2, Player2: &Player4}

	Player1.Team = &team1
	Player2.Team = &team2
	Player3.Team = &team1
	Player4.Team = &team2

	return &Game{
		Player1: Player1,
		Player2: Player2,
		Player3: Player3,
		Player4: Player4,
	}
}

func (game *Game) StartGameLoop(
	stateChangeCallback func(g *Game),
	playerTrumpAssignmentCallback func(player string, cards []Card) string,
	playerMoveCallback func(player string, cards []Card) Card,
) error {
	// Fresh new game starter.
	stateChangeCallback(game)

	for range maxPossibleNumberOfRounds {
		round, err := game.startNextRound()

		if err != nil {
			return err
		}

		// Round starter and cards got dealt.
		stateChangeCallback(game)

		trump := playerTrumpAssignmentCallback(
			round.starter.Name,
			round.getHand(round.starter).Cards,
		)

		round.assignTrump(trump)

		// Round trump assigned.
		stateChangeCallback(game)

		for range tricksPerRoundCount {
			trick := round.startNextTrick()
			starter := trick.starter

			// New trick started.
			stateChangeCallback(game)

			for p := starter; p.Name != starter.Name; p = *p.leftOpponent {
				card := playerMoveCallback(p.Name, round.availableCardsForMove(p))
				round.takeMove(p, card)

				// Move taken.
				stateChangeCallback(game)
			}
		}

		// TODO: End loop if there's a winner team. That can be checked
		// with a help of Score struct.
	}

	return nil
}

func (g Game) CurrentRound() *Round {
	if len(g.Rounds) == 0 {
		return nil
	}

	curRound := &g.Rounds[0]
	for i := 1; i < len(g.Rounds); i++ {
		if g.Rounds[i].number > curRound.number {
			curRound = &g.Rounds[i]
		}
	}

	return curRound
}

func (g *Game) startNextRound() (*Round, error) {
	players := []Player{g.Player1, g.Player2, g.Player3, g.Player4}
	round, err := newRound(players, g.CurrentRound())

	if err != nil {
		return nil, err
	}

	g.Rounds = append(g.Rounds, *round)

	return round, nil
}

// TODO
func NewGameFromJSON() {

}

// TODO
func (g Game) ToJSON() {

}

func (g Game) team1() Team {
	return *g.Player1.Team
}

func (g Game) team2() Team {
	return *g.Player3.Team
}

func (g Game) score() Score {
	return newScore(g.Rounds)
}
