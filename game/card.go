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

type Card struct {
	Rank    Rank
	Suit    Suit
	IsTrump bool
}

type HandCard struct {
	Card       Card
	IsPlayable bool
}

type PlayedCard struct {
	Player Player
	Card   Card
}

func (c Card) asHandCard(isPlayable bool) HandCard {
	return HandCard{Card: c, IsPlayable: isPlayable}
}

func (c Card) asPlayedCard(player Player) PlayedCard {
	return PlayedCard{Player: player, Card: c}
}

func NewCard(rank Rank, suit Suit) (Card, error) {
	if !rank.isValid() {
		return Card{}, newInvalidRankError(rank)
	}
	if !suit.isValid() {
		return Card{}, newInvalidSuitError(suit)
	}

	card := Card{Rank: rank, Suit: suit}
	if card.IsDefaultTrump() {
		card.IsTrump = true
	}

	return card, nil
}

func (c Card) IsDefaultTrump() bool {
	return c.Rank == SevenRank || c.Rank == JackRank
}

func (c Card) level() int {
	var level int

	if c.IsDefaultTrump() {
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

func (c Card) id() int {
	return int(c.Rank) + int(c.Suit)
}
