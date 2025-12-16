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
	Player        Card
	LeftOpponent  Card
	Teammate      Card
	RightOpponent Card

	Starter Player
	Winner  Player
}

type GameState struct {
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

	NextMoveBy Player

	CurrentTrick             Trick
	Tricks                   []Trick
	LeftOpponentTricksCount  int
	TeammateTricks           []Trick
	RightOpponentTricksCount int

	// Chooses trump after "round_started" event is sent.
	Trumper       Player
	Trump         Suit
	TrumperHasSix bool

	RoundNumber int

	// Must be set when "round_completed" event is sent.
	RoundWinnerTeam Team

	// Must be set when "game_completed" event is sent.
	GameWinnerTeam Team
}
