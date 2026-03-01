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

type MatchEvent struct {
	EventType  EventType
	MatchState MatchState
}

func newMatchEvent(m match, et EventType) MatchEvent {
	return MatchEvent{
		EventType:  et,
		MatchState: m.state(),
	}
}
