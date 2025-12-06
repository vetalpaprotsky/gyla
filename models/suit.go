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
