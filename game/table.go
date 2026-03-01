package game

type Table struct {
	Seat1 Seat
	Seat2 Seat
	Seat3 Seat
	Seat4 Seat
	Team1 Team
	Team2 Team
}

func newTable(p1, p2, p3, p4 Player, t1, t2 Team, ai1, ai2, ai3, ai4 bool) Table {
	return Table{
		Seat1: Seat{Player: p1, IsAI: ai1},
		Seat2: Seat{Player: p2, IsAI: ai2},
		Seat3: Seat{Player: p3, IsAI: ai3},
		Seat4: Seat{Player: p4, IsAI: ai4},
		Team1: t1,
		Team2: t2,
	}
}

func (t Table) getTeam(p Player) Team {
	switch p {
	case t.Seat1.Player, t.Seat3.Player:
		return t.Team1
	case t.Seat2.Player, t.Seat4.Player:
		return t.Team2
	default:
		return Team(0)
	}
}

func (t Table) getOpponentTeam(p Player) Team {
	return t.getTeam(t.getLeftOpponent(p))
}

func (t Table) getLeftOpponent(p Player) Player {
	switch p {
	case t.Seat1.Player:
		return t.Seat2.Player
	case t.Seat2.Player:
		return t.Seat3.Player
	case t.Seat3.Player:
		return t.Seat4.Player
	case t.Seat4.Player:
		return t.Seat1.Player
	}
	return Player(0)
}

func (t Table) getTeammate(p Player) Player {
	switch p {
	case t.Seat1.Player:
		return t.Seat3.Player
	case t.Seat2.Player:
		return t.Seat4.Player
	case t.Seat3.Player:
		return t.Seat1.Player
	case t.Seat4.Player:
		return t.Seat2.Player
	}
	return Player(0)
}

func (t Table) getRightOpponent(p Player) Player {
	switch p {
	case t.Seat1.Player:
		return t.Seat4.Player
	case t.Seat2.Player:
		return t.Seat1.Player
	case t.Seat3.Player:
		return t.Seat2.Player
	case t.Seat4.Player:
		return t.Seat3.Player
	}
	return Player(0)
}

func (t Table) getSeat(p Player) Seat {
	switch p {
	case t.Seat1.Player:
		return t.Seat1
	case t.Seat2.Player:
		return t.Seat2
	case t.Seat3.Player:
		return t.Seat3
	case t.Seat4.Player:
		return t.Seat4
	}
	return Seat{}
}

func (t Table) isAI(p Player) bool {
	switch p {
	case t.Seat1.Player:
		return t.Seat1.IsAI
	case t.Seat2.Player:
		return t.Seat2.IsAI
	case t.Seat3.Player:
		return t.Seat3.IsAI
	case t.Seat4.Player:
		return t.Seat4.IsAI
	}
	return false
}

func (t Table) players() []Player {
	return []Player{
		t.Seat1.Player,
		t.Seat2.Player,
		t.Seat3.Player,
		t.Seat4.Player,
	}
}

func (t Table) ViewFor(p Player) TableView {
	return TableView{
		You:           t.getSeat(p),
		LeftOpponent:  t.getSeat(t.getLeftOpponent(p)),
		Teammate:      t.getSeat(t.getTeammate(p)),
		RightOpponent: t.getSeat(t.getRightOpponent(p)),

		Team:         t.getTeam(p),
		OpponentTeam: t.getOpponentTeam(p),
	}
}

type Player int

func (p Player) IsZero() bool {
	return p == 0
}

type Team int

func (t Team) IsZero() bool {
	return t == 0
}

type TableView struct {
	You           Seat
	LeftOpponent  Seat
	Teammate      Seat
	RightOpponent Seat

	Team         Team
	OpponentTeam Team
}

type Seat struct {
	Player Player
	IsAI   bool
}
