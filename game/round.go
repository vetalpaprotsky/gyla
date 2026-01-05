package game

import (
	"errors"
	"fmt"
)

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
// start the next round. One player can't choose trumps always.
func newRound(curRound round) (round, error) {
	if curRound.number >= maxPossibleNumberOfRounds {
		// NOTE: This is note expected error. Panic?
		msg := "Max possible number of rounds started. Can't start a new one."
		return round{}, errors.New(msg)
	}

	winTeam, winTeamOk := curRound.winTeam()
	if !winTeamOk {
		// NOTE: This is note expected error. Panic?
		msg := "Current round doesn't have a win team. Can't start a new one."
		return round{}, errors.New(msg)
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
		panic("Player with nine of diamonds isn't found.")
	}

	return round
}

func (r *round) startNextTrick() (*trick, error) {
	var trick trick
	var err error

	if curTrick := r.currentTrick(); curTrick == nil {
		trick = newFirstTrick(r.starter, r.plrsRel)
	} else {
		trick, err = newTrick(*curTrick)
	}

	if err != nil {
		return nil, err
	}

	r.tricks = append(r.tricks, trick)

	return &r.tricks[len(r.tricks)-1], nil
}

func (r *round) assignTrump(suit Suit) error {
	if r.trump != "" {
		errorMsg := fmt.Sprintf(
			"Can't assign <%s> trump, it's already assigned to <%s>.",
			suit,
			r.trump,
		)
		return errors.New(errorMsg)
	}

	trumpIsValid := false
	for _, validSuit := range validSuits {
		if suit == validSuit {
			trumpIsValid = true
		}
	}
	if !trumpIsValid {
		msg := fmt.Sprintf("Trump <%s> is invalid, can't assign it.", suit)
		return errors.New(msg)
	}

	r.trump = suit
	for _, h := range r.hands {
		for i, c := range h.cards {
			if c.Suit == suit {
				h.cards[i].IsTrump = true
			}
		}
	}

	return nil
}

func (r *round) makeMove(player Player, card Card) error {
	trick := r.currentTrick()
	if trick == nil {
		return errors.New("No current trick. Can't make move.")
	}

	hand := r.getHand(player)
	if hand == nil {
		msg := fmt.Sprintf("Player <%s> hand isn't found. Can't make move.", player)
		return errors.New(msg)
	}

	if !hand.canMakeMove(card, *trick) {
		msg := fmt.Sprintf(
			"Player <%s> can't make a move with <%s> card.", player, card,
		)
		return errors.New(msg)
	}

	if err := trick.addCard(player, card); err != nil {
		return err
	}

	if ok := hand.removeCard(card); !ok {
		panic("Could not remove card from a hand when making a move.")
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

func (r round) tricksPerTeam() tricksPerTeam {
	return newTricksPerTeam(r)
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

func (r round) availableCardsForMove(player Player) []Card {
	trick := *r.currentTrick()
	hand := r.getHand(player)

	return hand.availableCardsForMove(trick)
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
				"Team <%s> with Player <%s> doesn't exist.",
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
		panic("Draw in a round. Impossible case.")
	}
}

func (r round) starterTeam() Team {
	t := r.plrsRel.getTeam(r.starter)

	if t == Team("") {
		panic(fmt.Sprintf("Starter player <%s> team is missing.", r.starter))
	}

	return t
}

func (r round) starterOpponentTeam() Team {
	t := r.plrsRel.getOpponentTeam(r.starter)

	if t == Team("") {
		panic(fmt.Sprintf("Starter player <%s> opponent team is missing.", r.starter))
	}

	return r.plrsRel.getOpponentTeam(r.starter)
}

func (r round) starterLeftOpponent() Player {
	opponent := r.plrsRel.getLeftOpponent(r.starter)

	if opponent == Player("") {
		panic(fmt.Sprintf("Starter player <%s> left opponent is missing.", r.starter))
	}

	return opponent
}
