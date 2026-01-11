package game

import (
	"maps"
)

type trick struct {
	number  int
	starter Player
	plrsRel playersRelation
	cards   map[Player]Card
}

func (t trick) deepCopy() trick {
	return trick{
		number:  t.number,
		starter: t.starter,
		plrsRel: t.plrsRel,
		cards:   maps.Clone(t.cards),
	}
}

func newFirstTrick(starter Player, plrsRel playersRelation) trick {
	return trick{number: 1, starter: starter, plrsRel: plrsRel, cards: make(map[Player]Card)}
}

func newTrick(curTrick trick) (trick, error) {
	if curTrick.number >= tricksPerRoundCount {
		return trick{}, newTooManyTricksPerRoundError()
	}

	winner, winnerOk := curTrick.winner()
	if !winnerOk {
		return trick{}, newNoTrickWinnerError()
	}

	return trick{
		number:  curTrick.number + 1,
		starter: winner,
		plrsRel: curTrick.plrsRel,
		cards:   make(map[Player]Card),
	}, nil
}

func (t *trick) addCard(player Player, card Card) error {
	if t.isCompleted() {
		return newTrickCompletedError()
	} else if !t.isPlayerValid(player) {
		return newInvalidPlayerError(player)
	} else if t.playerAlreadyAddedCard(player) {
		return newDuplicatedPlayerMoveError(player, t.cards[player])
	} else if expPlr := t.expectedNextPlayer(); expPlr != player {
		return newUnexpectedPlayerMoveError(player, expPlr)
	}

	t.cards[player] = card
	return nil
}

func (t trick) winner() (Player, bool) {
	if !t.isCompleted() {
		return Player(""), false
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

	return winPlayer, true
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

func (t trick) isPlayerValid(p Player) bool {
	return t.plrsRel.isPlayerValid(p)
}

func (t trick) playerAlreadyAddedCard(p Player) bool {
	_, ok := t.cards[p]

	return ok
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
		player = t.plrsRel.getLeftOpponent(player)
		if _, ok := t.cards[player]; !ok {
			return player
		}
	}
}
