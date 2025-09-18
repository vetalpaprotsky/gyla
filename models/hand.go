package models

import (
	"strings"
)

type Hand struct {
	player string
	cards  []Card
}

func (hand Hand) String() string {
	cardsStr := make([]string, len(hand.cards))

	for i, c := range hand.cards {
		cardsStr[i] = c.String()
	}

	return hand.player + ": " + strings.Join(cardsStr, " ")
}

func (hand Hand) assignTrump(suit string) {
	for i, c := range hand.cards {
		if c.suit == suit {
			hand.cards[i].isTrump = true
		}
	}
}
