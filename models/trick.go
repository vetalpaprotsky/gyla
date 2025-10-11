package models

type Trick struct {
	number  int
	starter Player
	turns   []Turn
}

func NewFirstTrick(starter Player) *Trick {
	return &Trick{number: 1, starter: starter}
}

func NewTrick(prevTrick *Trick) *Trick {
	return &Trick{number: prevTrick.number + 1, starter: prevTrick.winner()}
}

func (trick Trick) winningTurn() Turn {
	firstTurn := *trick.firstTurn()
	winningTurn := firstTurn

	if trick.hasAnyTrumps() {
		for _, turn := range trick.turns {
			if turn.card.level() > winningTurn.card.level() {
				winningTurn = turn
			}
		}
	} else {
		leadingSuit := firstTurn.card.suit

		for _, turn := range trick.turns {
			if turn.card.suit == leadingSuit && turn.card.level() > winningTurn.card.level() {
				winningTurn = turn
			}
		}
	}

	return winningTurn
}

func (trick Trick) winner() Player {
	return trick.winningTurn().player
}

func (trick Trick) firstTurn() *Turn {
	for _, turn := range trick.turns {
		if turn.player.Name == trick.starter.Name {
			return &turn
		}
	}

	// Not expected
	return nil
}

func (trick Trick) hasAnyTrumps() bool {
	for _, turn := range trick.turns {
		if turn.card.isTrump {
			return true
		}
	}

	return false
}
