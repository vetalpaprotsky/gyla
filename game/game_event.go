package game

// Events are ordered by their lifecycle.
const (
	GameStartedEvent                 = "game_started"
	RoundStartedEvent                = "round_started"
	TrumpAssignedEvent               = "trump_assigned"
	TrickStartedEvent                = "trick_started"
	CardPlayedEvent                  = "card_played"
	CardPlayedAndTrickCompletedEvent = "card_played_and_trick_completed"
	CardPlayedAndRoundCompletedEvent = "card_played_and_round_completed"
	CardPlayedAndGameCompletedEvent  = "card_played_and_game_completed"
)

type EventType string

type GameEvent struct {
	EventType EventType
	GameState GameState
}
