package models

type Game struct {
	rounds  []Round
	player1 Player
	player2 Player
	player3 Player
	player4 Player
}

func NewGame(p1, p2, p3, p4 string) *Game {
	player1 := Player{Name: p1}
	player2 := Player{Name: p2}
	player3 := Player{Name: p3}
	player4 := Player{Name: p4}

	player1.leftOpponent = &player2
	player1.teammate = &player3
	player1.rightOpponent = &player4

	player2.leftOpponent = &player3
	player2.teammate = &player4
	player2.rightOpponent = &player1

	player3.leftOpponent = &player4
	player3.teammate = &player1
	player3.rightOpponent = &player2

	player4.leftOpponent = &player1
	player4.teammate = &player2
	player4.rightOpponent = &player3

	team1 := Team{Name: "team1", Player1: &player1, Player2: &player3}
	team2 := Team{Name: "team2", Player1: &player2, Player2: &player4}

	player1.Team = &team1
	player2.Team = &team2
	player3.Team = &team1
	player4.Team = &team2

	return &Game{
		player1: player1,
		player2: player2,
		player3: player3,
		player4: player4,
	}
}

func (game *Game) StartGameLoop(
	stateChangeCallback func(g *Game),
	playerTrumpAssignmentCallback func(p Player, cards []Card) string,
	playerMoveCallback func(p Player, cards []Card) Card,
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
			round.starter,
			round.getHand(round.starter).cards,
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
				card := playerMoveCallback(p, round.availableCardsForMove(p))
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

func (g Game) currentRound() *Round {
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

func (g *Game) startNextRound() (*Round, error) {
	players := []Player{g.player1, g.player2, g.player3, g.player4}
	round, err := newRound(players, g.currentRound())

	if err != nil {
		return nil, err
	}

	g.rounds = append(g.rounds, *round)

	return round, nil
}

// TODO
func NewGameFromJSON() {

}

// TODO
func (g Game) ToJSON() {

}

func (g Game) team1() Team {
	return *g.player1.Team
}

func (g Game) team2() Team {
	return *g.player3.Team
}

func (g Game) score() Score {
	return newScore(g.rounds)
}
