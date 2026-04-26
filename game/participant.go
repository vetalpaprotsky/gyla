package game

type Participant struct {
	Player     Player
	Team       Team
	PlayerName string
	TeamName   string
	IsAI       bool
	ExternalID string
}
