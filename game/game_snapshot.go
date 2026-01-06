package game

type gameSnapshot struct {
	round   round
	score   score
	plrsRel playersRelation
}

func (gs gameSnapshot) getGameSnapshotFor(p Player) GameSnapshotForPlayer {
	// TODO

	return GameSnapshotForPlayer{}
}

// Snapshot from player point of view. It doesn't reveal non intended details,
// like other player cards or tricks.
type GameSnapshotForPlayer struct {
	// Get initialized on "game_started".
	Team              Team   // Immutable after initialization.
	OpponentTeam      Team   // Immutable after initialization.
	TeamScore         int    // Can change after "round_completed".
	OpponentTeamScore int    // Can change after "round_completed".
	Player            Player // Immutable after initialization.
	LeftOpponent      Player // Immutable after initialization.
	Teammate          Player // Immutable after initialization.
	RightOpponent     Player // Immutable after initialization.
	LeftOpponentBot   bool   // Immutable after initialization.
	TeammateBot       bool   // Immutable after initialization.
	RightOpponentBot  bool   // Immutable after initialization.

	// Get initialized on "round_started".
	RoundNumber              int     // Immutable after initialization.
	Trumper                  Player  // Immutable after initialization.
	TrumperHasSix            bool    // Immutable after initialization.
	Cards                    []Card  // Can change after "player_moved".
	LeftOpponentCardsCount   int     // Can change after "player_moved".
	TeammateCardsCount       int     // Can change after "player_moved".
	RightOpponentCardsCount  int     // Can change after "player_moved".
	Tricks                   []Trick // Can change after "trick_completed".
	LeftOpponentTricksCount  int     // Can change after "trick_completed".
	TeammateTricks           []Trick // Can change after "trick_completed".
	RightOpponentTricksCount int     // Can change after "trick_completed".

	// Gets initialized on "trump_chosen".
	Trump Suit // Immutable after initialization.

	// Gets initialized on "trick_started".
	CurrentTrick Trick // Changes after "player_moved" or "trick_completed".

	// Gets initialized on "player_moved".
	LastMove Move // Changes after "player_moved".

	// Gets initialized on "round_completed".
	RoundWinnerTeam Team // Immutable after initialization.

	// Gets initialized on "game_completed".
	GameWinnerTeam Team // Immutable after initialization.

	// If it's set, an action is expected from a player.
	ExpectedNextAction ExpectedAction
}

// This lifecycle gets applied only to CurrentTrick field.
type Trick struct {
	// Gets initialized on "trick_started".
	Starter Player // Immutable after initialization.

	// One field at a time get initialized on "player_moved".
	Player        Card // Immutable after initialization.
	LeftOpponent  Card // Immutable after initialization.
	Teammate      Card // Immutable after initialization.
	RightOpponent Card // Immutable after initialization.

	// Gets initialized on "trick_completed".
	Winner Player // Immutable after initialization.
}

type Move struct {
	Player Player
	Card   Card
}
