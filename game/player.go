package game

type Player int

type PlayerInfo struct {
	Player     Player
	TeamInfo   TeamInfo
	Name       string
	IsAI       bool
	ExternalID string
}

const (
	Player1 = Player(1)
	Player2 = Player(2)
	Player3 = Player(3)
	Player4 = Player(4)
)

var allPlayers = [4]Player{
	Player1,
	Player2,
	Player3,
	Player4,
}

func (p Player) isZero() bool {
	return p == Player(0)
}

func (p Player) team() Team {
	switch p {
	case Player1, Player3:
		return Team1
	case Player2, Player4:
		return Team2
	}

	return Team(0)
}

func (p Player) opponentTeam() Team {
	return p.team().opponent()
}

func (p Player) leftOpponent() Player {
	switch p {
	case Player1:
		return Player2
	case Player2:
		return Player3
	case Player3:
		return Player4
	case Player4:
		return Player1
	}

	return Player(0)
}

func (p Player) teammate() Player {
	switch p {
	case Player1:
		return Player3
	case Player2:
		return Player4
	case Player3:
		return Player1
	case Player4:
		return Player2
	}

	return Player(0)
}

func (p Player) rightOpponent() Player {
	switch p {
	case Player1:
		return Player4
	case Player2:
		return Player1
	case Player3:
		return Player2
	case Player4:
		return Player3
	}

	return Player(0)
}
