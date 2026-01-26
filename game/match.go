package game

// TODO: Knows nothing about events or actions.
// Basically it works like Game struct used to work.
type match struct {
	rounds           []round
	plrsRel          playersRelation
	isMatchCompleted bool
}

func (m *match) startNextRound() error {
	if m.isMatchCompleted {
		return newMatchCompletedError()
	}

	var round round
	var err error

	if curRound := m.currentRound(); curRound == nil {
		round = newFirstRound(m.plrsRel)
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
	if m.isMatchCompleted {
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
	if m.isMatchCompleted {
		return newMatchCompletedError()
	}

	curRound := m.currentRound()
	if curRound == nil {
		return newNoCurrentRoundError()
	}

	if err := curRound.playCard(rank, suit, player); err != nil {
		return err
	}

	if newScore(*m).isMatchCompleted() {
		m.isMatchCompleted = true
	}

	return nil
}

func (m *match) assignTrumpForCurrentRound(suit Suit, player Player) error {
	if m.isMatchCompleted {
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

	round := &m.rounds[0]
	for i := 1; i < len(m.rounds); i++ {
		if m.rounds[i].number > round.number {
			round = &m.rounds[i]
		}
	}

	return round
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

func (m match) createSnapshot() matchSnapshot {
	curRound := m.currentRound()

	if curRound == nil {
		return matchSnapshot{plrsRel: m.plrsRel}
	}

	return matchSnapshot{
		round:   curRound.deepCopy(),
		score:   newScore(m),
		plrsRel: m.plrsRel,
	}
}
