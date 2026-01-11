package game

// TODO: Knows nothing about events or actions.
// Basically it works like Game struct used to work.
// It shoud have simple methods like startNextRound()
// startNextTrick(), isTrickCompleted(), isRoundCompleted(),
// isMatchCompleted(), assignTrumpForRound(), playCard(), ...
type match struct {
	rounds  []round
	plrsRel playersRelation
}
