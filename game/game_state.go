package game

type GameState struct {
	Round        RoundState
	Stats        GameStats
	Participants []Participant
}

func newGameState(g Game) GameState {
	curRound := g.currentRound()
	var curRoundState RoundState

	if curRound != nil {
		curRoundState = curRound.state()
	}

	return GameState{
		Round:        curRoundState,
		Stats:        g.stats.deepCopy(),
		Participants: append([]Participant{}, g.participants...),
	}
}

func (gs GameState) getParticipant(p Player) Participant {
	for _, participant := range gs.Participants {
		if participant.Player == p {
			return participant
		}
	}

	return Participant{}
}

func (gs GameState) ViewFor(p Player) GameView {
	return GameView{
		You:           gs.getParticipant(p),
		LeftOpponent:  gs.getParticipant(p.leftOpponent()),
		Teammate:      gs.getParticipant(p.teammate()),
		RightOpponent: gs.getParticipant(p.rightOpponent()),

		Round: gs.Round.ViewFor(p),
		Stats: gs.Stats,
	}
}

type GameView struct {
	You           Participant
	LeftOpponent  Participant
	Teammate      Participant
	RightOpponent Participant

	Round RoundView
	Stats GameStats
}
