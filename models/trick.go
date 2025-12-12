package models

import (
	"errors"
	"fmt"
)

// TODO: We better not store moves in array. It's no really safe.
// I think it's better to create 4 fields, since we always know that 4 moves
// must be done.
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

	winner, winnerOk := curTrick.Winner()
	if !winnerOk {
		msg := "Current trick winner isn't determined. Can't start a new one."
		return Trick{}, errors.New(msg)
	}

	return Trick{Number: curTrick.Number + 1, starter: winner}, nil
}

// TODO: Validate order! Players must make moves in correct order.
func (t *Trick) addMove(player Player, card Card) error {
	if t.IsCompleted() {
		msg := "Trick is completed. Can't add a new card to it."
		return errors.New(msg)
	}

	for _, m := range t.Moves {
		if m.Player == player {
			msg := fmt.Sprintf("Player <%s> already made a move in a trick", player)
			return errors.New(msg)
		}

		if m.Card == card {
			msg := fmt.Sprintf("Card <%s> already in a trick", card)
			return errors.New(msg)
		}
	}

	t.Moves = append(t.Moves, Move{Player: player, Card: card})

	return nil
}

func (t Trick) winMove() (Move, bool) {
	if len(t.Moves) != movesPerTrickCount {
		return Move{}, false
	}

	// It's safe to skip bool value in this case, since we're sure that
	// the first move is present in the trick at this point.
	firstMove, _ := t.firstMove()
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

	return winMove, true
}

func (t Trick) Winner() (Player, bool) {
	if move, ok := t.winMove(); ok {
		return move.Player, true
	}

	return Player(""), false
}

func (t Trick) IsCompleted() bool {
	if _, ok := t.winMove(); ok {
		return true
	}

	return false
}

func (t Trick) firstMove() (Move, bool) {
	for i := 0; i < len(t.Moves); i++ {
		move := t.Moves[i]
		if move.Player == t.starter {
			return move, true
		}
	}

	return Move{}, false
}

func (t Trick) hasAnyTrumps() bool {
	for _, move := range t.Moves {
		if move.Card.IsTrump {
			return true
		}
	}

	return false
}
