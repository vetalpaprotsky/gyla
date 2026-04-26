package game

type CardState struct {
	Rank       Rank
	Suit       Suit
	IsTrump    bool
	IsPlayable bool
}

func newCardState(c card, isPlayable bool) CardState {
	return CardState{
		Rank:       c.rank,
		Suit:       c.suit,
		IsTrump:    c.isTrump,
		IsPlayable: isPlayable,
	}
}
