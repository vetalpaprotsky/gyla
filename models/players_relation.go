package models

type PlayersRelation struct {
	Team1   Team
	Player1 Player
	Player3 Player
	Team2   Team
	Player2 Player
	Player4 Player
}

func (pr PlayersRelation) getTeam(p Player) Team {
	team := Team("")

	if p == pr.Player1 || p == pr.Player3 {
		team = pr.Team1
	} else if p == pr.Player2 || p == pr.Player4 {
		team = pr.Team2
	}

	return team
}

func (pr PlayersRelation) getOpponentTeam(p Player) Team {
	switch currentTeam := pr.getTeam(p); currentTeam {
	case pr.Team1:
		return pr.Team2
	case pr.Team2:
		return pr.Team1
	default:
		return Team("")
	}
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

func (pr PlayersRelation) allPlayers() []Player {
	return []Player{pr.Player1, pr.Player2, pr.Player3, pr.Player4}
}
