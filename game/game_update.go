package game

type GameUpdate struct {
	State  GameState
	Events []GameEvent
}

func (gu *GameUpdate) addState(g *Game) {
	gu.State = NewGameState(g)
}

func (gu *GameUpdate) addEvent(eventType string, g *Game) {
	gu.Events = append(
		gu.Events,
		NewGameEvent(eventType, g),
	)
}
