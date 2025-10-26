package models

type Hand struct {
	player Player
	cards  []Card
}

func (hand Hand) assignTrump(suit string) {
	for i, c := range hand.cards {
		if c.suit == suit {
			hand.cards[i].isTrump = true
		}
	}
}
