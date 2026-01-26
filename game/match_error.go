package game

import "fmt"

const (
	noCurrentTrickError = "no_current_trick"
	noCurrentRoundError = "no_current_round"

	noTrickWinnerError  = "no_trick_winner"
	noRoundWinTeamError = "no_round_win_team"

	tooManyCardsPerTrickError  = "too_many_cards_per_trick"
	tooManyTricksPerRoundError = "too_many_tricks_per_round"
	tooManyRoundsPerMatchError = "too_many_rounds_per_match"

	invalidRankError = "invalid_rank"
	invalidSuitError = "invalid_suit"

	handNotFoundError            = "hand_not_found"
	invalidCardForPlayError      = "invalid_card_for_play"
	unexpectedPlayerError        = "unexpected_player"
	repeatedTrumpAssignmentError = "repeated_trump_assignment"
	invalidTrumpError            = "invalid_trump"
	unexpectedTrumperError       = "unexpected_trumper"
	matchCompletedError          = "match_completed"
)

type matchError struct {
	error_type string
	msg        string
}

func (me matchError) Error() string {
	return fmt.Sprintf("%s - %s", me.error_type, me.msg)
}

func newNoCurrentTrickError() matchError {
	return matchError{noCurrentTrickError, "current trick isn't present."}
}

func newNoCurrentRoundError() matchError {
	return matchError{noCurrentRoundError, "current round isn't present."}
}

func newNoTrickWinnerError() matchError {
	return matchError{
		noTrickWinnerError, "Trick has no winner.",
	}
}

func newNoRoundWinTeamError() matchError {
	return matchError{
		noRoundWinTeamError, "Round has no win team.",
	}
}

func newTooManyCardsPerTrickError() matchError {
	return matchError{
		tooManyCardsPerTrickError, "Too many cards per trick.",
	}
}

func newTooManyTricksPerRoundError() matchError {
	return matchError{
		tooManyTricksPerRoundError, "Too many tricks per round.",
	}
}

func newTooManyRoundsPerMatchError() matchError {
	return matchError{
		tooManyRoundsPerMatchError, "Too many rounds per match.",
	}
}

func newInvalidRankError(rank Rank) matchError {
	return matchError{
		invalidRankError,
		fmt.Sprintf("Rank <%s> is invalid.", rank),
	}
}

func newInvalidSuitError(suit Suit) matchError {
	return matchError{
		invalidSuitError,
		fmt.Sprintf("Suit <%s> is invalid.", suit),
	}
}

func newHandNotFoundError(player Player) matchError {
	return matchError{
		handNotFoundError, fmt.Sprintf("Player <%s> hand isn't found.", player),
	}
}

func newInvalidCardForPlayError(player Player, card Card) matchError {
	return matchError{
		invalidCardForPlayError,
		fmt.Sprintf("Player <%s> can't play with <%s> card.", player, card),
	}
}

func newUnexpectedPlayerError(actual Player, expected Player) matchError {
	return matchError{
		unexpectedPlayerError,
		fmt.Sprintf(
			"Player <%s> is expected to play a card, not <%s>.", expected, actual,
		),
	}
}

func newRepeatedTrumpAssignmentError(trump Suit, currentTrump Suit) matchError {
	return matchError{
		repeatedTrumpAssignmentError,
		fmt.Sprintf(
			"Can't assign <%s> trump, it's already assigned to <%s>.",
			trump, currentTrump,
		),
	}
}

func newInvalidTrumpError(trump Suit) matchError {
	return matchError{
		invalidTrumpError,
		fmt.Sprintf("Trump <%s> is invalid, can't assign it.", trump),
	}
}

func newUnexpectedTrumperError(actual Player, expected Player) matchError {
	return matchError{
		unexpectedTrumperError,
		fmt.Sprintf(
			"<%s> can't assign trump. <%s> must do it", actual, expected,
		),
	}
}

func newMatchCompletedError() matchError {
	return matchError{
		matchCompletedError, "Match is completed.",
	}
}
