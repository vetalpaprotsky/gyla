package game

const (
	sevenOfClubs    = int(SevenRank) + int(ClubsSuit)
	sevenOfSpades   = int(SevenRank) + int(SpadesSuit)
	sevenOfHearts   = int(SevenRank) + int(HeartsSuit)
	sevenOfDiamonds = int(SevenRank) + int(DiamondsSuit)
	jackOfClubs     = int(JackRank) + int(ClubsSuit)
	jackOfSpades    = int(JackRank) + int(SpadesSuit)
	jackOfHearts    = int(JackRank) + int(HeartsSuit)
	jackOfDiamonds  = int(JackRank) + int(DiamondsSuit)
)

type card struct {
	rank    Rank
	suit    Suit
	isTrump bool
}

func newCard(rank Rank, suit Suit) (card, error) {
	if !rank.isValid() {
		return card{}, newInvalidRankError(rank)
	}
	if !suit.isValid() {
		return card{}, newInvalidSuitError(suit)
	}

	card := card{rank: rank, suit: suit}
	if card.isDefaultTrump() {
		card.isTrump = true
	}

	return card, nil
}

func (c card) isDefaultTrump() bool {
	return c.rank == SevenRank || c.rank == JackRank
}

func (c card) level() int {
	var level int

	if c.isDefaultTrump() {
		switch c.id() {
		case sevenOfClubs:
			level = 21
		case sevenOfSpades:
			level = 20
		case sevenOfHearts:
			level = 19
		case sevenOfDiamonds:
			level = 18
		case jackOfClubs:
			level = 17
		case jackOfSpades:
			level = 16
		case jackOfHearts:
			level = 15
		case jackOfDiamonds:
			level = 14
		}
	} else if c.isTrump {
		switch c.rank {
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
		switch c.rank {
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

func (c card) id() int {
	return int(c.rank) + int(c.suit)
}

func (c card) state(isPlayable bool) CardState {
	return newCardState(c, isPlayable)
}
