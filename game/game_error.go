package game

import "fmt"

const (
	noTrickWinnerError = iota + 1
	noRoundWinTeamError
	tooManyCardsPerTrickError
	tooManyTricksPerRoundError
	tooManyRoundsPerGameError
	invalidRankError
	invalidSuitError
	handNotFoundError
	invalidCardForPlayError
	unexpectedPlayerError
	noTrumpAssignedError
	repeatedTrumpAssignmentError
	invalidTrumpError
	unexpectedTrumperError
	gameAlreadyStartedError
	gameCompletedError
)

type gameErrorCode int

type gameError struct {
	code gameErrorCode
	msg  string
}

func (ge gameError) Error() string {
	return ge.msg
}

func newNoTrickWinnerError() gameError {
	return gameError{
		noTrickWinnerError, "trick has no winner",
	}
}

func newNoRoundWinTeamError() gameError {
	return gameError{
		noRoundWinTeamError, "round has no winning team",
	}
}

func newTooManyCardsPerTrickError() gameError {
	return gameError{
		tooManyCardsPerTrickError, "too many cards per trick",
	}
}

func newTooManyTricksPerRoundError() gameError {
	return gameError{
		tooManyTricksPerRoundError, "too many tricks per round",
	}
}

func newTooManyRoundsPerGameError() gameError {
	return gameError{
		tooManyRoundsPerGameError, "too many rounds per game",
	}
}

func newInvalidRankError(rank Rank) gameError {
	return gameError{
		invalidRankError,
		fmt.Sprintf("rank %v is invalid", rank),
	}
}

func newInvalidSuitError(suit Suit) gameError {
	return gameError{
		invalidSuitError,
		fmt.Sprintf("suit %v is invalid", suit),
	}
}

func newHandNotFoundError(player Player) gameError {
	return gameError{
		handNotFoundError, fmt.Sprintf("player %v hand not found", player),
	}
}

func newInvalidCardForPlayError(player Player, card Card) gameError {
	return gameError{
		invalidCardForPlayError,
		fmt.Sprintf("player %v cannot play card %v", player, card),
	}
}

func newNoTrumpAssignedError() gameError {
	return gameError{
		noTrumpAssignedError, "trump not assigned",
	}
}

func newUnexpectedPlayerError(actual Player, expected Player) gameError {
	return gameError{
		unexpectedPlayerError,
		fmt.Sprintf(
			"expected player %v to play, not %v", expected, actual,
		),
	}
}

func newRepeatedTrumpAssignmentError(trump Suit, currentTrump Suit) gameError {
	return gameError{
		repeatedTrumpAssignmentError,
		fmt.Sprintf(
			"cannot assign trump %v: already assigned to %v",
			trump, currentTrump,
		),
	}
}

func newInvalidTrumpError(trump Suit) gameError {
	return gameError{
		invalidTrumpError,
		fmt.Sprintf("trump %v is invalid", trump),
	}
}

func newUnexpectedTrumperError(actual Player, expected Player) gameError {
	return gameError{
		unexpectedTrumperError,
		fmt.Sprintf(
			"%v cannot assign trump: %v must do it", actual, expected,
		),
	}
}

func newGameAlreadyStartedError() gameError {
	return gameError{
		gameAlreadyStartedError, "game already started",
	}
}

func newGameCompletedError() gameError {
	return gameError{
		gameCompletedError, "game is completed",
	}
}
