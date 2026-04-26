package game

import "fmt"

type Game struct {
	rounds       []round
	eventsQueue  []GameEvent
	participants []Participant
	stats        GameStats
}

func NewGame(p1, p2, p3, p4 string, t1, t2 string, ai1, ai2, ai3, ai4 bool) Game {
	game := Game{
		participants: []Participant{
			{
				Player: Player1, Team: Team1,
				PlayerName: p1, TeamName: t1, IsAI: ai1,
			},
			{
				Player: Player2, Team: Team2,
				PlayerName: p2, TeamName: t2, IsAI: ai2,
			},
			{
				Player: Player3, Team: Team1,
				PlayerName: p3, TeamName: t1, IsAI: ai3,
			},
			{
				Player: Player4, Team: Team2,
				PlayerName: p4, TeamName: t2, IsAI: ai4,
			},
		},
	}

	return game
}

func (g *Game) Start() ([]GameEvent, error) {
	if !g.currentRound().isZero() {
		return nil, newGameAlreadyStartedError()
	}

	g.enqueueEvent(GameStartedEvent)
	g.startNextRound()
	applyAIActions(g)

	return g.dequeueEvents(), nil
}

func (g *Game) Apply(action Action) ([]GameEvent, error) {
	err := g.apply(action)

	if err != nil {
		return nil, err
	} else {
		applyAIActions(g)
	}

	return g.dequeueEvents(), nil
}

func (g *Game) apply(action Action) error {
	next := g.nextAction()

	if next.isZero() {
		return newNoActionExpectedError()
	}

	if action.Name != next.Name {
		return newUnexpectedActionError(action.Name, next.Name)
	}

	if action.Player != next.Player {
		return newUnexpectedPlayerError(action.Player, next.Player)
	}

	switch action.Name {
	case AssignTrumpAction:
		return g.assignTrump(action.Suit, action.Player)
	case PlayCardAction:
		return g.playCard(action.Rank, action.Suit, action.Player)
	default:
		panic(fmt.Sprintf("unknown action: %s", action.Name))
	}
}

func (g *Game) startNextRound() {
	if g.isCompleted() {
		panic(newGameCompletedError())
	}

	var round round
	var err error

	if curRound := g.currentRound(); curRound.isZero() {
		round = newFirstRound()
	} else {
		round, err = newRound(curRound)
	}

	if err != nil {
		panic(err)
	}

	g.rounds = append(g.rounds, round)
	g.enqueueEvent(RoundStartedEvent)
}

func (g *Game) startNextTrick() {
	if g.isCompleted() {
		panic(newGameCompletedError())
	}

	if err := g.currentRoundPtr().startNextTrick(); err != nil {
		panic(err)
	}

	g.enqueueEvent(TrickStartedEvent)
}

func (g *Game) playCard(rank Rank, suit Suit, player Player) error {
	curRound := g.currentRoundPtr()
	if err := curRound.playCard(rank, suit, player); err != nil {
		return err
	}

	if curRound.isCompleted() {
		g.recalcStats()
	}

	if g.isCompleted() {
		g.enqueueEvent(CardPlayedAndGameCompletedEvent)
	} else if curRound.isCompleted() {
		g.enqueueEvent(CardPlayedAndRoundCompletedEvent)
		g.startNextRound()
	} else if curRound.currentTrick().isCompleted() {
		g.enqueueEvent(CardPlayedAndTrickCompletedEvent)
		g.startNextTrick()
	} else {
		g.enqueueEvent(CardPlayedEvent)
	}

	applyAIActions(g)

	return nil
}

func (g *Game) assignTrump(suit Suit, player Player) error {
	if err := g.currentRoundPtr().assignTrump(suit, player); err != nil {
		return err
	}

	g.enqueueEvent(TrumpAssignedEvent)
	g.startNextTrick()
	applyAIActions(g)

	return nil
}

func (g *Game) recalcStats() {
	g.stats = newGameStats(*g)
}

func (g *Game) enqueueEvent(et EventType) {
	g.eventsQueue = append(
		g.eventsQueue,
		GameEvent{EventType: et, GameState: g.State()},
	)
}

func (g *Game) dequeueEvents() []GameEvent {
	events := g.eventsQueue

	g.eventsQueue = nil

	return events
}

func (g Game) State() GameState {
	return newGameState(g)
}

func (g Game) currentRound() round {
	if len(g.rounds) == 0 {
		return round{}
	}

	return g.rounds[len(g.rounds)-1]
}

func (g Game) currentRoundPtr() *round {
	if len(g.rounds) == 0 {
		panic("no current round")
	}

	return &g.rounds[len(g.rounds)-1]
}

func (g Game) isCompleted() bool {
	return !g.stats.WinTeam.isZero()
}

func (g Game) getParticipant(p Player) Participant {
	for _, participant := range g.participants {
		if participant.Player == p {
			return participant
		}
	}

	return Participant{}
}

func (g Game) nextAction() NextAction {
	curRound := g.currentRound()
	if curRound.isZero() {
		return NextAction{}
	}

	if !curRound.isTrumpAssigned() {
		return NextAction{
			Player: curRound.starter,
			Name:   AssignTrumpAction,
		}
	}

	curTrick := curRound.currentTrick()
	if curTrick.isZero() {
		return NextAction{}
	}

	nextPlayer := curTrick.expectedNextPlayer()
	if !nextPlayer.isZero() {
		return NextAction{
			Player: nextPlayer,
			Name:   PlayCardAction,
		}
	}

	return NextAction{}
}
