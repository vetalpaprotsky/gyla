package models

import (
	"fmt"
)

type Hand struct {
	Player Player
	Cards  []Card
}

// TODO: I don't think I need a pointer here.
func (h *Hand) assignTrump(suit Suit) {
	for i, c := range h.Cards {
		if c.Suit == suit {
			h.Cards[i].IsTrump = true
		}
	}
}

func (h *Hand) availableCardsForMove(trick Trick) []Card {
	// First move in a trick, any card works
	if len(trick.Moves) == 0 {
		return h.Cards
	}

	// First card in the trick is a trump:
	// - You have to go with trump as well
	// - If you don't have trumps then any card works
	trickFirstCard := trick.Moves[0].Card
	if trickFirstCard.IsTrump {
		if len(h.trumps()) > 0 {
			return h.trumps()
		} else {
			return h.Cards
		}
	}

	// First card in the trick is a plain suit:
	// - You have to go with the same plain suit card
	// - If you don't any plain cards with that suit then any card works
	plainSuitCards := h.plainSuitCards(trickFirstCard.Suit)
	if len(plainSuitCards) > 0 {
		return plainSuitCards
	} else {
		return h.Cards
	}
}

func (h *Hand) makeMove(card Card) error {
	newCards := make([]Card, 0, len(h.Cards))

	for _, c := range h.Cards {
		if c != card {
			newCards = append(newCards, c)
		}
	}

	// This means that none of the Cards in the hand match the card that
	// was passed to the function. It's not expected. You can't take a move with
	// a card that you don't have.
	if len(newCards) == len(h.Cards) {
		return fmt.Errorf("Player %s doesn't have card %s", h.Player, card.ID())
	}

	h.Cards = newCards

	return nil
}

func (h *Hand) trumps() []Card {
	// capacity = 2, just a simple guess
	trumps := make([]Card, 0, 2)

	for _, c := range h.Cards {
		if c.IsTrump {
			trumps = append(trumps, c)
		}
	}

	return trumps
}

func (h *Hand) plainSuitCards(suit Suit) []Card {
	// capacity = 2, just a simple guess
	plainCards := make([]Card, 0, 2)

	for _, c := range h.Cards {
		if !c.IsTrump && c.Suit == suit {
			plainCards = append(plainCards, c)
		}
	}

	return plainCards
}
