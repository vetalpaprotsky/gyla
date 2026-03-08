package game

type TrickState struct {
	Number      int
	Next        Player
	PlayedCards []PlayedCard
	Winner      Player
}

func newTrickState(t trick) TrickState {
	return TrickState{
		Number:      t.number,
		Next:        t.expectedNextPlayer(),
		PlayedCards: append([]PlayedCard{}, t.playedCards...),
		Winner:      t.winner(),
	}
}
