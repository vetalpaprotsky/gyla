package game

type TrickState struct {
	Number      int
	Next        Player
	PlayedCards []PlayedCard
	Winner      Player
}

type trick struct {
	number      int
	starter     Player
	table       Table
	playedCards []PlayedCard
}

func newFirstTrick(starter Player, table Table) trick {
	return trick{
		number:      1,
		starter:     starter,
		table:       table,
		playedCards: make([]PlayedCard, 0, playersCount),
	}
}

func newTrick(curTrick trick) (trick, error) {
	if curTrick.number >= tricksPerRoundCount {
		return trick{}, newTooManyTricksPerRoundError()
	}

	winner := curTrick.winner()
	if winner == Player("") {
		return trick{}, newNoTrickWinnerError()
	}

	return trick{
		number:      curTrick.number + 1,
		starter:     winner,
		table:       curTrick.table,
		playedCards: make([]PlayedCard, 0, playersCount),
	}, nil
}

func (t *trick) addCard(player Player, card Card) error {
	if t.isCompleted() {
		return newTooManyCardsPerTrickError()
	} else if expPlr := t.expectedNextPlayer(); expPlr != player {
		return newUnexpectedPlayerError(player, expPlr)
	}

	t.playedCards = append(t.playedCards, PlayedCard{Player: player, Card: card})
	return nil
}

func (t trick) state() TrickState {
	return TrickState{
		Number:      t.number,
		Next:        t.expectedNextPlayer(),
		PlayedCards: append([]PlayedCard{}, t.playedCards...),
		Winner:      t.winner(),
	}
}

func (t trick) winner() Player {
	if !t.isCompleted() {
		return Player("")
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
	return len(t.playedCards) == playersCount
}

func (t trick) expectedNextPlayer() Player {
	if t.isCompleted() {
		return Player("")
	}

	if t.isEmpty() {
		return t.starter
	}

	lastPlayer := t.playedCards[len(t.playedCards)-1].Player
	return t.table.getLeftOpponent(lastPlayer)
}
