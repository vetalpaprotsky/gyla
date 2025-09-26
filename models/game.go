package models

type Game struct {
	rounds  []Round
	player1 *Player
	player2 *Player
	player3 *Player
	player4 *Player
}

// TODO: It should receive names of all players and teams, and initialize
// all every field like in main.go. It can even start a first round?
func NewGame() {

}

// TODO
func NewGameFromJSON() {

}

// TODO
func (g Game) ToJSON() {

}

func (g Game) team1() Team {
	return *g.player1.Team
}

func (g Game) team2() Team {
	return *g.player3.Team
}

func (g Game) score() Score {
	return NewScore(g.rounds)
}
