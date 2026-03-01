package game

type Player string
type Team string

type table struct {
	player1 Player
	player2 Player
	player3 Player
	player4 Player

	team1 Team
	team2 Team

	ai1 bool
	ai2 bool
	ai3 bool
	ai4 bool
}

func newTable(p1, p2, p3, p4 Player, t1, t2 Team, ai1, ai2, ai3, ai4 bool) table {
	return table{
		player1: p1,
		player2: p2,
		player3: p3,
		player4: p4,

		team1: t1,
		team2: t2,

		ai1: ai1,
		ai2: ai2,
		ai3: ai3,
		ai4: ai4,
	}
}

func (t table) getTeam(p Player) Team {
	switch p {
	case t.player1, t.player3:
		return t.team1
	case t.player2, t.player4:
		return t.team2
	default:
		return Team("")
	}
}

func (t table) getOpponentTeam(p Player) Team {
	return t.getTeam(t.getLeftOpponent(p))
}

func (t table) getLeftOpponent(p Player) Player {
	switch p {
	case t.player1:
		return t.player2
	case t.player2:
		return t.player3
	case t.player3:
		return t.player4
	case t.player4:
		return t.player1
	}
	return Player("")
}

func (t table) getTeammate(p Player) Player {
	switch p {
	case t.player1:
		return t.player3
	case t.player2:
		return t.player4
	case t.player3:
		return t.player1
	case t.player4:
		return t.player2
	}
	return Player("")
}

func (t table) getRightOpponent(p Player) Player {
	switch p {
	case t.player1:
		return t.player4
	case t.player2:
		return t.player1
	case t.player3:
		return t.player2
	case t.player4:
		return t.player3
	}
	return Player("")
}

func (t table) isAI(p Player) bool {
	switch p {
	case t.player1:
		return t.ai1
	case t.player2:
		return t.ai2
	case t.player3:
		return t.ai3
	case t.player4:
		return t.ai4
	}
	return false
}

func (t table) players() []Player {
	return []Player{t.player1, t.player2, t.player3, t.player4}
}
