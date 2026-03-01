package game

import "fmt"

const maxPossibleNumberOfRounds = 19

type round struct {
	number         int
	starter        Player
	trumpedWithSix bool
	trump          Suit
	hands          []Hand
	tricks         []trick
	table          Table
}

// TODO: When score of the starter team goes over 30, then their teammate must
// start the next round. One player can't assign trumps always.
func newRound(curRound round) (round, error) {
	if curRound.number >= maxPossibleNumberOfRounds {
		return round{}, newTooManyRoundsPerMatchError()
	}

	winTeam := curRound.winTeam()
	if winTeam == Team("") {
		return round{}, newNoRoundWinTeamError()
	}

	round := round{
		table:  curRound.table,
		number: curRound.number + 1,
		hands:  newDeck().deal(curRound.table),
		tricks: make([]trick, 0, tricksPerRoundCount),
	}

	if winTeam == curRound.starterTeam() {
		round.starter = curRound.starter
	} else {
		round.starter = curRound.starterLeftOpponent()
	}

	return round, nil
}

func newFirstRound(t Table) round {
	round := round{table: t, number: 1, hands: newDeck().deal(t)}

	if starter, ok := round.findPlayerWithNineOfDiamonds(); ok {
		round.starter = starter
	} else {
		panic("player with nine of diamonds not found")
	}

	return round
}

func (r *round) startNextTrick() error {
	var trick trick
	var err error

	if curTrick := r.currentTrick(); curTrick == nil {
		if r.isTrumpAssigned() {
			trick = newFirstTrick(r.starter, r.table)
		} else {
			err = newNoTrumpAssignedError()
		}
	} else {
		trick, err = newTrick(*curTrick)
	}

	if err != nil {
		return err
	}

	r.tricks = append(r.tricks, trick)
	return nil
}

func (r *round) assignTrump(trump Suit, player Player) error {
	if r.isTrumpAssigned() {
		return newRepeatedTrumpAssignmentError(trump, r.trump)
	}

	trumpIsValid := false
	for _, validSuit := range validSuits {
		if trump == validSuit {
			trumpIsValid = true
		}
	}
	if !trumpIsValid {
		return newInvalidTrumpError(trump)
	}

	if player != r.trumper() {
		return newUnexpectedTrumperError(player, r.trumper())
	}

	r.trump = trump

	for _, h := range r.hands {
		for i, c := range h.Cards {
			if c.Suit == trump {
				h.Cards[i].IsTrump = true
			}
		}
	}

	for _, c := range r.getHand(r.trumper()).Cards {
		if c.Rank == SixRank && c.IsTrump {
			r.trumpedWithSix = true
		}
	}

	return nil
}

func (r *round) playCard(rank Rank, suit Suit, player Player) error {
	trick := r.currentTrick()
	if trick == nil {
		return newNoCurrentTrickError()
	}

	hand := r.getHand(player)
	if hand == nil {
		return newHandNotFoundError(player)
	}

	card := hand.getCard(rank, suit)
	if !hand.canPlayCard(card, *trick) {
		return newInvalidCardForPlayError(player, card)
	}

	if err := trick.addCard(player, card); err != nil {
		return err
	}

	if ok := hand.removeCard(card); !ok {
		panic("could not remove card from hand")
	}

	return nil
}

func (r round) currentTrick() *trick {
	if len(r.tricks) == 0 {
		return nil
	}

	return &r.tricks[len(r.tricks)-1]
}

func (r round) getHand(player Player) *Hand {
	for i := 0; i < len(r.hands); i++ {
		if r.hands[i].Player == player {
			return &r.hands[i]
		}
	}

	return nil
}

func (r round) findPlayerWithNineOfDiamonds() (Player, bool) {
	for i := 0; i < len(r.hands); i++ {
		hand := r.hands[i]

		for j := 0; j < len(hand.Cards); j++ {
			card := hand.Cards[j]

			if card.Rank == NineRank && card.Suit == DiamondsSuit {
				return hand.Player, true
			}
		}
	}

	return Player(""), false
}

func (r round) playableCardsFor(player Player) []Card {
	trick := *r.currentTrick()
	hand := r.getHand(player)

	return hand.playableCardsFor(trick)
}

func (r round) isCompleted() bool {
	if len(r.tricks) != tricksPerRoundCount {
		return false
	}

	for _, trick := range r.tricks {
		if !trick.isCompleted() {
			return false
		}
	}

	return true
}

func (r round) winTeam() Team {
	if !r.isCompleted() {
		return Team("")
	}

	starterTeam := r.starterTeam()
	opponentTeam := r.starterOpponentTeam()
	starterTeamTricks := 0
	opponentTeamTricks := 0

	for _, trick := range r.tricks {
		winner := trick.winner()

		switch winnerTeam := r.table.getTeam(winner); winnerTeam {
		case starterTeam:
			starterTeamTricks += 1
		case opponentTeam:
			opponentTeamTricks += 1
		default:
			msg := fmt.Sprintf(
				"team %s with player %s does not exist",
				winner,
				winnerTeam,
			)
			panic(msg)
		}
	}

	if starterTeamTricks > opponentTeamTricks {
		return starterTeam
	} else if starterTeamTricks < opponentTeamTricks {
		return opponentTeam
	} else {
		panic("draw in round: impossible case")
	}
}

func (r round) isTrumpAssigned() bool {
	return r.trump != ""
}

func (r round) trumper() Player {
	return r.starter
}

func (r round) starterTeam() Team {
	team := r.table.getTeam(r.starter)

	if team == Team("") {
		panic(fmt.Sprintf("starter player %s team is missing", r.starter))
	}

	return team
}

func (r round) starterOpponentTeam() Team {
	team := r.table.getOpponentTeam(r.starter)

	if team == Team("") {
		panic(fmt.Sprintf("starter player %s opponent team is missing", r.starter))
	}

	return r.table.getOpponentTeam(r.starter)
}

func (r round) starterLeftOpponent() Player {
	opponent := r.table.getLeftOpponent(r.starter)

	if opponent == Player("") {
		panic(fmt.Sprintf("starter player %s left opponent is missing", r.starter))
	}

	return opponent
}

func (r round) state() RoundState {
	hands := make([]Hand, 0, len(r.hands)-1)
	for _, h := range r.hands {
		hands = append(hands, h.deepCopy())
	}

	tricks := make([]TrickState, 0, len(r.tricks)-1)
	for _, t := range r.tricks {
		tricks = append(tricks, t.state())
	}

	return RoundState{
		Number:         r.number,
		Trumper:        r.trumper(),
		TrumpedWithSix: r.trumpedWithSix,
		Trump:          r.trump,
		Hands:          hands,
		Tricks:         tricks,
		WinTeam:        r.winTeam(),
		Table:          r.table,
	}
}

type RoundState struct {
	Number         int
	Trumper        Player
	TrumpedWithSix bool
	Trump          Suit
	Hands          []Hand
	Tricks         []TrickState
	WinTeam        Team
	Table          Table
}

func (rs RoundState) getHand(p Player) Hand {
	for _, h := range rs.Hands {
		if h.Player == p {
			return h
		}
	}

	return Hand{}
}

func (rs RoundState) ViewFor(p Player) RoundView {
	return RoundView{
		Number:            rs.Number,
		Trumper:           rs.Trumper,
		TrumpedWithSix:    rs.TrumpedWithSix,
		Trump:             rs.Trump,
		Hand:              rs.getHand(p),
		LeftOpponentHand:  len(rs.getHand(rs.Table.getLeftOpponent(p)).Cards),
		TeammateHand:      len(rs.getHand(rs.Table.getTeammate(p)).Cards),
		RightOpponentHand: len(rs.getHand(rs.Table.getRightOpponent(p)).Cards),
		Tricks:            rs.Tricks,
		WinTeam:           rs.WinTeam,
	}
}

type RoundView struct {
	Number            int
	Trumper           Player
	TrumpedWithSix    bool
	Trump             Suit
	Hand              Hand
	LeftOpponentHand  int
	TeammateHand      int
	RightOpponentHand int
	Tricks            []TrickState
	WinTeam           Team
}
