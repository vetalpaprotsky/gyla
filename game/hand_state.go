package game

type HandState struct {
	Player Player
	Cards  []HandCard
}

func newHandState(h hand, t trick) HandState {
	state := HandState{Player: h.player}
	playableCards := h.playableCardsFor(t)

	for _, c := range h.cards {
		isPlayable := false

		for _, pc := range playableCards {
			if c.id() == pc.id() {
				isPlayable = true
				break
			}
		}

		state.Cards = append(state.Cards, c.asHandCard(isPlayable))
	}

	return state
}
