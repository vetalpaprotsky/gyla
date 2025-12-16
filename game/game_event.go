package game

// Clients must store events in the FIFO queue on their side, and read them
// one by one. Server will expect actions to be sent after some events, but not all.

const (
	// After this event, players must see what move a player made.
	PlayerMovedEvent = "player_moved"

	// After this event, players must see what suit is a trump now.
	TrumpChosenEvent = "trump_chosen"

	// After this event, round starter must choose a trump.
	RoundStartedEvent = "round_started"

	// After this event, players must see who won the round.
	RoundCompletedEvent = "round_completed"

	// After this event, trick starter must made the first move.
	TrickStartedEvent = "trick_started"

	// After this event, players must see which player won a trick.
	TrickCompletedEvent = "trick_completed"

	// After this event, players must see a welcome message.
	GameStartedEvent = "game_started"

	// After this event, players must see which team won the game.
	GameCompletedEvent = "game_completed"
)

type GameEvent struct {
	GameState GameState
	Event     string
}
