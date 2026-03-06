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
}

// TODO: When score of the starter team goes over 30, then their teammate must
// start the next round. One player can't assign trumps always.
func newRound(curRound round) (round, error) {
	if curRound.number >= maxPossibleNumberOfRounds {
		return round{}, newTooManyRoundsPerGameError()
	}

	winTeam := curRound.winTeam()
	if winTeam.isZero() {
		return round{}, newNoRoundWinTeamError()
	}

	round := round{
		number: curRound.number + 1,
		hands:  newDeck().deal(),
		tricks: make([]trick, 0, tricksPerRoundCount),
	}

	if winTeam == curRound.starterTeam() {
		round.starter = curRound.starter
	} else {
		round.starter = curRound.starterLeftOpponent()
	}

	return round, nil
}

func newFirstRound() round {
	round := round{number: 1, hands: newDeck().deal()}

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
			trick = newFirstTrick(r.starter)
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

	if !trump.isValid() {
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

	if ok := hand.playCard(card); !ok {
		panic("could not play card")
	}

	return nil
}

func (r round) currentTrick() *trick {
	if len(r.tricks) == 0 {
		return nil
	}

	return &r.tricks[len(r.tricks)-1]
}

func (r round) mustCurrentTrick() *trick {
	curTrick := r.currentTrick()
	if curTrick == nil {
		panic("no current trick")
	}

	return curTrick
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

	return Player(0), false
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
		return Team(0)
	}

	starterTeam := r.starterTeam()
	opponentTeam := r.starterOpponentTeam()
	starterTeamTricks := 0
	opponentTeamTricks := 0

	for _, trick := range r.tricks {
		winner := trick.winner()

		switch winner.team() {
		case starterTeam:
			starterTeamTricks += 1
		case opponentTeam:
			opponentTeamTricks += 1
		default:
			msg := fmt.Sprintf(
				"team %v with player %v does not exist",
				winner.team(),
				winner,
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
	return r.trump.isZero()
}

func (r round) trumper() Player {
	return r.starter
}

func (r round) starterTeam() Team {
	team := r.starter.team()

	if team.isZero() {
		panic(fmt.Sprintf("starter player %v team is missing", r.starter))
	}

	return team
}

func (r round) starterOpponentTeam() Team {
	team := r.starter.opponentTeam()

	if team.isZero() {
		panic(fmt.Sprintf("starter player %v opponent team is missing", r.starter))
	}

	return team
}

func (r round) starterLeftOpponent() Player {
	opponent := r.starter.leftOpponent()

	if opponent.isZero() {
		panic(fmt.Sprintf("starter player %v left opponent is missing", r.starter))
	}

	return opponent
}

func (r round) state() RoundState {
	return newRoundState(r)
}
