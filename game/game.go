package game

import (
	"fmt"
)

type Game struct {
	match  match
	events []Event
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

func (g *Game) StartMatch() {
	g.addEvent(matchStartedEvent)
	g.startNextRound()
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

// TODO: Let's retry only possible errors, the ones that could be introduced
// by invalid method params. Other - panic.
func (g *Game) assignTrumpForCurrentRound(suit Suit, player Player) error {
	if err := g.match.assignTrumpForCurrentRound(suit, player); err != nil {
		switch err.(matchError).error_type {
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

// TODO: Let's retry only possible errors, the ones that could be introduced
// by invalid method params. Other - panic.
func (g *Game) playCard(rank Rank, suit Suit, player Player) error {
	if err := g.match.playCard(rank, suit, player); err != nil {
		switch err.(matchError).error_type {
		case noCurrentRoundError, noCurrentTrickError, tooManyCardsPerTrickError:
			panic(err)
		default:
			return err
		}
	}

	g.addEvent(cardPlayedEvent)

	if g.match.isCurrentTrickCompleted() {
		g.addEvent(trickCompletedEvent)

		if g.match.isCurrentRoundCompleted() {
			g.addEvent(roundCompletedEvent)

			if g.match.isMatchCompleted {
				g.addEvent(matchCompletedEvent)
			} else {
				g.startNextRound()
			}
		} else {
			g.startNextTrick()
		}
	}

	return nil
}

func (g *Game) DoAction(action Action) ActionResult {
	expAct := g.expectedNextAction()
	if expAct.Name == "" {
		return ActionResult{Succeeded: false, ErrorMsg: "No action is expected."}
	}

	if expAct.Name != action.Name || expAct.Player != action.Player {
		msg := fmt.Sprintf(
			"Unexpected action. <%s> action is expected from <%s> player.",
			expAct.Name, expAct.Player,
		)
		return ActionResult{
			ErrorMsg:  msg,
			Succeeded: false,
		}
	}

	var err error
	switch action.Name {
	case AssignTrumpAction:
		err = g.assignTrumpForCurrentRound(action.Suit, action.Player)
	case PlayCardAction:
		err = g.playCard(action.Rank, action.Suit, action.Player)
	default:
		panic(fmt.Sprintf("Unexpected action <%s>", action.Name))
	}

	if err == nil {
		return ActionResult{Succeeded: true}
	} else {
		return ActionResult{Succeeded: false, ErrorMsg: err.Error()}
	}
}

func (g Game) GetEvent() Event {
	// TODO: Remove event from slice.

	return Event{}
}

func (g *Game) addEvent(name string) {
	g.events = append(g.events, Event{name, g.match.createSnapshot()})
}

// TODO: Use match methods, like m.currentRound()
func (g Game) expectedNextAction() ExpectedAction {
	curRound := g.match.currentRound()
	curTrick := curRound.currentTrick()

	if curRound.number == 1 && curTrick == nil {
		return ExpectedAction{AssignTrumpAction, curRound.starter}
	}

	if nextPlayer := curTrick.expectedNextPlayer(); nextPlayer != Player("") {
		return ExpectedAction{PlayCardAction, nextPlayer}
	}

	return ExpectedAction{}
}
