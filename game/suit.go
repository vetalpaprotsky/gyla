package game

type Suit int

const (
	ClubsSuit    = Suit(10)
	SpadesSuit   = Suit(20)
	HeartsSuit   = Suit(30)
	DiamondsSuit = Suit(40)
)

var allSuits = [4]Suit{
	ClubsSuit,
	SpadesSuit,
	HeartsSuit,
	DiamondsSuit,
}

func (s Suit) isZero() bool {
	return s == Suit(0)
}

func (s Suit) isValid() bool {
	for _, suit := range allSuits {
		if s == suit {
			return true
		}
	}

	return false
}
