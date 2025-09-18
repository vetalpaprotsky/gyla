package models

type Trick struct {
	turns []Turn
}

func (t Trick) leadingTurn() Turn {
	return t.turns[0]
}

func (t Trick) winningTurn() Turn {
	winningTurn := t.leadingTurn()

	if t.hasAnyTrumps() {
		for _, turn := range t.turns {
			if turn.card.level() > winningTurn.card.level() {
				winningTurn = turn
			}
		}
	} else {
		leadingSuit := t.leadingTurn().card.suit

		for _, turn := range t.turns {
			if turn.card.suit == leadingSuit && turn.card.level() > winningTurn.card.level() {
				winningTurn = turn
			}
		}
	}

	return winningTurn
}

func (t Trick) hasAnyTrumps() bool {
	for _, turn := range t.turns {
		if turn.card.isTrump {
			return true
		}
	}

	return false
}
