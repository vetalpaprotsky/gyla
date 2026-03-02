package game

import "math/rand/v2"

func getAIAction(g *Game) Action {
	curRound := g.match.currentRound()
	if !curRound.isTrumpAssigned() && g.isAI(curRound.trumper()) {
		return Action{
			Name:   AssignTrumpAction,
			Player: curRound.trumper(),
			Suit:   randomSuit(),
		}
	}

	curTrick := curRound.currentTrick()
	if curTrick != nil {
		player := curTrick.expectedNextPlayer()
		if g.isAI(player) {
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
