package game

type Player string

type Team string

type Table struct {
	Player1 Player
	Player2 Player
	Player3 Player
	Player4 Player

	Team1 Team
	Team2 Team

	AI1 bool
	AI2 bool
	AI3 bool
	AI4 bool
}

type TableView struct {
	You           Player
	LeftOpponent  Player
	Teammate      Player
	RightOpponent Player

	Team         Team
	OpponentTeam Team

	// YouAI - it doesn't make any sense :D
	LeftOpponentAI  bool
	TeammateAI      bool
	RightOpponentAI bool
}

func newTable(p1, p2, p3, p4 Player, t1, t2 Team, ai1, ai2, ai3, ai4 bool) Table {
	return Table{
		Player1: p1,
		Player2: p2,
		Player3: p3,
		Player4: p4,

		Team1: t1,
		Team2: t2,

		AI1: ai1,
		AI2: ai2,
		AI3: ai3,
		AI4: ai4,
	}
}

func (t Table) ViewFor(p Player) TableView {
	return TableView{
		You:           p,
		LeftOpponent:  t.getLeftOpponent(p),
		Teammate:      t.getTeammate(p),
		RightOpponent: t.getRightOpponent(p),

		Team:         t.getTeam(p),
		OpponentTeam: t.getOpponentTeam(p),

		LeftOpponentAI:  t.isAI(t.getLeftOpponent(p)),
		TeammateAI:      t.isAI(t.getTeammate(p)),
		RightOpponentAI: t.isAI(t.getRightOpponent(p)),
	}
}

func (t Table) getTeam(p Player) Team {
	switch p {
	case t.Player1, t.Player3:
		return t.Team1
	case t.Player2, t.Player4:
		return t.Team2
	default:
		return Team("")
	}
}

func (t Table) getOpponentTeam(p Player) Team {
	return t.getTeam(t.getLeftOpponent(p))
}

func (t Table) getLeftOpponent(p Player) Player {
	switch p {
	case t.Player1:
		return t.Player2
	case t.Player2:
		return t.Player3
	case t.Player3:
		return t.Player4
	case t.Player4:
		return t.Player1
	}
	return Player("")
}

func (t Table) getTeammate(p Player) Player {
	switch p {
	case t.Player1:
		return t.Player3
	case t.Player2:
		return t.Player4
	case t.Player3:
		return t.Player1
	case t.Player4:
		return t.Player2
	}
	return Player("")
}

func (t Table) getRightOpponent(p Player) Player {
	switch p {
	case t.Player1:
		return t.Player4
	case t.Player2:
		return t.Player1
	case t.Player3:
		return t.Player2
	case t.Player4:
		return t.Player3
	}
	return Player("")
}

func (t Table) isAI(p Player) bool {
	switch p {
	case t.Player1:
		return t.AI1
	case t.Player2:
		return t.AI2
	case t.Player3:
		return t.AI3
	case t.Player4:
		return t.AI4
	}
	return false
}

func (t Table) players() []Player {
	return []Player{t.Player1, t.Player2, t.Player3, t.Player4}
}
