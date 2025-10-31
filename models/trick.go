package models

type Trick struct {
	number  int
	starter Player
	moves   []Move
}

func newFirstTrick(starter Player) *Trick {
	return &Trick{number: 1, starter: starter}
}

// TODO: Add error is more than 4 tricks started
func newTrick(curTrick *Trick) *Trick {
	return &Trick{number: curTrick.number + 1, starter: curTrick.winner()}
}

// TODO: Add error if more than 4 moves added, or same player added, or same card
// added.
func (t *Trick) addMove(player Player, card Card) {
	t.moves = append(t.moves, Move{player: player, card: card})
}

// TODO: Add error if trick isn't completed (len(t.moves) != 4)
// In that case it should return *Move and error
func (t Trick) winMove() Move {
	firstMove := *t.firstMove()
	winMove := firstMove

	if t.hasAnyTrumps() {
		for _, move := range t.moves {
			if move.card.level() > winMove.card.level() {
				winMove = move
			}
		}
	} else {
		leadingSuit := firstMove.card.suit

		for _, move := range t.moves {
			if move.card.suit == leadingSuit && move.card.level() > winMove.card.level() {
				winMove = move
			}
		}
	}

	return winMove
}

func (t Trick) winner() Player {
	return t.winMove().player
}

func (t Trick) firstMove() *Move {
	for i := 0; i < len(t.moves); i++ {
		move := &t.moves[i]
		if move.player.Name == t.starter.Name {
			return move
		}
	}

	// Not expected
	return nil
}

func (t Trick) hasAnyTrumps() bool {
	for _, move := range t.moves {
		if move.card.isTrump {
			return true
		}
	}

	return false
}
