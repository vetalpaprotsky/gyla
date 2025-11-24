package models

type Trick struct {
	Number  int
	starter Player
	Moves   []Move
}

func newFirstTrick(starter Player) Trick {
	return Trick{Number: 1, starter: starter}
}

// TODO: Add error is more than 9 tricks started, or curTrick isn't finished.
func newTrick(curTrick Trick) Trick {
	return Trick{Number: curTrick.Number + 1, starter: curTrick.Winner()}
}

// TODO: Add error if more than 4 moves added, or same player added, or same card
// added, or order is incorrect(last player move be a right opponent of the current player).
func (t *Trick) addMove(player Player, card Card) {
	t.Moves = append(t.Moves, Move{Player: player, Card: card})
}

// TODO: Add error if trick isn't completed (len(t.Moves) != 4)
// In that case it should return *Move and error
func (t Trick) winMove() Move {
	firstMove := *t.firstMove()
	winMove := firstMove

	if t.hasAnyTrumps() {
		for _, move := range t.Moves {
			if move.Card.level() > winMove.Card.level() {
				winMove = move
			}
		}
	} else {
		leadingSuit := firstMove.Card.Suit

		for _, move := range t.Moves {
			if move.Card.Suit == leadingSuit && move.Card.level() > winMove.Card.level() {
				winMove = move
			}
		}
	}

	return winMove
}

func (t Trick) Winner() Player {
	return t.winMove().Player
}

func (t Trick) firstMove() *Move {
	for i := 0; i < len(t.Moves); i++ {
		move := &t.Moves[i]
		if move.Player.Name == t.starter.Name {
			return move
		}
	}

	// Not expected
	return nil
}

func (t Trick) hasAnyTrumps() bool {
	for _, move := range t.Moves {
		if move.Card.isTrump {
			return true
		}
	}

	return false
}

func (t Trick) IsCompleted() bool {
	return len(t.Moves) == movesInTrickCount
}
