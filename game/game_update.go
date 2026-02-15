package game

type GameUpdate struct {
	State  GameState
	Events []GameEvent
}

func NewGameUpdate(g *Game) GameUpdate {
	return GameUpdate{
		State: GameState{
			round: g.match.currentRound().deepCopy(),
			score: newScore(g.match),
			ai:    g.ai,
		},
		Events: g.events,
	}
}
