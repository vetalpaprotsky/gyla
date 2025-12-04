package models

type Player string
type Team string

type PlayersRelation struct {
	Team1   Team
	Player1 Player
	Player3 Player
	Team2   Team
	Player2 Player
	Player4 Player
}

func (pr PlayersRelation) getTeam(p Player) Team {
	switch p {
	case pr.Player1, pr.Player3:
		return pr.Team1
	case pr.Player2, pr.Player4:
		return pr.Team2
	default:
		return Team("")
	}
}

func (pr PlayersRelation) getOpponentTeam(p Player) Team {
	return pr.getTeam(pr.getLeftOpponent(p))
}

func (pr PlayersRelation) getLeftOpponent(p Player) Player {
	switch p {
	case pr.Player1:
		return pr.Player2
	case pr.Player2:
		return pr.Player3
	case pr.Player3:
		return pr.Player4
	case pr.Player4:
		return pr.Player1
	default:
		return Player("")
	}
}

func (pr PlayersRelation) getTeammate(p Player) Player {
	switch p {
	case pr.Player1:
		return pr.Player3
	case pr.Player2:
		return pr.Player4
	case pr.Player3:
		return pr.Player1
	case pr.Player4:
		return pr.Player2
	default:
		return Player("")
	}
}

func (pr PlayersRelation) teamPlayers(t Team) []Player {
	switch t {
	case pr.Team1:
		return []Player{pr.Player1, pr.Player3}
	case pr.Team2:
		return []Player{pr.Player2, pr.Player4}
	default:
		return []Player{}
	}
}

func (pr PlayersRelation) allPlayers() []Player {
	return []Player{pr.Player1, pr.Player2, pr.Player3, pr.Player4}
}
