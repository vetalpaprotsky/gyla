package game

// Events are ordered by their lifecycle.
const (
	// After this event, players must see a welcome message.
	matchStartedEvent = "match_started"

	// After this event, round starter must assign a trump.
	roundStartedEvent = "round_started"

	// After this event, players must see what suit is a trump now.
	trumpAssignedEvent = "trump_assigned"

	// After this event, trick starter must play first card.
	trickStartedEvent = "trick_started"

	// After this event, players must see what card was played.
	cardPlayedEvent = "card_played"

	// After this event, players must see which player won the trick.
	trickCompletedEvent = "trick_completed"

	// After this event, players must see which team won the round.
	roundCompletedEvent = "round_completed"

	// After this event, players must see which team won the match.
	matchCompletedEvent = "match_completed"
)

type Event struct {
	Name         string
	gameSnapshot gameSnapshot
}

func (e Event) getGameSnapshotFor(p Player) GameSnapshotForPlayer {
	return e.gameSnapshot.getGameSnapshotFor(p)
}
