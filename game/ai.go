package game

import (
	"fmt"
	"math/rand/v2"
)

func applyAIActions(g *Game) {
	for {
		if !applyAIAction(g) {
			return
		}
	}
}

func applyAIAction(g *Game) bool {
	action := getAIAction(g)
	if action.isZero() {
		return false
	}

	err := g.apply(action)
	if err != nil {
		msg := fmt.Sprintf(
			"ai %v action %s failed: %s", action.Player, action.Name, err,
		)
		panic(msg)
	}

	return true
}

func getAIAction(g *Game) Action {
	if action := tryToAssignTrump(g); !action.isZero() {
		return action
	}

	if action := tryToPlayCard(g); !action.isZero() {
		return action
	}

	return Action{}
}

func tryToAssignTrump(g *Game) Action {
	curRound := g.currentRound()

	if !curRound.isTrumpAssigned() && g.getParticipant(curRound.trumper()).IsAI {
		return Action{
			Name:   AssignTrumpAction,
			Player: curRound.trumper(),
			Suit:   randomSuit(),
		}
	}

	return Action{}
}

func tryToPlayCard(g *Game) Action {
	curRound := g.currentRound()
	curTrick := curRound.currentTrick()
	if curTrick.isZero() {
		return Action{}
	}

	player := curTrick.expectedNextPlayer()
	if player.isZero() || !g.getParticipant(player).IsAI {
		return Action{}
	}

	hand := curRound.getHand(player)
	playableCards := hand.playableCardsFor(curTrick)
	if len(playableCards) == 0 {
		panic("AI: no playable cards for expected next player")
	}

	card := playableCards[rand.IntN(len(playableCards)-1)]
	return Action{
		Name:   PlayCardAction,
		Player: player,
		Rank:   card.rank,
		Suit:   card.suit,
	}
}

func randomSuit() Suit {
	return allSuits[rand.IntN(4)]
}
