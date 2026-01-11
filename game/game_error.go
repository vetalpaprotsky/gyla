package game

import "fmt"

const (
	noCurrentTrickError            = "no_current_trick"
	noCurrentRoundError            = "no_current_round"
	handNotFoundError              = "hand_not_found"
	invalidCardForMoveError        = "invalid_card_for_move"
	trickCompletedError            = "trick_completed"
	invalidPlayerError             = "invalid_player"
	duplicatedPlayerMoveError      = "duplicated_player_move"
	unexpectedPlayerMoveError      = "unexpected_player_move"
	tooManyTricksPerRoundError     = "too_many_tricks_per_round"
	noTrickWinnerError             = "no_trick_winner"
	tooManyRoundsPerGameError      = "too_many_rounds_per_game"
	noRoundWinTeamError            = "no_round_win_team"
	duplicatedTrumpAssignmentError = "duplicated_trump_assignment"
	invalidTrumpError              = "invalid_trump"
	unexpectedTrumperError         = "unexpected_trumper"
	invalidRankError               = "invalid_rank"
	invalidSuitError               = "invalid_suit"
)

type gameError struct {
	error_type string
	msg        string
}

func (ge gameError) Error() string {
	return fmt.Sprintf("%s - %s", ge.error_type, ge.msg)
}

func newNoCurrentRoundError() gameError {
	return gameError{noCurrentRoundError, "current round isn't present."}
}

func newNoCurrentTrickError() gameError {
	return gameError{noCurrentTrickError, "current trick isn't present."}
}

func newHandNotFoundError(player Player) gameError {
	return gameError{
		handNotFoundError, fmt.Sprintf("Player <%s> hand isn't found.", player),
	}
}

func newInvalidCardForMoveError(player Player, card Card) gameError {
	return gameError{
		invalidCardForMoveError,
		fmt.Sprintf("Player <%s> can't make a move with <%s> card.", player, card),
	}
}

func newTrickCompletedError() gameError {
	return gameError{
		trickCompletedError, "Trick is completed. Can't add a new card to it.",
	}
}

func newInvalidPlayerError(player Player) gameError {
	return gameError{
		invalidPlayerError, fmt.Sprintf("Player <%s> is invalid.", player),
	}
}

func newDuplicatedPlayerMoveError(player Player, card Card) gameError {
	return gameError{
		duplicatedPlayerMoveError,
		fmt.Sprintf("Player <%s> already made a move with <%s>.", player, card),
	}
}

func newUnexpectedPlayerMoveError(actual Player, expected Player) gameError {
	return gameError{
		unexpectedPlayerMoveError,
		fmt.Sprintf(
			"Player <%s> is expected to make a move, not <%s>.", expected, actual,
		),
	}
}

func newTooManyTricksPerRoundError() gameError {
	return gameError{
		tooManyTricksPerRoundError, "Too many tricks per round.",
	}
}

func newNoTrickWinnerError() gameError {
	return gameError{
		noTrickWinnerError, "Trick has no winner.",
	}
}

func newTooManyRoundsPerGameError() gameError {
	return gameError{
		tooManyRoundsPerGameError, "Too many rounds per game.",
	}
}

func newNoRoundWinTeamError() gameError {
	return gameError{
		noRoundWinTeamError, "Round has no win team.",
	}
}

func newDuplicatedTrumpAssignmentError(trump Suit, currentTrump Suit) gameError {
	return gameError{
		duplicatedTrumpAssignmentError,
		fmt.Sprintf(
			"Can't assign <%s> trump, it's already assigned to <%s>.",
			trump, currentTrump,
		),
	}
}

func newInvalidTrumpError(trump Suit) gameError {
	return gameError{
		invalidTrumpError,
		fmt.Sprintf("Trump <%s> is invalid, can't assign it.", trump),
	}
}

func newUnexpectedTrumperError(actual Player, expected Player) gameError {
	return gameError{
		unexpectedTrumperError,
		fmt.Sprintf(
			"<%s> can't assign trump. <%s> must do it", actual, expected,
		),
	}
}

func newInvalidRankError(rank Rank) gameError {
	return gameError{
		invalidRankError,
		fmt.Sprintf("Rank <%s> is invalid.", rank),
	}
}

func newInvalidSuitError(suit Suit) gameError {
	return gameError{
		invalidSuitError,
		fmt.Sprintf("Suit <%s> is invalid.", suit),
	}
}
