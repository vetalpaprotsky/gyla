package game

// Events are ordered by their lifecycle.
const (
	MatchStartedEvent                = "match_started"
	RoundStartedEvent                = "round_started"
	TrumpAssignedEvent               = "trump_assigned"
	TrickStartedEvent                = "trick_started"
	CardPlayedEvent                  = "card_played"
	CardPlayedAndTrickCompletedEvent = "card_played_and_trick_completed"
	CardPlayedAndRoundCompletedEvent = "card_played_and_round_completed"
	CardPlayedAndMatchCompletedEvent = "card_played_and_match_completed"
)

type EventType string

type GameEvent struct {
	EventType EventType
	GameState GameState
}

func newGameEvent(g *Game, et EventType) GameEvent {
	return GameEvent{
		EventType: et,
		GameState: newGameState(g),
	}
}
