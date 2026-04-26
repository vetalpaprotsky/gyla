package game

const (
	PlayCardAction    = "play_card"
	AssignTrumpAction = "assign_trump"
)

type NextAction struct {
	Name   string
	Player Player
}

func (na NextAction) isZero() bool {
	return na.Name == ""
}

type Action struct {
	Name   string
	Player Player
	Rank   Rank
	Suit   Suit
}

func (a Action) isZero() bool {
	return a.Name == ""
}
