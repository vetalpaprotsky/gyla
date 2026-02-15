package game

import (
	"fmt"
)

type Game struct {
	match  match
	ai     ai
	events []GameEvent
}

func NewGame(t1, p1, p3, t2, p2, p4 string) Game {
	team1 := Team(t1)
	player1 := Player(p1)
	player3 := Player(p3)
	team2 := Team(t2)
	player2 := Player(p2)
	player4 := Player(p4)

	game := Game{
		match: match{
			plrsRel: playersRelation{
				team1:   team1,
				player1: player1,
				player3: player3,
				team2:   team2,
				player2: player2,
				player4: player4,
			},
		},
	}

	return game
}

func (g *Game) StartMatch() GameUpdate {
	g.addEvent(matchStartedEvent)
	g.startNextRound()
	g.applyAiActions()

	return g.createGameUpdate()
}

func (g *Game) Apply(action Action) (ActionResult, GameUpdate) {
	actRes := g.apply(action)

	if actRes.Succeeded {
		g.applyAiActions()
	} else {
		return actRes, GameUpdate{}
	}

	return actRes, g.createGameUpdate()
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
	action := g.ai.getAction(g.match)
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

	g.addEvent(roundStartedEvent)
}

func (g *Game) startNextTrick() {
	if err := g.match.startNextTrick(); err != nil {
		panic(err)
	}

	g.addEvent(trickStartedEvent)
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

	g.addEvent(trumpAssignedEvent)
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

	g.addEvent(cardPlayedEvent)
	if g.match.isMatchCompleted {
		g.addEvent(trickCompletedEvent)
		g.addEvent(roundCompletedEvent)
		g.addEvent(matchCompletedEvent)
	} else if g.match.isCurrentRoundCompleted() {
		g.addEvent(trickCompletedEvent)
		g.addEvent(roundCompletedEvent)
		g.startNextRound()
	} else if g.match.isCurrentTrickCompleted() {
		g.addEvent(trickCompletedEvent)
		g.startNextTrick()
	}

	return nil
}

func (g *Game) addEvent(name string) {
	g.events = append(g.events, NewGameEvent(name, g))
}

func (g *Game) createGameUpdate() GameUpdate {
	update := NewGameUpdate(g)

	// It's important to delete events, so they
	// don't end up in the next game update.
	g.events = nil

	return update
}
