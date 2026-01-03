package models

import (
	"errors"
	"fmt"
)

type Round struct {
	Number  int
	Hands   []Hand
	Tricks  []Trick
	Trump   Suit
	starter Player
	plrsRel PlayersRelation
}

// TODO: When score of the starter team goes over 30, then their teammate must
// start the next round. One player can't choose trumps always.
func newRound(curRound Round) (Round, error) {
	if curRound.Number >= maxPossibleNumberOfRounds {
		// NOTE: This is note expected error. Panic?
		msg := "Max possible number of rounds started. Can't start a new one."
		return Round{}, errors.New(msg)
	}

	winTeam, winTeamOk := curRound.winTeam()
	if !winTeamOk {
		// NOTE: This is note expected error. Panic?
		msg := "Current round doesn't have a win team. Can't start a new one."
		return Round{}, errors.New(msg)
	}

	round := Round{
		plrsRel: curRound.plrsRel,
		Number:  curRound.Number + 1,
		Hands:   newDeck().deal(curRound.plrsRel),
	}

	if winTeam == curRound.starterTeam() {
		round.starter = curRound.starter
	} else {
		round.starter = curRound.starterLeftOpponent()
	}

	return round, nil
}

func newFirstRound(pr PlayersRelation) Round {
	round := Round{plrsRel: pr, Number: 1, Hands: newDeck().deal(pr)}

	if starter, ok := round.findPlayerWithNineOfDiamonds(); ok {
		round.starter = starter
	} else {
		panic("Player with nine of diamonds isn't found.")
	}

	return round
}

func (r *Round) startNextTrick() (*Trick, error) {
	var trick Trick
	var err error

	if curTrick := r.CurrentTrick(); curTrick == nil {
		trick = newFirstTrick(r.starter, r.plrsRel)
	} else {
		trick, err = newTrick(*curTrick)
	}

	if err != nil {
		return nil, err
	}

	r.Tricks = append(r.Tricks, trick)

	return &r.Tricks[len(r.Tricks)-1], nil
}

func (r *Round) assignTrump(suit Suit) error {
	if r.Trump != "" {
		errorMsg := fmt.Sprintf(
			"Can't assign <%s> trump, it's already assigned to <%s>.",
			suit,
			r.Trump,
		)
		return errors.New(errorMsg)
	}

	trumpIsValid := false
	for _, validSuit := range ValidSuits {
		if suit == validSuit {
			trumpIsValid = true
		}
	}
	if !trumpIsValid {
		msg := fmt.Sprintf("Trump <%s> is invalid, can't assign it.", suit)
		return errors.New(msg)
	}

	r.Trump = suit
	for _, h := range r.Hands {
		for i, c := range h.Cards {
			if c.Suit == suit {
				h.Cards[i].IsTrump = true
			}
		}
	}

	return nil
}

func (r *Round) makeMove(player Player, card Card) error {
	trick := r.CurrentTrick()
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

func (r Round) CurrentTrick() *Trick {
	if len(r.Tricks) == 0 {
		return nil
	}

	trick := &r.Tricks[0]
	for i := 1; i < len(r.Tricks); i++ {
		if r.Tricks[i].Number > trick.Number {
			trick = &r.Tricks[i]
		}
	}

	return trick
}

func (r Round) getHand(player Player) *Hand {
	for i := 0; i < len(r.Hands); i++ {
		if r.Hands[i].Player == player {
			return &r.Hands[i]
		}
	}

	return nil
}

func (r Round) TricksPerTeam() TricksPerTeam {
	return newTricksPerTeam(r)
}

func (r Round) findPlayerWithNineOfDiamonds() (Player, bool) {
	for i := 0; i < len(r.Hands); i++ {
		hand := r.Hands[i]

		for j := 0; j < len(hand.Cards); j++ {
			card := hand.Cards[j]

			if card.Rank == NineRank && card.Suit == DiamondsSuit {
				return hand.Player, true
			}
		}
	}

	return Player(""), false
}

func (r Round) availableCardsForMove(player Player) []Card {
	trick := *r.CurrentTrick()
	hand := r.getHand(player)

	return hand.availableCardsForMove(trick)
}

func (r Round) IsCompleted() bool {
	if len(r.Tricks) != tricksPerRoundCount {
		return false
	}

	for _, trick := range r.Tricks {
		if !trick.IsCompleted() {
			return false
		}
	}

	return true
}

func (r Round) winTeam() (Team, bool) {
	if !r.IsCompleted() {
		return Team(""), false
	}

	starterTeam := r.starterTeam()
	opponentTeam := r.starterOpponentTeam()
	starterTeamTricks := 0
	opponentTeamTricks := 0

	for _, trick := range r.Tricks {
		// It's safe to skip bool value in this case, since we're sure that
		// winner is present, since all tricks are have a winner at this point.
		winner, _ := trick.Winner()

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

func (r Round) starterTeam() Team {
	team := r.plrsRel.getTeam(r.starter)

	if team == Team("") {
		panic(fmt.Sprintf("Starter player <%s> team is missing.", r.starter))
	}

	return team
}

func (r Round) starterOpponentTeam() Team {
	team := r.plrsRel.getOpponentTeam(r.starter)

	if team == Team("") {
		panic(fmt.Sprintf("Starter player <%s> opponent team is missing.", r.starter))
	}

	return r.plrsRel.getOpponentTeam(r.starter)
}

func (r Round) starterLeftOpponent() Player {
	opponent := r.plrsRel.getLeftOpponent(r.starter)

	if opponent == Player("") {
		panic(fmt.Sprintf("Starter player <%s> left opponent is missing.", r.starter))
	}

	return opponent
}
