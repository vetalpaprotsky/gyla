package game

// Maybe a different approach is needed?

type gameSnapshot struct {
	round   round
	score   score
	plrsRel playersRelation
}

type GameSnapshot struct {
	Team              Team
	OpponentTeam      Team
	TeamScore         int
	OpponentTeamScore int

	Player        Player
	LeftOpponent  Player
	Teammate      Player
	RightOpponent Player

	LeftOpponentBot  bool
	TeammateBot      bool
	RightOpponentBot bool

	Cards                   []Card
	LeftOpponentCardsCount  int
	TeammateCardsCount      int
	RightOpponentCardsCount int

	CurrentTrick             Trick
	Tricks                   []Trick
	LeftOpponentTricksCount  int
	TeammateTricks           []Trick
	RightOpponentTricksCount int

	RoundNumber   int
	Trumper       Player
	TrumperHasSix bool
	Trump         Suit

	ExpectedNextAction   string
	ExpectedNextActionBy Player

	// Gets set by server when "round_completed" event is sent.
	RoundWinnerTeam Team

	// Gets set by server when "game_completed" event is sent.
	GameWinnerTeam Team
}
