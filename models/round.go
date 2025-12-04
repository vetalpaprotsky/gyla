package models

import (
	"errors"
	"fmt"
	"math/rand"
)

type Round struct {
	Number   int
	Hands    []Hand
	Tricks   []Trick
	Trump    Suit
	starter  Player
	relation PlayersRelation
}

func newRound(curRound Round) (Round, error) {
	if curRound.Number >= maxPossibleNumberOfRounds {
		msg := "Max possible number of rounds already started."
		return Round{}, errors.New(msg)
	}

	if !curRound.IsCompleted() {
		msg := "Current round isn't completed. New round can't be started."
		return Round{}, errors.New(msg)
	}

	round := Round{relation: curRound.relation, Number: curRound.Number + 1}

	err := round.dealHands()
	if err != nil {
		return Round{}, err
	}

	if curRound.winnerTeam() == curRound.starterTeam() {
		round.starter = curRound.starter
	} else {
		round.starter = curRound.starterLeftOpponent()
	}

	return round, nil
}

func newFirstRound(pr PlayersRelation) (Round, error) {
	round := Round{relation: pr, Number: 1}

	err := round.dealHands()
	if err != nil {
		return Round{}, err
	}

	round.starter = round.findPlayerWithNineOfDiamonds()

	return round, nil
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

func (r *Round) startNextTrick() *Trick {
	curTrick := r.CurrentTrick()
	var trick Trick

	if curTrick == nil {
		trick = newFirstTrick(r.starter)
	} else {
		trick = newTrick(*curTrick)
	}

	r.Tricks = append(r.Tricks, trick)

	return &r.Tricks[len(r.Tricks)-1]
}

func (r *Round) getHand(player Player) *Hand {
	for i := 0; i < len(r.Hands); i++ {
		if r.Hands[i].Player == player {
			return &r.Hands[i]
		}
	}

	// Not expected
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

func (r *Round) dealHands() error {
	if len(r.Hands) > 0 {
		return errors.New("Hands have been dealt already")
	}

	players := r.relation.allPlayers()
	deck := createShuffledDeckOfCards()
	r.Hands = make([]Hand, playersCount)
	for i := range players {
		start := i * cardsInHandCount
		end := start + cardsInHandCount
		r.Hands[i] = Hand{Player: players[i], Cards: deck[start:end]}
	}

	return nil
}

func (r *Round) findPlayerWithNineOfDiamonds() Player {
	for i := 0; i < len(r.Hands); i++ {
		hand := r.Hands[i]

		for j := 0; j < len(hand.Cards); j++ {
			card := hand.Cards[j]

			if card.Rank == NineRank && card.Suit == DiamondsSuit {
				return hand.Player
			}
		}
	}

	// Not expected
	return Player("")
}

// TODO: Check for errors and return them if needed.
func (r *Round) takeMove(player Player, card Card) {
	hand := r.getHand(player)
	hand.takeMove(card)
	r.CurrentTrick().addMove(player, card)
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

func (r *Round) winnerTeam() Team {
	tricksCount := map[Team]int{}

	if !r.IsCompleted() {
		return Team("")
	}

	for _, trick := range r.Tricks {
		tricksCount[r.relation.getTeam(trick.Winner())] += 1
	}

	starterTeam := r.starterTeam()
	opponentTeam := r.starterOpponentTeam()

	// Draw is impossible.
	if tricksCount[starterTeam] > tricksCount[opponentTeam] {
		return starterTeam
	} else {
		return opponentTeam
	}
}

func (r *Round) starterTeam() Team {
	return r.relation.getTeam(r.starter)
}

func (r *Round) starterOpponentTeam() Team {
	return r.relation.getOpponentTeam(r.starter)
}

func (r *Round) starterLeftOpponent() Player {
	return r.relation.getLeftOpponent(r.starter)
}

func createShuffledDeckOfCards() []Card {
	shuffledIndexes := createSliceWithShuffledIndexes()
	deck := make([]Card, cardsCount)
	k := 0

	for i := range ranksCount {
		for j := range suitsCount {
			randInx := shuffledIndexes[k]
			card, _ := newCard(ValidRanks[i], ValidSuits[j]) // safe to ignore error
			deck[randInx] = card
			k++
		}
	}

	return deck
}

// Fisher–Yates shuffle
func createSliceWithShuffledIndexes() []int {
	// [0, 1, 2, 3, ..., 35]
	indexes := make([]int, cardsCount)
	for i := range indexes {
		indexes[i] = i
	}

	// [4, 9, 11, 23, ..., 19]
	for i := len(indexes) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		indexes[i], indexes[j] = indexes[j], indexes[i]
	}

	return indexes
}
