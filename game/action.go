package game

const (
	MoveAction        = "move"
	TrumpChoiceAction = "trump_choice"
)

type Action struct {
	Name   string
	Player Player
	Rank   Rank
	Suit   Suit
}

type ActionRejectedError struct {
	Action Action
	Msg    string
}
