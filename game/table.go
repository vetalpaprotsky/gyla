package game

type Player string
type Team string

type table struct {
	players []Player
	teams   map[Player]Team
	ai      map[Player]bool
}

func newTable(p1, p2, p3, p4 Player, t1, t2 Team, ai1, ai2, ai3, ai4 bool) table {
	return table{
		players: []Player{p1, p2, p3, p4},
		teams: map[Player]Team{
			p1: t1,
			p3: t1,
			p2: t2,
			p4: t2,
		},
		ai: map[Player]bool{
			p1: ai1,
			p2: ai2,
			p3: ai3,
			p4: ai4,
		},
	}
}

func (t table) getTeam(p Player) Team {
	return t.teams[p]
}

func (t table) getOpponentTeam(p Player) Team {
	return t.getTeam(t.getLeftOpponent(p))
}

func (t table) getLeftOpponent(p Player) Player {
	for i, player := range t.players {
		if player == p {
			return t.players[(i+1)%4]
		}
	}
	return Player("")
}

func (t table) getTeammate(p Player) Player {
	for i, player := range t.players {
		if player == p {
			return t.players[(i+2)%4]
		}
	}
	return Player("")
}

func (t table) getRightOpponent(p Player) Player {
	for i, player := range t.players {
		if player == p {
			return t.players[(i+3)%4]
		}
	}
	return Player("")
}

func (t table) getTeamPlayers(team Team) []Player {
	var players []Player
	for _, p := range t.players {
		if t.teams[p] == team {
			players = append(players, p)
		}
	}
	return players
}

func (t table) getAllPlayers() []Player {
	return t.players[:]
}

func (t table) isAI(p Player) bool {
	return t.ai[p]
}

func (t table) getTeams() []Team {
	team1 := t.teams[t.players[0]]
	team2 := t.teams[t.players[1]]
	return []Team{team1, team2}
}
