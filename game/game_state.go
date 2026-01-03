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

	CurrentTrick             Trick
	Tricks                   []Trick
	LeftOpponentTricksCount  int
	TeammateTricks           []Trick
	RightOpponentTricksCount int

	// Server expects this player to choose a trump when "round_started" event is sent.
	Trumper       Player
	TrumperHasSix bool
	Trump         Suit

	RoundNumber int

	// Gets set by server when "round_completed" event is sent.
	RoundWinnerTeam Team

	// Gets set by server when "game_completed" event is sent.
	GameWinnerTeam Team

	// When this value is non zero, server expects "move" action from a client.
	ExpextingNextMoveBy Player

	// When this set to true, server expects "trump_choice" action from a client.
	ExpextingToChooseTrump bool
}
