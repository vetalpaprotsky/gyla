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

	player1.LeftOpponent = &player2
	player1.Teammate = &player3
	player1.RightOpponent = &player4

	player2.LeftOpponent = &player3
	player2.Teammate = &player4
	player2.RightOpponent = &player1

	player3.LeftOpponent = &player4
	player3.Teammate = &player1
	player3.RightOpponent = &player2

	player4.LeftOpponent = &player1
	player4.Teammate = &player2
	player4.RightOpponent = &player3

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

// TODO: These callback functions must be reconsindered for sure.
// They shouldn't force user to implement some game logic.
// They must allow user only perform an action that is required
// from the user: choose trump, make a move.
func (game *Game) StartGameLoop(
	stateChange func(g *Game),
	playerMove func(p Player, availableCardsForMove []Card) Card,
	trumpAssignment func(h Hand) string,
) error {
	// Fresh new game starter
	stateChange(game)

	for roundNum := 1; roundNum < maxPossibleNumberOfRounds; roundNum++ {
		err := game.startRound()

		if err != nil {
			return err
		}

		// Round starter and cards got dealt.
		stateChange(game)

		round := game.CurrentRound()
		trump := trumpAssignment(*round.starterHand())
		round.assignTrump(trump)

		// Round trump assigned
		stateChange(game)

		for trickNum := 1; trickNum <= tricksPerRoundCount; trickNum++ {
			round.startTrick()

			// TODO: Players take moves, they are stored in the trick.
			// Once the trick is complete, we know which team won.
			// Then the next trick starts.

			// 4 moves loop
		}

		// TODO: End loop if there's a winner team
	}

	return nil
}

func (g Game) CurrentRound() *Round {
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

func (g *Game) startRound() error {
	players := []Player{g.player1, g.player2, g.player3, g.player4}
	round, err := NewRound(players, g.CurrentRound())

	if err != nil {
		return err
	}

	g.rounds = append(g.rounds, *round)

	return nil
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
	return NewScore(g.rounds)
}
