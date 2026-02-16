package game

import "fmt"

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

type GameEvent interface {
	EventType() string
	PayloadFor(Player) any
}

func NewGameEvent(eventType string, g *Game) GameEvent {
	switch eventType {
	case matchStartedEvent:
		return MatchStartedEvent{plrsRel: g.match.plrsRel}
	case roundStartedEvent:
		return RoundStartedEvent{round: g.match.currentRound().deepCopy()}
	case trumpAssignedEvent:
		return TrumpAssignedEvent{round: g.match.currentRound().deepCopy()}
	// TODO
	default:
		panic(fmt.Sprintf("invalid event type: %s", eventType))
	}
}

/******************************************************************************/

type MatchStartedEvent struct {
	plrsRel playersRelation
}

func (MatchStartedEvent) EventType() string {
	return matchStartedEvent
}

func (e MatchStartedEvent) PayloadFor(p Player) any {
	return MatchStartedEventPayload{
		Team:          e.plrsRel.getTeam(p),
		OpponentTeam:  e.plrsRel.getOpponentTeam(p),
		You:           p,
		LeftOpponent:  e.plrsRel.getLeftOpponent(p),
		Teammate:      e.plrsRel.getTeammate(p),
		RightOpponent: e.plrsRel.getRightOpponent(p),
		// TODO: e.plrsRel.getAIMap()
		// AI:            e.ai,
	}
}

type MatchStartedEventPayload struct {
	Team          Team
	OpponentTeam  Team
	You           Player
	LeftOpponent  Player
	Teammate      Player
	RightOpponent Player
	AI            map[Player]bool
}

/******************************************************************************/

type RoundStartedEvent struct {
	round round
}

func (RoundStartedEvent) EventType() string {
	return roundStartedEvent
}

func (e RoundStartedEvent) PayloadFor(p Player) any {
	round := e.round
	plrsRel := round.plrsRel

	return RoundStartedEventPayload{
		RoundNumber: round.number,
		Trumper:     round.trumper(),
		YourCards:   round.getHand(p).cards,
		CardsPerPlayer: map[Player]int{
			plrsRel.player1: len(round.getHand(plrsRel.player1).cards),
			plrsRel.player2: len(round.getHand(plrsRel.player2).cards),
			plrsRel.player3: len(round.getHand(plrsRel.player3).cards),
			plrsRel.player4: len(round.getHand(plrsRel.player4).cards),
		},
	}
}

type RoundStartedEventPayload struct {
	RoundNumber    int
	Trumper        Player
	YourCards      []Card
	CardsPerPlayer map[Player]int
}

/******************************************************************************/

type TrumpAssignedEvent struct {
	round round
}

func (TrumpAssignedEvent) EventType() string {
	return trumpAssignedEvent
}

func (e TrumpAssignedEvent) PayloadFor(p Player) any {
	return TrumpAssignedEventPayload{
		Trump:         e.round.trump,
		YourCards:     e.round.getHand(p).cards,
		TrumperHasSix: e.round.trumperHasSix(),
	}
}

type TrumpAssignedEventPayload struct {
	Trump         Suit
	YourCards     []Card
	TrumperHasSix bool
}

/******************************************************************************/

type TrickStartedPayload struct {
	Number  int
	Starter Player
}

/******************************************************************************/

type CardPlayedPaylaod struct {
	Player Player
	Card   Card
	Next   Player
}

/******************************************************************************/

type TrickCompletedPayload struct {
	Winner Player
}

/******************************************************************************/

type RoundCompletedPayload struct {
	WinTeam           Team
	TeamScore         int
	OpponentTeamScore int
}

/******************************************************************************/

type MatchCompletedPaylaod struct {
	WinTeam Team
}
