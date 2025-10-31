package models

import (
	"errors"
	"fmt"
	"math/rand"
)

type Round struct {
	number  int
	hands   []Hand
	tricks  []Trick
	trump   string
	Starter Player
}

// TODO: Add error if more than max number of rounds started
func newRound(players []Player, curRound *Round) (*Round, error) {
	newRound := Round{}
	err := newRound.dealHands(players)
	if err != nil {
		return nil, err
	}

	if curRound == nil {
		newRound.Starter = *newRound.findPlayerWithNineOfDiamonds()
		newRound.number = 1
	} else {
		if curRound.winnerTeam().Name == curRound.Starter.Name {
			newRound.Starter = curRound.Starter
		} else {
			newRound.Starter = *curRound.Starter.LeftOpponent
		}
		newRound.number = curRound.number + 1
	}

	return &newRound, nil
}

func (r *Round) currentTrick() *Trick {
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

func (r *Round) startNextTrick() *Trick {
	curTrick := r.currentTrick()
	var trick *Trick

	if curTrick == nil {
		trick = newFirstTrick(r.Starter)
	} else {
		trick = newTrick(curTrick)
	}

	r.tricks = append(r.tricks, *trick)

	return trick
}

func (r *Round) starterHand() *Hand {
	for i := 0; i < len(r.hands); i++ {
		if r.hands[i].player.Name == r.Starter.Name {
			return &r.hands[i]
		}
	}

	// Not expected
	return nil
}

func (r *Round) findHand(player Player) *Hand {
	for i := 0; i < len(r.hands); i++ {
		if r.hands[i].player.Name == player.Name {
			return &r.hands[i]
		}
	}

	// Not expected
	return nil
}

func (r *Round) nextTrickStarterHand() *Hand {
	if len(r.tricks) == 0 {
		return r.starterHand()
	} else {
		return r.findHand(r.currentTrick().winner())
	}
}

func (r *Round) assignTrump(suit string) error {
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

func (r *Round) dealHands(players []Player) error {
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

func (r *Round) findPlayerWithNineOfDiamonds() *Player {
	for i := 0; i < len(r.hands); i++ {
		hand := r.hands[i]

		for j := 0; j < len(hand.cards); j++ {
			card := hand.cards[j]

			if card.rank == NineRank && card.suit == DiamondsSuit {
				return &hand.player
			}
		}
	}

	// Not expected
	return nil
}

func (r *Round) winnerTeam() Team {
	// TODO
	return Team{}
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
