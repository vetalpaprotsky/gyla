package models

type Player struct {
	Name          string
	Team          *Team
	leftOpponent  *Player
	teammate      *Player // Is this needed?
	rightOpponent *Player
}

// TODO: This simple approach should work very good
// type Player string
// type Team string

// type PlayersRelation struct {
// 	player1 Player
// 	player3 Player
// 	team1   Team

// 	player2 Player
// 	player4 Player
// 	team2   Team
// }

// func (pr PlayersRelation) getTeam(p Player) Team {
// }

// func (pr PlayersRelation) getOpponentTeam(p Player) Team {
// }

// func (pr PlayersRelation) getLeftOpponent(p Player) Player {
// }
