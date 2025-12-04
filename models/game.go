package models

type Game struct {
	Rounds   []Round
	Relation PlayersRelation
}

func NewGame(t1, p1, p3, t2, p2, p4 string) *Game {
	team1 := Team(t1)
	player1 := Player(p1)
	player3 := Player(p3)
	team2 := Team(t2)
	player2 := Player(p2)
	player4 := Player(p4)

	return &Game{
		Relation: PlayersRelation{
			Team1:   team1,
			Player1: player1,
			Player3: player3,
			Team2:   team2,
			Player2: player2,
			Player4: player4,
		},
	}
}

// TODO: We might want to add more callbacks to make things easier for the
// client to understand what happed. Things like:
// 1. Move take.
// 2. Round started.
// 3. Trick completed.
// 4. Game completed.
// 5. ...
//
// That way client won't need to dig in the game state to understand
// what happened. But, let's see. For now let's keep things as they are.
func (g *Game) StartGameLoop(
	stateChangeCallback func(g *Game),
	playerTrumpAssignmentCallback func(p Player, cards []Card) Suit,
	playerMoveCallback func(p Player, cards []Card) Card,
) error {
	// Fresh new game starter.
	stateChangeCallback(g)

	for range maxPossibleNumberOfRounds {
		round, err := g.startNextRound()

		if err != nil {
			return err
		}

		// Round starter and cards got dealt.
		stateChangeCallback(g)

		trump := playerTrumpAssignmentCallback(
			round.starter,
			round.getHand(round.starter).Cards,
		)

		round.assignTrump(trump)

		// Round trump assigned.
		stateChangeCallback(g)

		for range tricksPerRoundCount {
			trick := round.startNextTrick()
			starter := trick.starter

			// New trick started.
			stateChangeCallback(g)

			for p := starter; ; p = g.Relation.getLeftOpponent(p) {
				card := playerMoveCallback(p, round.availableCardsForMove(p))
				round.takeMove(p, card)

				// Move taken.
				stateChangeCallback(g)

				if g.Relation.getLeftOpponent(p) == starter {
					break
				}
			}
		}

		if g.Score().isGameCompleted() {
			// Game completed.
			stateChangeCallback(g)
			return nil
		}
	}

	// Game completed. Max possible number of rounds played.
	stateChangeCallback(g)
	return nil
}

func (g *Game) CurrentRound() *Round {
	if len(g.Rounds) == 0 {
		return nil
	}

	curRound := &g.Rounds[0]
	for i := 1; i < len(g.Rounds); i++ {
		if g.Rounds[i].Number > curRound.Number {
			curRound = &g.Rounds[i]
		}
	}

	return curRound
}

func (g *Game) startNextRound() (*Round, error) {
	round, err := newRound(g.Relation, g.CurrentRound())

	if err != nil {
		return nil, err
	}

	g.Rounds = append(g.Rounds, *round)

	return &g.Rounds[len(g.Rounds)-1], nil
}

func (g *Game) Score() Score {
	return newScore(*g)
}
