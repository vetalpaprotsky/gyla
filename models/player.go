package models

type Player struct {
	Name          string
	Team          *Team
	leftOpponent  *Player
	teammate      *Player // Is this needed?
	rightOpponent *Player
}
