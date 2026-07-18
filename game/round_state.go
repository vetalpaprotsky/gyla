package game

type RoundState struct {
	Number          int
	Trumper         Player
	TrumpedWithSix  bool
	Trump           Suit
	Hands           []HandState
	CompletedTricks []TrickState
	CurrentTrick    TrickState
	WinTeam         Team
}

func newRoundState(r round) RoundState {
	curTrick := r.currentTrick()
	hands := make([]HandState, 0, len(r.hands))
	for _, h := range r.hands {
		hands = append(hands, h.state(curTrick))
	}

	completedTricks := make([]TrickState, 0, len(r.tricks))
	for _, t := range r.tricks {
		if t.isCompleted() {
			completedTricks = append(completedTricks, t.state())
		}
	}

	return RoundState{
		Number:          r.number,
		Trumper:         r.starter,
		TrumpedWithSix:  r.trumpedWithSix,
		Trump:           r.trump,
		Hands:           hands,
		CompletedTricks: completedTricks,
		CurrentTrick:    curTrick.state(),
		WinTeam:         r.winTeam(),
	}
}

func (rs RoundState) getHand(p Player) HandState {
	for _, h := range rs.Hands {
		if h.Player == p {
			return h
		}
	}

	return HandState{}
}

func (rs RoundState) ViewFor(p Player) RoundView {
	return RoundView{
		Number:            rs.Number,
		Trumper:           rs.Trumper,
		TrumpedWithSix:    rs.TrumpedWithSix,
		Trump:             rs.Trump,
		Hand:              rs.getHand(p),
		LeftOpponentHand:  len(rs.getHand(p.leftOpponent()).Cards),
		TeammateHand:      len(rs.getHand(p.teammate()).Cards),
		RightOpponentHand: len(rs.getHand(p.rightOpponent()).Cards),
		CompletedTricks:   rs.CompletedTricks,
		CurrentTrick:      rs.CurrentTrick,
		WinTeam:           rs.WinTeam,
	}
}

type RoundView struct {
	Number            int
	Trumper           Player
	TrumpedWithSix    bool
	Trump             Suit
	Hand              HandState
	LeftOpponentHand  int
	TeammateHand      int
	RightOpponentHand int
	CompletedTricks   []TrickState
	CurrentTrick      TrickState
	WinTeam           Team
}
