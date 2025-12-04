package models

import (
	"errors"
	"strings"
)

type Card struct {
	Rank    Rank
	Suit    Suit
	IsTrump bool
}

func newCard(rank Rank, suit Suit) (Card, error) {
	if !rank.IsValid() {
		return Card{}, errors.New("Invalid Rank: " + string(rank))
	}
	if !suit.IsValid() {
		return Card{}, errors.New("Invalid Suit: " + string(suit))
	}

	card := Card{Rank: rank, Suit: suit}
	if card.isDefaultTrump() {
		card.IsTrump = true
	}

	return card, nil
}

func NewCardFromCardID(id string) (Card, error) {
	rank := Rank(strings.ToUpper(id[:len(id)-1]))
	suit := Suit(strings.ToUpper(id[len(id)-1:]))

	return newCard(rank, suit)
}

func (c Card) isDefaultTrump() bool {
	return c.Rank == SevenRank || c.Rank == JackRank
}

func (c Card) Level() int {
	var level int

	if c.isDefaultTrump() {
		switch c.ID() {
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

func (c Card) ID() string {
	return string(c.Rank) + string(c.Suit)
}
