package game

const (
	PlayerMoveAction  = "player_move"
	TrumpChoiceAction = "trump_choice"
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
