package game

import (
	"math/rand/v2"
)

func getAIAction(match match) Action {
	plrsRel := match.plrsRel
	curRound := match.currentRound()
	if !curRound.isTrumpAssigned() && plrsRel.isAI(curRound.trumper()) {
		return Action{
			Name:   AssignTrumpAction,
			Player: curRound.trumper(),
			Suit:   randomSuit(),
		}
	}

	curTrick := curRound.currentTrick()
	if curTrick != nil {
		player := curTrick.expectedNextPlayer()
		if plrsRel.isAI(player) {
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
	return validSuits[rand.IntN(4)]
}
