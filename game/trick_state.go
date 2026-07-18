package game

type TrickState struct {
	Number      int
	Next        Player
	PlayedCards map[Player]Card
	Winner      Player
}

func newTrickState(t trick) TrickState {
	return TrickState{
		Number:      t.number,
		Next:        t.expectedNextPlayer(),
		PlayedCards: t.playedCards,
		Winner:      t.winner(),
	}
}
