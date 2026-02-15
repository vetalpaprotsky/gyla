package game

import (
	"fmt"
)

// TODO: Hands must be implemented in a different way. Consider map instead of
// slice.
type round struct {
	number  int
	hands   []hand
	tricks  []trick
	trump   Suit
	starter Player
	plrsRel playersRelation
}

func (r round) deepCopy() round {
	hands := make([]hand, 0, len(r.hands))
	for _, h := range r.hands {
		hands = append(hands, h.deepCopy())
	}

	tricks := make([]trick, 0, len(r.tricks))
	for _, t := range r.tricks {
		tricks = append(tricks, t.deepCopy())
	}

	return round{
		number:  r.number,
		hands:   hands,
		tricks:  tricks,
		trump:   r.trump,
		starter: r.starter,
		plrsRel: r.plrsRel,
	}
}

// TODO: When score of the starter team goes over 30, then their teammate must
// start the next round. One player can't assign trumps always.
func newRound(curRound round) (round, error) {
	if curRound.number >= maxPossibleNumberOfRounds {
		return round{}, newTooManyRoundsPerMatchError()
	}

	winTeam, winTeamOk := curRound.winTeam()
	if !winTeamOk {
		return round{}, newNoRoundWinTeamError()
	}

	round := round{
		plrsRel: curRound.plrsRel,
		number:  curRound.number + 1,
		hands:   newDeck().deal(curRound.plrsRel),
	}

	if winTeam == curRound.starterTeam() {
		round.starter = curRound.starter
	} else {
		round.starter = curRound.starterLeftOpponent()
	}

	return round, nil
}

func newFirstRound(pr playersRelation) round {
	round := round{plrsRel: pr, number: 1, hands: newDeck().deal(pr)}

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
			trick = newFirstTrick(r.starter, r.plrsRel)
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

func (r round) isTrumpAssigned() bool {
	return r.trump != ""
}

func (r round) trumper() Player {
	return r.starter
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
		for i, c := range h.cards {
			if c.Suit == trump {
				h.cards[i].IsTrump = true
			}
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

	trick := &r.tricks[0]
	for i := 1; i < len(r.tricks); i++ {
		if r.tricks[i].number > trick.number {
			trick = &r.tricks[i]
		}
	}

	return trick
}

func (r round) getHand(player Player) *hand {
	for i := 0; i < len(r.hands); i++ {
		if r.hands[i].player == player {
			return &r.hands[i]
		}
	}

	return nil
}

func (r round) findPlayerWithNineOfDiamonds() (Player, bool) {
	for i := 0; i < len(r.hands); i++ {
		hand := r.hands[i]

		for j := 0; j < len(hand.cards); j++ {
			card := hand.cards[j]

			if card.Rank == nineRank && card.Suit == diamondsSuit {
				return hand.player, true
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

func (r round) winTeam() (Team, bool) {
	if !r.isCompleted() {
		return Team(""), false
	}

	starterTeam := r.starterTeam()
	opponentTeam := r.starterOpponentTeam()
	starterTeamTricks := 0
	opponentTeamTricks := 0

	for _, trick := range r.tricks {
		// It's safe to skip bool value in this case, since we're sure that
		// winner is present, since all tricks are have a winner at this point.
		winner, _ := trick.winner()

		switch winnerTeam := r.plrsRel.getTeam(winner); winnerTeam {
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
		return starterTeam, true
	} else if starterTeamTricks < opponentTeamTricks {
		return opponentTeam, true
	} else {
		panic("draw in round: impossible case")
	}
}

func (r round) trumperHasSix() bool {
	for _, c := range r.getHand(r.trumper()).cards {
		if c.Rank == sixRank && c.IsTrump {
			return true
		}
	}

	return false
}

func (r round) starterTeam() Team {
	t := r.plrsRel.getTeam(r.starter)

	if t == Team("") {
		panic(fmt.Sprintf("starter player %s team is missing", r.starter))
	}

	return t
}

func (r round) starterOpponentTeam() Team {
	t := r.plrsRel.getOpponentTeam(r.starter)

	if t == Team("") {
		panic(fmt.Sprintf("starter player %s opponent team is missing", r.starter))
	}

	return r.plrsRel.getOpponentTeam(r.starter)
}

func (r round) starterLeftOpponent() Player {
	opponent := r.plrsRel.getLeftOpponent(r.starter)

	if opponent == Player("") {
		panic(fmt.Sprintf("starter player %s left opponent is missing", r.starter))
	}

	return opponent
}
