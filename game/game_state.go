package game

type GameState struct {
	match match
}

// NOTE: If we need to optimise this, we could store only info that is needed,
// not the whole match copy. But for now let's ignore this.
func newGameState(g *Game) GameState {
	return GameState{match: g.match.deepCopy()}
}

type PlayerView struct {
	Team          Team
	OpponentTeam  Team
	You           Player
	LeftOpponent  Player
	Teammate      Player
	RightOpponent Player
	AI            map[Player]bool

	RoundNumber    int
	Trumper        Player
	TrumpedWithSix bool
	Trump          Suit

	YourCards      []Card
	CardsPerPlayer map[Player]int

	TricksHistory []Trick
	CurrentTrick  Trick

	TeamPoints         int
	OpponentTeamPoints int

	RoundWinTeam Team
	MatchWinTeam Team
}

type Trick struct {
	Number int
	Next   Player
	Cards  map[Player]Card
	Winner Player
}

func (gs GameState) ViewFor(p Player) PlayerView {
	curRound := gs.match.currentRound()
	table := gs.match.table

	if curRound == nil {
		return PlayerView{
			Team:          table.getTeam(p),
			OpponentTeam:  table.getOpponentTeam(p),
			You:           p,
			LeftOpponent:  table.getLeftOpponent(p),
			Teammate:      table.getTeammate(p),
			RightOpponent: table.getRightOpponent(p),
			AI:            table.ai,
		}
	}

	score := newScore(gs.match)

	return PlayerView{
		Team:          table.getTeam(p),
		OpponentTeam:  table.getOpponentTeam(p),
		You:           p,
		LeftOpponent:  table.getLeftOpponent(p),
		Teammate:      table.getTeammate(p),
		RightOpponent: table.getRightOpponent(p),
		AI:            table.ai,

		RoundNumber:    curRound.number,
		Trumper:        curRound.trumper(),
		TrumpedWithSix: curRound.trumpedWithSix,
		Trump:          curRound.trump,

		YourCards:      curRound.getHand(p).cards,
		CardsPerPlayer: cardsPerPlayer(*curRound),

		TricksHistory: tricksHistory(*curRound),
		CurrentTrick:  currentTrick(*curRound),

		TeamPoints:         score.points[table.getTeam(p)],
		OpponentTeamPoints: score.points[table.getOpponentTeam(p)],

		RoundWinTeam: curRound.winTeam(),
		MatchWinTeam: score.winTeam(),
	}
}

func cardsPerPlayer(r round) map[Player]int {
	result := make(map[Player]int)

	for _, h := range r.hands {
		result[h.player] = len(h.cards)
	}

	return result
}

func tricksHistory(r round) []Trick {
	curTrick := r.currentTrick()
	if curTrick == nil {
		return nil
	}

	var tricks []Trick
	for _, t := range r.tricks {
		if t.number != curTrick.number {
			tricks = append(tricks, toTrick(t))
		}
	}

	return tricks
}

func currentTrick(r round) Trick {
	curTrick := r.currentTrick()
	if curTrick == nil {
		return Trick{}
	}

	return toTrick(*curTrick)
}

func toTrick(t trick) Trick {
	return Trick{
		Number: t.number,
		Next:   t.expectedNextPlayer(),
		Cards:  t.cards,
		Winner: t.winner(),
	}
}
