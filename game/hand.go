package game

import "slices"

type hand struct {
	player Player
	cards  []card
}

func (h *hand) playCard(c card) bool {
	if !slices.Contains(h.cards, c) {
		return false
	}

	newCards := make([]card, 0, len(h.cards)-1)
	for _, existing := range h.cards {
		if existing != c {
			newCards = append(newCards, existing)
		}
	}

	h.cards = newCards

	return true
}

func (h hand) deepCopy() hand {
	return hand{
		player: h.player,
		cards:  append([]card{}, h.cards...),
	}
}

func (h hand) getCard(rank Rank, suit Suit) card {
	for _, c := range h.cards {
		if c.rank == rank && c.suit == suit {
			return c
		}
	}

	return card{}
}

func (h hand) playableCardsFor(trick trick) []card {
	// Not h.player turn. No cards playable.
	if h.player != trick.expectedNextPlayer() {
		return nil
	}

	// First card in a trick, any card works.
	if trick.isEmpty() {
		return h.cards
	}

	// First card in the trick is a trump:
	// - You have to go with a trump as well.
	// - If you don't have trumps then any card works.
	trickFirstCard := trick.firstCard()
	if trickFirstCard.isTrump {
		if len(h.trumps()) > 0 {
			return h.trumps()
		} else {
			return h.cards
		}
	}

	// First card in the trick is a plain suit:
	// - You have to go with the same plain suit card.
	// - If you don't have any plain cards with that suit then any card works.
	plainSuitCards := h.plainSuitCards(trickFirstCard.suit)
	if len(plainSuitCards) > 0 {
		return plainSuitCards
	} else {
		return h.cards
	}
}

func (h hand) canPlayCard(card card, trick trick) bool {
	return slices.Contains(h.playableCardsFor(trick), card)
}

func (h hand) trumps() []card {
	// capacity = 2, just a simple guess
	trumps := make([]card, 0, 2)

	for _, c := range h.cards {
		if c.isTrump {
			trumps = append(trumps, c)
		}
	}

	return trumps
}

func (h hand) plainSuitCards(suit Suit) []card {
	// capacity = 2, just a simple guess
	plainCards := make([]card, 0, 2)

	for _, c := range h.cards {
		if !c.isTrump && c.suit == suit {
			plainCards = append(plainCards, c)
		}
	}

	return plainCards
}

func (h hand) state(t trick) HandState {
	return newHandState(h, t)
}
