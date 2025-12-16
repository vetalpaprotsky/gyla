package game

const (
	MoveAction        = "move"
	TrumpChoiceAction = "trump_choice"
)

type PlayerAction struct {
	Player Player
	Action string
	Rank   Rank
	Suit   Suit
}
