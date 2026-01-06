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

	game.addEvent("game_started")
	// Starts first round. After this call, currentRound() must never return nil.
	game.startNextRound()

	return game
}

// It's going to be a complex method :D
func (g *Game) DoAction(action Action) (ActionRejectedError, bool) {
	// TODO: Check for expectedNextAction().
	// TODO: if action is valid, apply it + create proper events.
	// TODO: if trick is completed, start a new one + create proper events.
	// TODO: if round is completed, start a new one + create proper events.
	// TODO: if game is completed, create "game_completed" event.
	// ...

	return ActionRejectedError{}, true
}

func (g Game) GetEvent() Event {
	// TODO: Remove event from slice.

	return Event{}
}

func (g Game) expectedNextAction() ExpectedAction {
	curRound := g.currentRound()
	curTrick := curRound.currentTrick()

	if curRound.number == 1 && curTrick == nil {
		return ExpectedAction{"trump_choice", curRound.starter}
	}

	return ExpectedAction{}
}

func (g Game) currentRound() *round {
	if len(g.rounds) == 0 {
		return nil
	}

	round := &g.rounds[0]
	for i := 1; i < len(g.rounds); i++ {
		if g.rounds[i].number > round.number {
			round = &g.rounds[i]
		}
	}

	return round
}

func (g *Game) startNextRound() {
	var round round
	var err error

	if curRound := g.currentRound(); curRound == nil {
		round = newFirstRound(g.plrsRel)
	} else {
		round, err = newRound(*curRound)
	}

	if err != nil {
		panic(err)
	}

	g.rounds = append(g.rounds, round)
	g.addEvent("round_started")
}

func (g *Game) startNextTrick() {
	if err := g.currentRound().startNextTrick(); err != nil {
		panic(err)
	}

	g.addEvent("trick_started")
}

func (g *Game) addEvent(name string) {
	g.events = append(g.events, Event{name, g.createSnapshot()})
}

func (g Game) createSnapshot() gameSnapshot {
	curRound := g.currentRound()

	if curRound == nil {
		return gameSnapshot{plrsRel: g.plrsRel}
	}

	return gameSnapshot{
		round:   curRound.deepCopy(),
		score:   newScore(g),
		plrsRel: g.plrsRel,
	}
}
