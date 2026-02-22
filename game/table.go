package game

// TODO: We could use int instead of string. There's no need to use
// string. This will optimize things a lot.
type Player string
type Team string

type table struct {
	team1 Team
	team2 Team

	player1 Player
	player2 Player
	player3 Player
	player4 Player

	ai1 bool
	ai2 bool
	ai3 bool
	ai4 bool
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
	default:
		return Player("")
	}
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
	default:
		return Player("")
	}
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
	default:
		return Player("")
	}
}

func (t table) getTeamPlayers(team Team) []Player {
	switch team {
	case t.team1:
		return []Player{t.player1, t.player3}
	case t.team2:
		return []Player{t.player2, t.player4}
	default:
		return []Player{}
	}
}

func (t table) getAllPlayers() []Player {
	return []Player{t.player1, t.player2, t.player3, t.player4}
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
	default:
		return false
	}
}
