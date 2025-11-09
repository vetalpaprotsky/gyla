package models

import (
	"errors"
	"fmt"
	"math/rand"
)

type Round struct {
	number  int
	Hands   []Hand
	tricks  []Trick
	Trump   string
	starter Player
}

// TODO: Add error if more than max number of rounds started
func newRound(players []Player, curRound *Round) (*Round, error) {
	newRound := Round{}
	err := newRound.dealHands(players)
	if err != nil {
		return nil, err
	}

	if curRound == nil {
		newRound.starter = *newRound.findPlayerWithNineOfDiamonds()
		newRound.number = 1
	} else {
		if curRound.winnerTeam().Name == curRound.starter.Name {
			newRound.starter = curRound.starter
		} else {
			newRound.starter = *curRound.starter.leftOpponent
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
		trick = newFirstTrick(r.starter)
	} else {
		trick = newTrick(curTrick)
	}

	r.tricks = append(r.tricks, *trick)

	return trick
}

func (r *Round) getHand(player Player) *Hand {
	for i := 0; i < len(r.Hands); i++ {
		if r.Hands[i].Player.Name == player.Name {
			return &r.Hands[i]
		}
	}

	// Not expected
	return nil
}

func (r *Round) assignTrump(suit string) error {
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

func (r *Round) dealHands(players []Player) error {
	if len(players) != handsCount {
		errorMsg := fmt.Sprintf(
			"Number of players must be %d, not %d",
			handsCount,
			len(players),
		)
		return errors.New(errorMsg)
	}

	if len(r.Hands) > 0 {
		return errors.New("Hands have been dealt already")
	}

	deck := createShuffledDeckOfCards()
	r.Hands = make([]Hand, handsCount)
	for i := range handsCount {
		start := i * cardsInHandCount
		end := start + cardsInHandCount
		r.Hands[i] = Hand{Player: players[i], Cards: deck[start:end]}
	}

	return nil
}

func (r *Round) findPlayerWithNineOfDiamonds() *Player {
	for i := 0; i < len(r.Hands); i++ {
		hand := r.Hands[i]

		for j := 0; j < len(hand.Cards); j++ {
			card := hand.Cards[j]

			if card.Rank == NineRank && card.Suit == DiamondsSuit {
				return &hand.Player
			}
		}
	}

	// Not expected
	return nil
}

// TODO: Check for errors and return them if needed.
func (r *Round) takeMove(player Player, card Card) {
	hand := r.getHand(player)
	hand.takeMove(card)
	r.currentTrick().addMove(player, card)
}

func (r *Round) availableCardsForMove(player Player) []Card {
	trick := *r.currentTrick()
	hand := r.getHand(player)

	return hand.availableCardsForMove(trick)
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
			card, _ := newCard(ValidRanks[i], ValidSuits[j]) // safe to ignore error
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
