package game

import (
	"math/rand/v2"
)

type ai map[Player]bool

func (ai ai) getAction(match match) Action {
	curRound := match.currentRound()
	if !curRound.isTrumpAssigned() && ai[curRound.trumper()] {
		return Action{
			Name:   AssignTrumpAction,
			Player: curRound.trumper(),
			Suit:   randomSuit(),
		}
	}

	curTrick := curRound.currentTrick()
	if curTrick != nil {
		player := curTrick.expectedNextPlayer()
		if ai[player] {
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
