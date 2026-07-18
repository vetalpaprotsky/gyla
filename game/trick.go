package game

const tricksPerRoundCount = 9

type trick struct {
	number      int
	starter     Player
	playedCards map[Player]Card
}

func newFirstTrick(starter Player) trick {
	return trick{
		number:      1,
		starter:     starter,
		playedCards: make(map[Player]Card, len(allPlayers)),
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
		playedCards: make(map[Player]Card, len(allPlayers)),
	}, nil
}

func (t *trick) addCard(player Player, card Card) error {
	if t.isCompleted() {
		return newTooManyCardsPerTrickError()
	} else if expPlr := t.expectedNextPlayer(); expPlr != player {
		return newUnexpectedPlayerError(player, expPlr)
	}

	t.playedCards[player] = card
	return nil
}

func (t trick) winner() Player {
	if !t.isCompleted() {
		return Player(0)
	}

	winPlayer := t.starter
	winCard := t.firstCard()

	if t.hasAnyTrumps() {
		for player, card := range t.playedCards {
			if card.level() > winCard.level() {
				winPlayer = player
				winCard = card
			}
		}
	} else {
		leadingSuit := t.firstCard().Suit

		for player, card := range t.playedCards {
			if card.Suit == leadingSuit && card.level() > winCard.level() {
				winPlayer = player
				winCard = card
			}
		}
	}

	return winPlayer
}

func (t trick) firstCard() Card {
	if t.isEmpty() {
		return Card{}
	}

	return t.playedCards[t.starter]
}

func (t trick) hasAnyTrumps() bool {
	for _, card := range t.playedCards {
		if card.IsTrump {
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

	player := t.starter
	for i := 0; i < len(t.playedCards); i++ {
		player = player.leftOpponent()
	}

	return player
}

func (t trick) isZero() bool {
	return t.number == 0
}

func (t trick) state() TrickState {
	return newTrickState(t)
}
