package game

import (
	"strings"
)

type Card struct {
	Rank    Rank
	Suit    Suit
	IsTrump bool
}

func (s Suit) isValid() bool {
	for _, validSuit := range validSuits {
		if s == validSuit {
			return true
		}
	}

	return false
}

func (r Rank) isValid() bool {
	for _, validRank := range validRanks {
		if r == validRank {
			return true
		}
	}

	return false
}

func newCard(rank Rank, suit Suit) (Card, error) {
	if !rank.isValid() {
		return Card{}, newInvalidRankError(rank)
	}
	if !suit.isValid() {
		return Card{}, newInvalidSuitError(suit)
	}

	card := Card{Rank: rank, Suit: suit}
	if card.isDefaultTrump() {
		card.IsTrump = true
	}

	return card, nil
}

func newCardFromRankAndSuit(rankAndSuit string) (Card, error) {
	rank := Rank(strings.ToUpper(rankAndSuit[:len(rankAndSuit)-1]))
	suit := Suit(strings.ToUpper(rankAndSuit[len(rankAndSuit)-1:]))

	return newCard(rank, suit)
}

func (c Card) isDefaultTrump() bool {
	return c.Rank == SevenRank || c.Rank == JackRank
}

func (c Card) level() int {
	var level int

	if c.isDefaultTrump() {
		switch c.rankAndSuit() {
		case string(SevenRank) + string(ClubsSuit):
			level = 21
		case string(SevenRank) + string(SpadesSuit):
			level = 20
		case string(SevenRank) + string(HeartsSuit):
			level = 19
		case string(SevenRank) + string(DiamondsSuit):
			level = 18
		case string(JackRank) + string(ClubsSuit):
			level = 17
		case string(JackRank) + string(SpadesSuit):
			level = 16
		case string(JackRank) + string(HeartsSuit):
			level = 15
		case string(JackRank) + string(DiamondsSuit):
			level = 14
		}
	} else if c.IsTrump {
		switch c.Rank {
		case SixRank:
			level = 22
		case AceRank:
			level = 13
		case KingRank:
			level = 12
		case QueenRank:
			level = 11
		case TenRank:
			level = 10
		case NineRank:
			level = 9
		case EightRank:
			level = 8
		}
	} else {
		switch c.Rank {
		case AceRank:
			level = 7
		case KingRank:
			level = 6
		case QueenRank:
			level = 5
		case TenRank:
			level = 4
		case NineRank:
			level = 3
		case EightRank:
			level = 2
		case SixRank:
			level = 1
		}
	}

	return level
}

func (c Card) rankAndSuit() string {
	return string(c.Rank) + string(c.Suit)
}

func (c Card) String() string {
	var suffix string

	if c.IsTrump {
		suffix = "-T"
	} else {
		suffix = "-P"
	}

	return c.rankAndSuit() + suffix
}
