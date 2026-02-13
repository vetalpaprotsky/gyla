package game

import (
	"fmt"
	"math/rand/v2"
)

type Bot struct {
	player Player
	round  round
}

func (b Bot) getAction(actName string) Action {
	switch actName {
	case AssignTrumpAction:
		return Action{
			Name:   actName,
			Player: b.player,
			Suit:   randomSuit(),
		}
	case PlayCardAction:
		card := b.getRandomCard()
		return Action{
			Name:   actName,
			Player: b.player,
			Rank:   card.Rank,
			Suit:   card.Suit,
		}
	default:
		panic(fmt.Sprintf("unexpected action: %s", actName))
	}
}

func (b Bot) getRandomCard() Card {
	cards := b.round.playableCardsFor(b.player)

	return cards[rand.IntN(len(cards)-1)]
}

func randomSuit() Suit {
	return validSuits[rand.IntN(4)]
}
