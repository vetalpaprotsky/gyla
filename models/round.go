package models

import (
	"errors"
	"fmt"
)

type Round struct {
	Number   int
	Hands    []Hand
	Tricks   []Trick
	Trump    Suit
	starter  Player
	relation PlayersRelation
}

// TODO: When score of the starter(the one which chooses a trump) goes over
// 30, then its teammate must start the next round. One player can't choose
// a trump always.
func newRound(curRound Round) (Round, error) {
	if curRound.Number >= maxPossibleNumberOfRounds {
		msg := "Max possible number of rounds started. Can't start a new one."
		return Round{}, errors.New(msg)
	}

	winTeam, winTeamOk := curRound.winTeam()
	if !winTeamOk {
		msg := "Current round doesn't have a win team. Can't start a new one."
		return Round{}, errors.New(msg)
	}

	round := Round{
		relation: curRound.relation,
		Number:   curRound.Number + 1,
		Hands:    newDeck().deal(curRound.relation),
	}

	if winTeam == curRound.starterTeam() {
		round.starter = curRound.starter
	} else {
		round.starter = curRound.starterLeftOpponent()
	}

	return round, nil
}

func newFirstRound(pr PlayersRelation) Round {
	round := Round{relation: pr, Number: 1, Hands: newDeck().deal(pr)}

	if starter, ok := round.findPlayerWithNineOfDiamonds(); ok {
		round.starter = starter
	} else {
		panic("Player with nine of diamonds isn't found")
	}

	return round
}

func (r *Round) CurrentTrick() *Trick {
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

func (r *Round) startNextTrick() (*Trick, error) {
	var trick Trick
	var err error

	if curTrick := r.CurrentTrick(); curTrick == nil {
		trick = newFirstTrick(r.starter)
	} else {
		trick, err = newTrick(*curTrick)
	}

	if err != nil {
		return nil, err
	}

	r.Tricks = append(r.Tricks, trick)

	return &r.Tricks[len(r.Tricks)-1], nil
}

func (r *Round) getHand(player Player) *Hand {
	for i := 0; i < len(r.Hands); i++ {
		if r.Hands[i].Player == player {
			return &r.Hands[i]
		}
	}

	return nil
}

func (r *Round) TricksPerTeam() TricksPerTeam {
	return newTricksPerTeam(*r)
}

func (r *Round) assignTrump(suit Suit) error {
	if r.Trump != "" {
		errorMsg := fmt.Sprintf(
			"Can't assign %s Trump, it's already assigned to %s",
			suit,
			r.Trump,
		)
		return errors.New(errorMsg)
	}

	r.Trump = suit

	for i := 0; i < len(r.Hands); i++ {
		r.Hands[i].assignTrump(suit)
	}

	return nil
}

func (r *Round) findPlayerWithNineOfDiamonds() (Player, bool) {
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

func (r *Round) makeMove(player Player, card Card) error {
	hand := r.getHand(player)

	if hand == nil {
		msg := fmt.Sprintf("Player <%s> hand isn't found. Can't make move.", player)
		return errors.New(msg)
	}

	trick := r.CurrentTrick()
	if trick == nil {
		return errors.New("No current trick. Can't make move.")
	}

	if err := hand.makeMove(card, trick); err != nil {
		return err
	}

	return nil
}

func (r *Round) availableCardsForMove(player Player) []Card {
	trick := *r.CurrentTrick()
	hand := r.getHand(player)

	return hand.availableCardsForMove(trick)
}

func (r *Round) IsCompleted() bool {
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

func (r *Round) winTeam() (Team, bool) {
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

		switch winnerTeam := r.relation.getTeam(winner); winnerTeam {
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

func (r *Round) starterTeam() Team {
	team := r.relation.getTeam(r.starter)

	if team == Team("") {
		panic(fmt.Sprintf("Starter player <%s> team is missing", r.starter))
	}

	return team
}

func (r *Round) starterOpponentTeam() Team {
	team := r.relation.getOpponentTeam(r.starter)

	if team == Team("") {
		panic(fmt.Sprintf("Starter player <%s> opponent team is missing", r.starter))
	}

	return r.relation.getOpponentTeam(r.starter)
}

func (r *Round) starterLeftOpponent() Player {
	opponent := r.relation.getLeftOpponent(r.starter)

	if opponent == Player("") {
		panic(fmt.Sprintf("Starter player <%s> left opponent is missing", r.starter))
	}

	return opponent
}
