package game

import "fmt"

const (
	noCurrentTrickError = iota + 1
	noCurrentRoundError
	noTrickWinnerError
	noRoundWinTeamError
	tooManyCardsPerTrickError
	tooManyTricksPerRoundError
	tooManyRoundsPerMatchError
	invalidRankError
	invalidSuitError
	handNotFoundError
	invalidCardForPlayError
	unexpectedPlayerError
	noTrumpAssignedError
	repeatedTrumpAssignmentError
	invalidTrumpError
	unexpectedTrumperError
	matchCompletedError
)

type matchErrorCode int

type matchError struct {
	code matchErrorCode
	msg  string
}

func (me matchError) Error() string {
	return me.msg
}

func newNoCurrentTrickError() matchError {
	return matchError{noCurrentTrickError, "current trick not present"}
}

func newNoCurrentRoundError() matchError {
	return matchError{noCurrentRoundError, "current round not present"}
}

func newNoTrickWinnerError() matchError {
	return matchError{
		noTrickWinnerError, "trick has no winner",
	}
}

func newNoRoundWinTeamError() matchError {
	return matchError{
		noRoundWinTeamError, "round has no winning team",
	}
}

func newTooManyCardsPerTrickError() matchError {
	return matchError{
		tooManyCardsPerTrickError, "too many cards per trick",
	}
}

func newTooManyTricksPerRoundError() matchError {
	return matchError{
		tooManyTricksPerRoundError, "too many tricks per round",
	}
}

func newTooManyRoundsPerMatchError() matchError {
	return matchError{
		tooManyRoundsPerMatchError, "too many rounds per match",
	}
}

func newInvalidRankError(rank Rank) matchError {
	return matchError{
		invalidRankError,
		fmt.Sprintf("rank %s is invalid", rank),
	}
}

func newInvalidSuitError(suit Suit) matchError {
	return matchError{
		invalidSuitError,
		fmt.Sprintf("suit %s is invalid", suit),
	}
}

func newHandNotFoundError(player Player) matchError {
	return matchError{
		handNotFoundError, fmt.Sprintf("player %s hand not found", player),
	}
}

func newInvalidCardForPlayError(player Player, card Card) matchError {
	return matchError{
		invalidCardForPlayError,
		fmt.Sprintf("player %s cannot play card %s", player, card),
	}
}

func newNoTrumpAssignedError() matchError {
	return matchError{
		noTrumpAssignedError, "trump not assigned",
	}
}

func newUnexpectedPlayerError(actual Player, expected Player) matchError {
	return matchError{
		unexpectedPlayerError,
		fmt.Sprintf(
			"expected player %s to play, not %s", expected, actual,
		),
	}
}

func newRepeatedTrumpAssignmentError(trump Suit, currentTrump Suit) matchError {
	return matchError{
		repeatedTrumpAssignmentError,
		fmt.Sprintf(
			"cannot assign trump %s: already assigned to %s",
			trump, currentTrump,
		),
	}
}

func newInvalidTrumpError(trump Suit) matchError {
	return matchError{
		invalidTrumpError,
		fmt.Sprintf("trump %s is invalid", trump),
	}
}

func newUnexpectedTrumperError(actual Player, expected Player) matchError {
	return matchError{
		unexpectedTrumperError,
		fmt.Sprintf(
			"%s cannot assign trump: %s must do it", actual, expected,
		),
	}
}

func newMatchCompletedError() matchError {
	return matchError{
		matchCompletedError, "match is completed",
	}
}
