package models

import "errors"

type Card struct {
	Rank    string
	Suit    string
	isTrump bool
}

func isSuitValid(suit string) bool {
	for _, validSuit := range ValidSuits {
		if suit == validSuit {
			return true
		}
	}

	return false
}

func isRankValid(rank string) bool {
	for _, validRank := range ValidRanks {
		if rank == validRank {
			return true
		}
	}

	return false
}

func newCard(rank, suit string) (*Card, error) {
	if !isRankValid(rank) {
		return nil, errors.New("Invalid Rank: " + rank)
	}
	if !isSuitValid(suit) {
		return nil, errors.New("Invalid Suit: " + suit)
	}

	card := Card{Rank: rank, Suit: suit}
	if card.isDefaultTrump() {
		card.isTrump = true
	}

	return &card, nil
}

func (c Card) isDefaultTrump() bool {
	return c.Rank == SevenRank || c.Rank == JackRank
}

func (c Card) level() int {
	var level int

	if c.isDefaultTrump() {
		switch c.ID() {
		case SevenRank + ClubsSuit:
			level = 21
		case SevenRank + SpadesSuit:
			level = 20
		case SevenRank + HeartsSuit:
			level = 19
		case SevenRank + DiamondsSuit:
			level = 18
		case JackRank + ClubsSuit:
			level = 17
		case JackRank + SpadesSuit:
			level = 16
		case JackRank + HeartsSuit:
			level = 15
		case JackRank + DiamondsSuit:
			level = 14
		}
	} else if c.isTrump {
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
	return c.Rank + c.Suit
}
