package game

import "fmt"

type Game struct {
	match       match
	matchEvents []MatchEvent
}

func NewGame(p1, p2, p3, p4 Player, t1, t2 Team, ai1, ai2, ai3, ai4 bool) Game {
	game := Game{
		match: match{
			table: newTable(p1, p2, p3, p4, t1, t2, ai1, ai2, ai3, ai4),
		},
	}

	return game
}

func (g *Game) StartMatch() []MatchEvent {
	g.addMatchEvent(MatchStartedEvent)

	g.startNextRound()
	g.applyAiActions()

	return g.clearMatchEvents()
}

func (g *Game) Apply(action Action) (ActionResult, []MatchEvent) {
	actRes := g.apply(action)

	if actRes.Succeeded {
		g.applyAiActions()
	} else {
		return actRes, nil
	}

	return actRes, g.clearMatchEvents()
}

func (g *Game) MatchState() MatchState {
	return g.match.state()
}

func (g *Game) apply(action Action) ActionResult {
	var err error
	switch action.Name {
	case AssignTrumpAction:
		err = g.assignTrump(action.Suit, action.Player)
	case PlayCardAction:
		err = g.playCard(action.Rank, action.Suit, action.Player)
	default:
		err = fmt.Errorf("unexpected action: %s", action.Name)
	}

	if err == nil {
		return ActionResult{Succeeded: true}
	} else {
		return ActionResult{Succeeded: false, ErrorMsg: err.Error()}
	}
}

func (g *Game) applyAiAction() bool {
	action := getAIAction(g.match)
	if action.Name == "" {
		return false
	}

	actRes := g.apply(action)
	if !actRes.Succeeded {
		msg := fmt.Sprintf(
			"ai %s action %s failed: %s",
			action.Player, action.Name, actRes.ErrorMsg,
		)
		panic(msg)
	}

	return true
}

func (g *Game) applyAiActions() {
	for {
		if !g.applyAiAction() {
			return
		}
	}
}

func (g *Game) startNextRound() {
	if err := g.match.startNextRound(); err != nil {
		panic(err)
	}

	g.addMatchEvent(RoundStartedEvent)
}

func (g *Game) startNextTrick() {
	if err := g.match.startNextTrick(); err != nil {
		panic(err)
	}

	g.addMatchEvent(TrickStartedEvent)
}

func (g *Game) assignTrump(suit Suit, player Player) error {
	if err := g.match.assignTrump(suit, player); err != nil {
		switch err.(matchError).code {
		case noCurrentRoundError:
			panic(err)
		default:
			return err
		}
	}

	g.addMatchEvent(TrumpAssignedEvent)
	g.startNextTrick()
	return nil
}

func (g *Game) playCard(rank Rank, suit Suit, player Player) error {
	if err := g.match.playCard(rank, suit, player); err != nil {
		switch err.(matchError).code {
		case noCurrentRoundError, noCurrentTrickError:
			panic(err)
		default:
			return err
		}
	}

	if g.match.isMatchCompleted() {
		g.addMatchEvent(CardPlayedAndMatchCompletedEvent)
	} else if g.match.isCurrentRoundCompleted() {
		g.addMatchEvent(CardPlayedAndRoundCompletedEvent)
		g.startNextRound()
	} else if g.match.isCurrentTrickCompleted() {
		g.addMatchEvent(CardPlayedAndTrickCompletedEvent)
		g.startNextTrick()
	} else {
		g.addMatchEvent(CardPlayedEvent)
	}

	return nil
}

func (g *Game) addMatchEvent(et EventType) {
	g.matchEvents = append(g.matchEvents, newMatchEvent(g.match, et))
}

func (g *Game) clearMatchEvents() []MatchEvent {
	events := g.matchEvents

	g.matchEvents = nil

	return events
}
