package game

type GameState struct {
	round round
	score score
}

func newGameState(g *Game) GameState {
	return GameState{
		round: g.match.currentRound().deepCopy(),
		score: newScore(g.match),
	}
}

type GameStatePayload struct {
	Team          Team
	OpponentTeam  Team
	You           Player
	LeftOpponent  Player
	Teammate      Player
	RightOpponent Player
	AI            map[Player]bool

	RoundNumber   int
	Trumper       Player
	TrumperHasSix bool
	Trump         Suit

	YourCards      []Card
	CardsPerPlayer map[Player]int

	CompletedTricks []Trick
	CurrentTrick    Trick

	TeamScore         int
	OpponentTeamScore int

	RoundWinTeam Team
	MatchWinTeam Team
}

func (gs GameState) PayloadFor(p Player) GameStatePayload {
	return GameStatePayload{}
}
