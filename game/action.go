package game

const (
	PlayCardAction    = "play_card"
	AssignTrumpAction = "assign_trump"
)

type Action struct {
	Name   string
	Player Player
	Rank   Rank
	Suit   Suit
}

type ExpectedAction struct {
	Name   string
	Player Player
}

type ActionResult struct {
	ErrorMsg  string
	Succeeded bool
}
