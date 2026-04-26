package game

const tricksPerRoundCount = 9

type trick struct {
	number      int
	starter     Player
	playedCards []PlayedCard
}

func newFirstTrick(starter Player) trick {
	return trick{
		number:      1,
		starter:     starter,
		playedCards: make([]PlayedCard, 0, len(allPlayers)),
	}
}

func newTrick(curTrick trick) (trick, error) {
	if curTrick.number >= tricksPerRoundCount {
		return trick{}, newTooManyTricksPerRoundError()
	}

	winner := curTrick.winner()
	if winner.isZero() {
		return trick{}, newNoTrickWinnerError()
	}

	return trick{
		number:      curTrick.number + 1,
		starter:     winner,
		playedCards: make([]PlayedCard, 0, len(allPlayers)),
	}, nil
}

func (t *trick) addCard(player Player, card Card) error {
	if t.isCompleted() {
		return newTooManyCardsPerTrickError()
	} else if expPlr := t.expectedNextPlayer(); expPlr != player {
		return newUnexpectedPlayerError(player, expPlr)
	}

	t.playedCards = append(t.playedCards, card.asPlayedCard(player))
	return nil
}

func (t trick) winner() Player {
	if !t.isCompleted() {
		return Player(0)
	}

	winPlayer := t.starter
	winCard := t.firstCard()

	if t.hasAnyTrumps() {
		for _, pc := range t.playedCards {
			if pc.Card.level() > winCard.level() {
				winPlayer = pc.Player
				winCard = pc.Card
			}
		}
	} else {
		leadingSuit := t.firstCard().Suit

		for _, pc := range t.playedCards {
			if pc.Card.Suit == leadingSuit && pc.Card.level() > winCard.level() {
				winPlayer = pc.Player
				winCard = pc.Card
			}
		}
	}

	return winPlayer
}

func (t trick) firstCard() Card {
	if t.isEmpty() {
		return Card{}
	}

	return t.playedCards[0].Card
}

func (t trick) hasAnyTrumps() bool {
	for _, pc := range t.playedCards {
		if pc.Card.IsTrump {
			return true
		}
	}

	return false
}

func (t trick) isEmpty() bool {
	return len(t.playedCards) == 0
}

func (t trick) isCompleted() bool {
	return len(t.playedCards) == len(allPlayers)
}

func (t trick) expectedNextPlayer() Player {
	if t.isCompleted() {
		return Player(0)
	}

	if t.isEmpty() {
		return t.starter
	}

	lastPlayer := t.playedCards[len(t.playedCards)-1].Player
	return lastPlayer.leftOpponent()
}

func (t trick) isZero() bool {
	return t.number == 0
}

func (t trick) state() TrickState {
	return newTrickState(t)
}
