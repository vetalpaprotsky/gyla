package game

type Player string
type Team string

type playersRelation struct {
	team1   Team
	player1 Player
	player3 Player
	team2   Team
	player2 Player
	player4 Player
}

func (pr playersRelation) getTeam(p Player) Team {
	switch p {
	case pr.player1, pr.player3:
		return pr.team1
	case pr.player2, pr.player4:
		return pr.team2
	default:
		return Team("")
	}
}

func (pr playersRelation) getOpponentTeam(p Player) Team {
	return pr.getTeam(pr.getLeftOpponent(p))
}

func (pr playersRelation) getLeftOpponent(p Player) Player {
	switch p {
	case pr.player1:
		return pr.player2
	case pr.player2:
		return pr.player3
	case pr.player3:
		return pr.player4
	case pr.player4:
		return pr.player1
	default:
		return Player("")
	}
}

func (pr playersRelation) getTeammate(p Player) Player {
	switch p {
	case pr.player1:
		return pr.player3
	case pr.player2:
		return pr.player4
	case pr.player3:
		return pr.player1
	case pr.player4:
		return pr.player2
	default:
		return Player("")
	}
}

func (pr playersRelation) getRightOpponent(p Player) Player {
	switch p {
	case pr.player1:
		return pr.player4
	case pr.player2:
		return pr.player1
	case pr.player3:
		return pr.player2
	case pr.player4:
		return pr.player3
	default:
		return Player("")
	}
}

func (pr playersRelation) getTeamPlayers(t Team) []Player {
	switch t {
	case pr.team1:
		return []Player{pr.player1, pr.player3}
	case pr.team2:
		return []Player{pr.player2, pr.player4}
	default:
		return []Player{}
	}
}

func (pr playersRelation) getAllPlayers() []Player {
	return []Player{pr.player1, pr.player2, pr.player3, pr.player4}
}
