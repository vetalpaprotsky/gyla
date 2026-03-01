package game

// Lifecycle:
// 1. Start next round.
// 2. Assign trump.
// 3. Start next trick.
// 4. Play card.
// 5. If current trick isn't completed, go to step 4.
// 6. If current round isn't completed, go to step 3.
// 7. If match isn't completed, go to step 1.
// 8. Match is completed.
type match struct {
	table  Table
	rounds []round
}

func (m *match) startNextRound() error {
	if m.isMatchCompleted() {
		return newMatchCompletedError()
	}

	var round round
	var err error

	if curRound := m.currentRound(); curRound == nil {
		round = newFirstRound(m.table)
	} else {
		round, err = newRound(*curRound)
	}

	if err != nil {
		return err
	}

	m.rounds = append(m.rounds, round)
	return nil
}

func (m *match) startNextTrick() error {
	if m.isMatchCompleted() {
		return newMatchCompletedError()
	}

	curRound := m.currentRound()
	if curRound == nil {
		return newNoCurrentRoundError()
	}

	if err := curRound.startNextTrick(); err != nil {
		return err
	}

	return nil
}

func (m *match) playCard(rank Rank, suit Suit, player Player) error {
	if m.isMatchCompleted() {
		return newMatchCompletedError()
	}

	curRound := m.currentRound()
	if curRound == nil {
		return newNoCurrentRoundError()
	}

	if err := curRound.playCard(rank, suit, player); err != nil {
		return err
	}

	return nil
}

func (m *match) assignTrump(suit Suit, player Player) error {
	if m.isMatchCompleted() {
		return newMatchCompletedError()
	}

	curRound := m.currentRound()
	if curRound == nil {
		return newNoCurrentRoundError()
	}

	if err := curRound.assignTrump(suit, player); err != nil {
		return err
	}

	return nil
}

func (m match) currentRound() *round {
	if len(m.rounds) == 0 {
		return nil
	}

	return &m.rounds[len(m.rounds)-1]
}

func (m match) currentTrick() *trick {
	curRound := m.currentRound()
	if curRound == nil {
		return nil
	}

	return curRound.currentTrick()
}

func (m match) isCurrentTrickCompleted() bool {
	curTrick := m.currentTrick()
	if curTrick == nil {
		return false
	}

	return curTrick.isCompleted()
}

func (m match) isCurrentRoundCompleted() bool {
	curRound := m.currentRound()
	if curRound == nil {
		return false
	}

	return curRound.isCompleted()
}

func (m match) isMatchCompleted() bool {
	return m.isCurrentRoundCompleted() && newMatchStats(m).isMatchCompleted()
}

func (m match) state() MatchState {
	curRound := m.currentRound()
	stats := newMatchStats(m)

	if curRound == nil {
		return MatchState{
			Table: m.table,
			Stats: stats,
		}
	}

	return MatchState{
		Table: m.table,
		Round: m.currentRound().state(),
		Stats: stats,
	}
}

type MatchState struct {
	Table Table
	Round RoundState
	Stats MatchStats
}

func (ms MatchState) ViewFor(p Player) MatchView {
	return MatchView{
		Table: ms.Table.ViewFor(p),
		Round: ms.Round.ViewFor(p),
		Stats: ms.Stats,
	}
}

type MatchView struct {
	Table TableView
	Round RoundView
	Stats MatchStats
}
