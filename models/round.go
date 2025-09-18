package models

import (
	"errors"
	"fmt"
	"math/rand"
)

type Round struct {
	hands         []Hand
	tricks        []Trick
	trump         string
	staringPlayer string
}

func NewRound(players []string) (*Round, error) {
	round := Round{}
	err := round.dealHands(players)
	if err != nil {
		return nil, err
	}

	return &round, nil
}

func (r *Round) AssignTrump(suit string) error {
	if r.trump != "" {
		errorMsg := fmt.Sprintf(
			"Can't assign %s trump, it's already assigned to %s",
			suit,
			r.trump,
		)
		return errors.New(errorMsg)
	}

	r.trump = suit

	for i := 0; i < len(r.hands); i++ {
		r.hands[i].assignTrump(suit)
	}

	return nil
}

func (r *Round) dealHands(players []string) error {
	if len(players) != handsCount {
		errorMsg := fmt.Sprintf(
			"Number of players must be %d, not %d",
			handsCount,
			len(players),
		)
		return errors.New(errorMsg)
	}

	if len(r.hands) > 0 {
		return errors.New("Hands have been dealt already")
	}

	deck := createShuffledDeckOfCards()
	r.hands = make([]Hand, handsCount)
	for i := range handsCount {
		start := i * cardsInHandCount
		end := start + cardsInHandCount
		r.hands[i] = Hand{player: players[i], cards: deck[start:end]}
	}

	return nil
}

func createShuffledDeckOfCards() []Card {
	shuffledIndexes := createSliceWithShuffledIndexes()
	deck := make([]Card, cardsCount)
	k := 0

	for i := range ranksCount {
		for j := range suitsCount {
			randInx := shuffledIndexes[k]
			card, _ := newCard(validRanks[i], validSuits[j]) // safe to ignore error
			deck[randInx] = *card
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
