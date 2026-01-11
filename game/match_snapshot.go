package game

type matchSnapshot struct {
	// RENAME: currentRound
	round   round
	score   score
	plrsRel playersRelation
}

func (gs matchSnapshot) getMatchSnapshotFor(p Player) MatchSnapshotForPlayer {
	// TODO

	return MatchSnapshotForPlayer{}
}

// Snapshot from player point of view. It doesn't reveal non intended details,
// like other player cards or tricks.
type MatchSnapshotForPlayer struct {
	// Get initialized on "match_started".
	Team              Team
	OpponentTeam      Team
	TeamScore         int // Can change after "round_completed".
	OpponentTeamScore int // Can change after "round_completed".
	Player            Player
	LeftOpponent      Player
	Teammate          Player
	RightOpponent     Player
	LeftOpponentBot   bool
	TeammateBot       bool
	RightOpponentBot  bool

	// Get initialized on "round_started".
	RoundNumber              int
	Trumper                  Player
	TrumperHasSix            bool
	Cards                    []Card  // Can change after "card_played".
	LeftOpponentCardsCount   int     // Can change after "card_played".
	TeammateCardsCount       int     // Can change after "card_played".
	RightOpponentCardsCount  int     // Can change after "card_played".
	Tricks                   []Trick // Can change after "trick_completed".
	LeftOpponentTricksCount  int     // Can change after "trick_completed".
	TeammateTricks           []Trick // Can change after "trick_completed".
	RightOpponentTricksCount int     // Can change after "trick_completed".

	// Gets initialized on "trump_assigned".
	Trump Suit

	// Gets initialized on "trick_started".
	CurrentTrick Trick // Changes after "card_played" or "trick_completed".

	// Gets initialized on "round_completed".
	RoundWinnerTeam Team

	// Gets initialized on "match_completed".
	MatchWinnerTeam Team

	// If it's set, an action is expected from a player.
	ExpectedNextAction ExpectedAction
}

// This lifecycle gets applied only to CurrentTrick field.
type Trick struct {
	// Gets initialized on "trick_started".
	Starter Player

	// One field at a time get initialized on "card_played".
	Player        Card
	LeftOpponent  Card
	Teammate      Card
	RightOpponent Card

	// Gets initialized on "trick_completed".
	Winner Player
}
