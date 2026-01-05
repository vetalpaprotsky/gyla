package game

const (
	// After this event, players must see what move a player made.
	playerMovedEvent = "player_moved"

	// After this event, players must see what suit is a trump now.
	trumpChosenEvent = "trump_chosen"

	// After this event, round starter must choose a trump.
	roundStartedEvent = "round_started"

	// After this event, players must see who won the round.
	roundCompletedEvent = "round_completed"

	// After this event, trick starter must make the first move.
	trickStartedEvent = "trick_started"

	// After this event, players must see which player won a trick.
	trickCompletedEvent = "trick_completed"

	// After this event, players must see a welcome message.
	gameStartedEvent = "game_started"

	// After this event, players must see which team won the game.
	gameCompletedEvent = "game_completed"
)

type Event struct {
	Name         string
	gameSnapshot gameSnapshot
}

func (e Event) getGameSnapshotFor(p Player) GameSnapshot {
	return GameSnapshot{}
}
