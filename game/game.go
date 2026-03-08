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
		applyAIActions(g)
	} else {
		return nil, err
	}

	return g.dequeueEvents(), nil
}

func (g *Game) apply(action Action) error {
	var err error
	switch action.Name {
	case AssignTrumpAction:
		err = g.assignTrump(action.Suit, action.Player)
	case PlayCardAction:
		err = g.playCard(action.Rank, action.Suit, action.Player)
	default:
		err = fmt.Errorf("unexpected action: %s", action.Name)
	}

	if err != nil {
		return err
	}

	return nil
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
	if g.isCompleted() {
		return newGameCompletedError()
	}

	curRound := g.currentRoundPtr()
	if err := curRound.playCard(rank, suit, player); err != nil {
		return err
	}

	g.recalcStats()

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
	if g.isCompleted() {
		return newGameCompletedError()
	}

	if err := g.currentRoundPtr().assignTrump(suit, player); err != nil {
		return err
	}

	g.enqueueEvent(TrumpAssignedEvent)
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
