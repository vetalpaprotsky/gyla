package models

import (
	"errors"
	"fmt"
	"slices"
)

type Hand struct {
	Player Player
	Cards  []Card
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

func (h *Hand) makeMove(card Card, trick *Trick) error {
	if !h.canMakeMove(card, *trick) {
		msg := fmt.Sprintf(
			"Player %s can't make a move with %s card", h.Player, card,
		)
		return errors.New(msg)
	}

	newCards, isCardRemoved := h.removeCard(card)
	if !isCardRemoved {
		msg := fmt.Sprintf(
			"Count not remove card %s from player %s hand", card, h.Player,
		)
		return errors.New(msg)
	}

	if err := trick.addMove(h.Player, card); err != nil {
		return err
	}

	// This must be the last step, since at this point we know that move can be
	// made for sure.
	h.Cards = newCards

	return nil
}

func (h *Hand) removeCard(card Card) ([]Card, bool) {
	if !slices.Contains(h.Cards, card) {
		return []Card{}, false
	}

	newCards := make([]Card, 0, len(h.Cards)-1)
	for _, c := range h.Cards {
		if c != card {
			newCards = append(newCards, c)
		}
	}

	return newCards, true
}

func (h *Hand) canMakeMove(card Card, trick Trick) bool {
	return slices.Contains(h.availableCardsForMove(trick), card)
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
