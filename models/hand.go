package models

import (
	"fmt"
)

type Hand struct {
	player Player
	cards  []Card
}

func (h *Hand) assignTrump(suit string) {
	for i, c := range h.cards {
		if c.suit == suit {
			h.cards[i].isTrump = true
		}
	}
}

func (h *Hand) availableCardsForMove(trick Trick) []Card {
	// First move in a trick, any card works
	if len(trick.moves) == 0 {
		return h.cards
	}

	// First card in the trick is a trump:
	// - You have to go with trump as well
	// - If you don't have trumps then any card works
	trickFirstCard := trick.moves[0].card
	if trickFirstCard.isTrump {
		if len(h.trumps()) > 0 {
			return h.trumps()
		} else {
			return h.cards
		}
	}

	// First card in the trick is a plain suit:
	// - You have to go with the same plain suit card
	// - If you don't any plain cards with that suit then any card works
	plainSuitCards := h.plainSuitCards(trickFirstCard.suit)
	if len(plainSuitCards) > 0 {
		return plainSuitCards
	} else {
		return h.cards
	}
}

func (h *Hand) takeMove(card Card) error {
	cardsAfterMove := make([]Card, 0, len(h.cards))

	for _, c := range h.cards {
		if c != card {
			cardsAfterMove = append(cardsAfterMove, c)
		}
	}

	// This means that none of the cards in the hand match the card that
	// was passed to the function. It's not expected. You can't make a move
	// with a card that you don't have.
	if len(cardsAfterMove) == len(h.cards) {
		return fmt.Errorf("Player %s doesn't have card %s", h.player.Name, card.id())
	}

	h.cards = cardsAfterMove

	return nil
}

func (h *Hand) trumps() []Card {
	// capacity = 2, just a simple guess
	trumps := make([]Card, 0, 2)

	for _, c := range h.cards {
		if c.isTrump {
			trumps = append(trumps, c)
		}
	}

	return trumps
}

func (h *Hand) plainSuitCards(suit string) []Card {
	// capacity = 2, just a simple guess
	plainCards := make([]Card, 0, 2)

	for _, c := range h.cards {
		if !c.isTrump && c.suit == suit {
			plainCards = append(plainCards, c)
		}
	}

	return plainCards
}
