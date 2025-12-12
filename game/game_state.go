package game

type Team string
type Player string
type Rank string
type Suit string

type Card struct {
	Rank    string
	Suit    string
	IsTrump bool
}

type Trick struct {
	Player1Move Card
	Player2Move Card
	Player3Move Card
	Player4Move Card

	Starter Player
	Winner  Player
}

type GameState struct {
	Team1   Team
	Player1 Player
	Player3 Player

	Team2   Team
	Player2 Player
	Player4 Player

	Team1Score int
	Team2Score int
	WinnerTeam Team

	Player1Bot bool
	Player2Bot bool
	Player3Bot bool
	Player4Bot bool

	Player1Cards []Card
	Player2Cards []Card
	Player3Cards []Card
	Player4Cards []Card
	Trump        Suit

	RoundNumber  int
	RoundStarter Player

	Team1Tricks  []Trick
	Team2Tricks  []Trick
	Trick        Trick
	TrickStarter Player
}
