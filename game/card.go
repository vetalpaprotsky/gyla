package game

import (
	"errors"
	"strings"
)

type Suit string
type Rank string
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
		return Card{}, errors.New("Invalid Rank: " + string(rank))
	}
	if !suit.isValid() {
		return Card{}, errors.New("Invalid Suit: " + string(suit))
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
	return c.Rank == sevenRank || c.Rank == jackRank
}

func (c Card) level() int {
	var level int

	if c.isDefaultTrump() {
		switch c.rankAndSuit() {
		case string(sevenRank) + string(clubsSuit):
			level = 21
		case string(sevenRank) + string(spadesSuit):
			level = 20
		case string(sevenRank) + string(heartsSuit):
			level = 19
		case string(sevenRank) + string(diamondsSuit):
			level = 18
		case string(jackRank) + string(clubsSuit):
			level = 17
		case string(jackRank) + string(spadesSuit):
			level = 16
		case string(jackRank) + string(heartsSuit):
			level = 15
		case string(jackRank) + string(diamondsSuit):
			level = 14
		}
	} else if c.IsTrump {
		switch c.Rank {
		case sixRank:
			level = 22
		case aceRank:
			level = 13
		case kingRank:
			level = 12
		case queenRank:
			level = 11
		case tenRank:
			level = 10
		case nineRank:
			level = 9
		case eightRank:
			level = 8
		}
	} else {
		switch c.Rank {
		case aceRank:
			level = 7
		case kingRank:
			level = 6
		case queenRank:
			level = 5
		case tenRank:
			level = 4
		case nineRank:
			level = 3
		case eightRank:
			level = 2
		case sixRank:
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
