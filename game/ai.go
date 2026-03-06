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
	curRound := g.mustCurrentRound()

	if !curRound.isTrumpAssigned() && g.getParticipant(curRound.trumper()).IsAI {
		return Action{
			Name:   AssignTrumpAction,
			Player: curRound.trumper(),
			Suit:   randomSuit(),
		}
	}

	curTrick := curRound.currentTrick()
	if curTrick != nil {
		player := curTrick.expectedNextPlayer()
		if g.getParticipant(player).IsAI {
			card := getRandomCard(*curRound, player)
			return Action{
				Name:   PlayCardAction,
				Player: player,
				Rank:   card.Rank,
				Suit:   card.Suit,
			}
		}
	}

	return Action{}
}

func getRandomCard(r round, p Player) Card {
	cards := r.playableCardsFor(p)

	return cards[rand.IntN(len(cards)-1)]
}

func randomSuit() Suit {
	return allSuits[rand.IntN(4)]
}
