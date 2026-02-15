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

type ActionResult struct {
	ErrorMsg  string
	Succeeded bool
}
