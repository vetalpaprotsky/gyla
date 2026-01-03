package models

import (
	"errors"
	"fmt"
)

type Trick struct {
	Number  int
	starter Player
	plrsRel PlayersRelation
	Cards   map[Player]Card
}

func newFirstTrick(starter Player, plrsRel PlayersRelation) Trick {
	return Trick{Number: 1, starter: starter, plrsRel: plrsRel, Cards: make(map[Player]Card)}
}

// NOTE: All these errors are NOT meant to be retriable.
func newTrick(curTrick Trick) (Trick, error) {
	if curTrick.Number >= tricksPerRoundCount {
		msg := "Max possible tricks per round started. Can't start a new one."
		return Trick{}, errors.New(msg)
	}

	winner, winnerOk := curTrick.Winner()
	if !winnerOk {
		msg := "Current trick winner isn't determined. Can't start a new one."
		return Trick{}, errors.New(msg)
	}

	return Trick{
		Number:  curTrick.Number + 1,
		starter: winner,
		plrsRel: curTrick.plrsRel,
		Cards:   make(map[Player]Card),
	}, nil
}

// NOTE: All these errors are meant to be retriable.
func (t *Trick) addCard(player Player, card Card) error {
	var errMsg string

	if t.IsCompleted() {
		errMsg = "Trick is completed. Can't add a new card to it."
	} else if !t.isPlayerValid(player) {
		errMsg = fmt.Sprintf("Player <%s> is invalid.", player)
	} else if t.playerAlreadyAddedCard(player) {
		errMsg = fmt.Sprintf(
			"Player <%s> already added card <%s> to a trick.",
			player,
			t.Cards[player],
		)
	} else if t.expectedNextPlayer() != player {
		errMsg = fmt.Sprintf(
			"Player <%s> is expected to add the next card to a trick, not <%s>.",
			t.expectedNextPlayer(),
			player,
		)
	}

	if errMsg != "" {
		return errors.New(errMsg)
	}

	t.Cards[player] = card
	return nil
}

func (t Trick) Winner() (Player, bool) {
	if !t.IsCompleted() {
		return Player(""), false
	}

	winPlayer := t.starter
	firstCard := t.firstCard()
	winCard := firstCard

	if t.hasAnyTrumps() {
		for player, card := range t.Cards {
			if card.Level() > winCard.Level() {
				winPlayer = player
				winCard = card
			}
		}
	} else {
		leadingSuit := firstCard.Suit

		for player, card := range t.Cards {
			if card.Suit == leadingSuit && card.Level() > winCard.Level() {
				winPlayer = player
				winCard = card
			}
		}
	}

	return winPlayer, true
}

func (t Trick) firstCard() Card {
	return t.Cards[t.starter]
}

func (t Trick) hasAnyTrumps() bool {
	for _, card := range t.Cards {
		if card.IsTrump {
			return true
		}
	}

	return false
}

func (t Trick) isPlayerValid(p Player) bool {
	return t.plrsRel.isPlayerValid(p)
}

func (t Trick) playerAlreadyAddedCard(p Player) bool {
	_, ok := t.Cards[p]

	return ok
}

func (t Trick) isEmpty() bool {
	return len(t.Cards) == 0
}

func (t Trick) IsCompleted() bool {
	return len(t.Cards) == playersCount
}

func (t Trick) expectedNextPlayer() Player {
	if t.IsCompleted() {
		return Player("")
	}

	if t.isEmpty() {
		return t.starter
	}

	player := t.starter
	for {
		player = t.plrsRel.getLeftOpponent(player)
		if _, ok := t.Cards[player]; !ok {
			return player
		}
	}
}
