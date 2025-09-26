package models

type Player struct {
	Name          string
	Team          *Team
	LeftOpponent  *Player
	Teammate      *Player
	RightOpponent *Player
}
