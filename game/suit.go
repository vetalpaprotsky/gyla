package game

type Suit string

// TODO: We could store them as bites(chars)
const (
	ClubsSuit    = Suit("C")
	SpadesSuit   = Suit("S")
	HeartsSuit   = Suit("H")
	DiamondsSuit = Suit("D")
)

var validSuits = [4]Suit{
	ClubsSuit,
	SpadesSuit,
	HeartsSuit,
	DiamondsSuit,
}

func (s Suit) isValid() bool {
	for _, validSuit := range validSuits {
		if s == validSuit {
			return true
		}
	}

	return false
}
