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
	if action.Name == "" {
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
	if action, ok := tryToAssignTrump(g); ok {
		return action
	}

	if action, ok := tryToPlayCard(g); ok {
		return action
	}

	return Action{}
}

func tryToAssignTrump(g *Game) (Action, bool) {
	curRound := g.mustCurrentRound()

	if !curRound.isTrumpAssigned() && g.getParticipant(curRound.trumper()).IsAI {
		return Action{
			Name:   AssignTrumpAction,
			Player: curRound.trumper(),
			Suit:   randomSuit(),
		}, true
	}

	return Action{}, false
}

func tryToPlayCard(g *Game) (Action, bool) {
	curRound := g.mustCurrentRound()
	curTrick := curRound.currentTrick()
	if curTrick == nil {
		return Action{}, false
	}

	player := curTrick.expectedNextPlayer()
	if player.isZero() || !g.getParticipant(player).IsAI {
		return Action{}, false
	}

	hand := curRound.getHand(player)
	playableCards := hand.playableCardsFor(*curTrick)
	if len(playableCards) == 0 {
		panic("AI: no playable cards for expected next player")
	}

	card := playableCards[rand.IntN(len(playableCards)-1)]
	return Action{
		Name:   PlayCardAction,
		Player: player,
		Rank:   card.Rank,
		Suit:   card.Suit,
	}, true
}

func randomSuit() Suit {
	return allSuits[rand.IntN(4)]
}
