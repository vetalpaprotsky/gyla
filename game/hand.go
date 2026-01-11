package game

import (
	"slices"
)

type hand struct {
	player Player
	cards  []Card
}

func (h hand) deepCopy() hand {
	return hand{
		player: h.player,
		cards:  append([]Card{}, h.cards...),
	}
}

func (h *hand) removeCard(card Card) bool {
	if !slices.Contains(h.cards, card) {
		return false
	}

	newCards := make([]Card, 0, len(h.cards)-1)
	for _, c := range h.cards {
		if c != card {
			newCards = append(newCards, c)
		}
	}

	h.cards = newCards

	return true
}

func (h hand) getCard(rank Rank, suit Suit) Card {
	for _, c := range h.cards {
		if c.Rank == rank && c.Suit == suit {
			return c
		}
	}

	return Card{}
}

func (h hand) availableCardsForMove(trick trick) []Card {
	// First move in a trick, any card works.
	if trick.isEmpty() {
		return h.cards
	}

	// First card in the trick is a trump:
	// - You have to go with a trump as well.
	// - If you don't have trumps then any card works.
	trickFirstCard := trick.firstCard()
	if trickFirstCard.IsTrump {
		if len(h.trumps()) > 0 {
			return h.trumps()
		} else {
			return h.cards
		}
	}

	// First card in the trick is a plain suit:
	// - You have to go with the same plain suit card.
	// - If you don't have any plain cards with that suit then any card works.
	plainSuitCards := h.plainSuitCards(trickFirstCard.Suit)
	if len(plainSuitCards) > 0 {
		return plainSuitCards
	} else {
		return h.cards
	}
}

func (h hand) canMakeMove(card Card, trick trick) bool {
	return slices.Contains(h.availableCardsForMove(trick), card)
}

func (h hand) trumps() []Card {
	// capacity = 2, just a simple guess
	trumps := make([]Card, 0, 2)

	for _, c := range h.cards {
		if c.IsTrump {
			trumps = append(trumps, c)
		}
	}

	return trumps
}

func (h hand) plainSuitCards(suit Suit) []Card {
	// capacity = 2, just a simple guess
	plainCards := make([]Card, 0, 2)

	for _, c := range h.cards {
		if !c.IsTrump && c.Suit == suit {
			plainCards = append(plainCards, c)
		}
	}

	return plainCards
}
