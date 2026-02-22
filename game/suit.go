package game

type Suit string

const (
	ClubsSuit    = Suit("C")
	SpadesSuit   = Suit("S")
	HeartsSuit   = Suit("H")
	DiamondsSuit = Suit("D")
)

var validSuits = [suitsCount]Suit{
	ClubsSuit,
	SpadesSuit,
	HeartsSuit,
	DiamondsSuit,
}
