package models

import (
	"errors"
	"fmt"
)

type Trick struct {
	Number  int
	starter Player
	Moves   []Move
}

func newFirstTrick(starter Player) Trick {
	return Trick{Number: 1, starter: starter}
}

func newTrick(curTrick Trick) (Trick, error) {
	if curTrick.Number >= tricksPerRoundCount {
		msg := "Max possible tricks per round started. Can't start a new one."
		return Trick{}, errors.New(msg)
	}

	return Trick{Number: curTrick.Number + 1, starter: curTrick.Winner()}, nil
}

func (t *Trick) addMove(player Player, card Card) error {
	if len(t.Moves) >= movesPerTrickCount {
		msg := "Max possible moves per trick added. Can't add a new one."
		return errors.New(msg)
	}

	for _, m := range t.Moves {
		if m.Player == player {
			msg := fmt.Sprintf("Player %s already made a move in a trick", player)
			return errors.New(msg)
		}

		if m.Card.ID() == card.ID() {
			msg := fmt.Sprintf("Card %s already in a trick", card.ID())
			return errors.New(msg)
		}
	}

	t.Moves = append(t.Moves, Move{Player: player, Card: card})

	return nil
}

// TODO: use ok idiom if trick isn't completed
func (t Trick) winMove() Move {
	firstMove := *t.firstMove()
	winMove := firstMove

	if t.hasAnyTrumps() {
		for _, move := range t.Moves {
			if move.Card.Level() > winMove.Card.Level() {
				winMove = move
			}
		}
	} else {
		leadingSuit := firstMove.Card.Suit

		for _, move := range t.Moves {
			if move.Card.Suit == leadingSuit && move.Card.Level() > winMove.Card.Level() {
				winMove = move
			}
		}
	}

	return winMove
}

// TODO: comma ok idiom
func (t Trick) Winner() Player {
	return t.winMove().Player
}

// TODO: Use comma ok idiom, don't return a pointer.
func (t Trick) firstMove() *Move {
	for i := 0; i < len(t.Moves); i++ {
		move := &t.Moves[i]
		if move.Player == t.starter {
			return move
		}
	}

	// Not expected
	return nil
}

func (t Trick) hasAnyTrumps() bool {
	for _, move := range t.Moves {
		if move.Card.IsTrump {
			return true
		}
	}

	return false
}

func (t Trick) IsCompleted() bool {
	return len(t.Moves) == movesPerTrickCount
}
