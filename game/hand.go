package game

import "slices"

type Hand struct {
	Player Player
	Cards  []Card
}

func (h *Hand) playCard(card Card) bool {
	if !slices.Contains(h.Cards, card) {
		return false
	}

	newCards := make([]Card, 0, len(h.Cards)-1)
	for _, c := range h.Cards {
		if c != card {
			newCards = append(newCards, c)
		}
	}

	h.Cards = newCards

	return true
}

func (h Hand) deepCopy() Hand {
	return Hand{
		Player: h.Player,
		Cards:  append([]Card{}, h.Cards...),
	}
}

func (h Hand) getCard(rank Rank, suit Suit) Card {
	for _, c := range h.Cards {
		if c.Rank == rank && c.Suit == suit {
			return c
		}
	}

	return Card{}
}

func (h Hand) playableCardsFor(trick trick) []Card {
	// Not h.Player turn. No cards playable.
	if h.Player != trick.expectedNextPlayer() {
		return nil
	}

	// First card in a trick, any card works.
	if trick.isEmpty() {
		return h.Cards
	}

	// First card in the trick is a trump:
	// - You have to go with a trump as well.
	// - If you don't have trumps then any card works.
	trickFirstCard := trick.firstCard()
	if trickFirstCard.IsTrump {
		if len(h.trumps()) > 0 {
			return h.trumps()
		} else {
			return h.Cards
		}
	}

	// First card in the trick is a plain suit:
	// - You have to go with the same plain suit card.
	// - If you don't have any plain cards with that suit then any card works.
	plainSuitCards := h.plainSuitCards(trickFirstCard.Suit)
	if len(plainSuitCards) > 0 {
		return plainSuitCards
	} else {
		return h.Cards
	}
}

func (h Hand) canPlayCard(card Card, trick trick) bool {
	return slices.Contains(h.playableCardsFor(trick), card)
}

func (h Hand) trumps() []Card {
	// capacity = 2, just a simple guess
	trumps := make([]Card, 0, 2)

	for _, c := range h.Cards {
		if c.IsTrump {
			trumps = append(trumps, c)
		}
	}

	return trumps
}

func (h Hand) plainSuitCards(suit Suit) []Card {
	// capacity = 2, just a simple guess
	plainCards := make([]Card, 0, 2)

	for _, c := range h.Cards {
		if !c.IsTrump && c.Suit == suit {
			plainCards = append(plainCards, c)
		}
	}

	return plainCards
}
