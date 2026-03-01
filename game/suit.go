package game

type Suit string

// TODO: We could store them as bites(chars)
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
