package models

type Suit string

func (s Suit) IsValid() bool {
	for _, validSuit := range ValidSuits {
		if s == validSuit {
			return true
		}
	}

	return false
}

// NOTE: it's part of UI, it shouldn't be here ideally.
func (s Suit) Tui() string {
	var red = "\033[31m"
	var black = "\033[30m"
	var reset = "\033[0m"
	var color string

	if s == HeartsSuit || s == DiamondsSuit {
		color = red
	} else {
		color = black
	}

	var suitSymbol string
	switch s {
	case ClubsSuit:
		suitSymbol = "♣"
	case SpadesSuit:
		suitSymbol = "♠"
	case HeartsSuit:
		suitSymbol = "♥"
	case DiamondsSuit:
		suitSymbol = "♦"
	}

	return color + suitSymbol + reset
}
