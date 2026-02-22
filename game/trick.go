package game

import (
	"maps"
)

type trick struct {
	number  int
	starter Player
	table   table
	cards   map[Player]Card
}

func (t trick) deepCopy() trick {
	return trick{
		number:  t.number,
		starter: t.starter,
		table:   t.table,
		cards:   maps.Clone(t.cards),
	}
}

func newFirstTrick(starter Player, table table) trick {
	return trick{number: 1, starter: starter, table: table, cards: make(map[Player]Card)}
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
		number:  curTrick.number + 1,
		starter: winner,
		table:   curTrick.table,
		cards:   make(map[Player]Card),
	}, nil
}

func (t *trick) addCard(player Player, card Card) error {
	if t.isCompleted() {
		return newTooManyCardsPerTrickError()
	} else if expPlr := t.expectedNextPlayer(); expPlr != player {
		return newUnexpectedPlayerError(player, expPlr)
	}

	t.cards[player] = card
	return nil
}

func (t trick) winner() Player {
	if !t.isCompleted() {
		return Player("")
	}

	winPlayer := t.starter
	firstCard := t.firstCard()
	winCard := firstCard

	if t.hasAnyTrumps() {
		for player, card := range t.cards {
			if card.level() > winCard.level() {
				winPlayer = player
				winCard = card
			}
		}
	} else {
		leadingSuit := firstCard.Suit

		for player, card := range t.cards {
			if card.Suit == leadingSuit && card.level() > winCard.level() {
				winPlayer = player
				winCard = card
			}
		}
	}

	return winPlayer
}

func (t trick) firstCard() Card {
	return t.cards[t.starter]
}

func (t trick) hasAnyTrumps() bool {
	for _, card := range t.cards {
		if card.IsTrump {
			return true
		}
	}

	return false
}

func (t trick) isEmpty() bool {
	return len(t.cards) == 0
}

func (t trick) isCompleted() bool {
	return len(t.cards) == playersCount
}

func (t trick) expectedNextPlayer() Player {
	if t.isCompleted() {
		return Player("")
	}

	if t.isEmpty() {
		return t.starter
	}

	player := t.starter
	for {
		player = t.table.getLeftOpponent(player)
		if _, ok := t.cards[player]; !ok {
			return player
		}
	}
}
