package game

type Suit int

const (
	ClubsSuit    = 10
	SpadesSuit   = 20
	HeartsSuit   = 30
	DiamondsSuit = 40
)

var validSuits = [4]Suit{
	ClubsSuit,
	SpadesSuit,
	HeartsSuit,
	DiamondsSuit,
}

func (s Suit) IsZero() bool {
	return s == Suit(0)
}

func (s Suit) isValid() bool {
	for _, validSuit := range validSuits {
		if s == validSuit {
			return true
		}
	}

	return false
}
