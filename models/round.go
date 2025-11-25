package models

import (
	"errors"
	"fmt"
	"math/rand"
)

type Round struct {
	Number  int
	Hands   []Hand
	Tricks  []Trick
	Trump   string
	starter Player
}

// TODO: Add error if more than max number of rounds started
// I guess there's no need to return a pointer?
func newRound(players []Player, curRound *Round) (*Round, error) {
	newRound := Round{}
	err := newRound.dealHands(players)
	if err != nil {
		return nil, err
	}

	if curRound == nil {
		newRound.starter = *newRound.findPlayerWithNineOfDiamonds()
		newRound.Number = 1
	} else {
		if curRound.winnerTeam().Name == curRound.starter.Name {
			newRound.starter = curRound.starter
		} else {
			newRound.starter = *curRound.starter.leftOpponent
		}
		newRound.Number = curRound.Number + 1
	}

	return &newRound, nil
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

func (r *Round) winnerTeam() *Team {
	tricksCount := map[string]int{}

	if !r.IsCompleted() {
		return nil
	}

	for _, trick := range r.Tricks {
		tricksCount[trick.Winner().Team.Name] += 1
	}

	starterTeam := r.starter.Team
	opponentTeam := r.starter.leftOpponent.Team

	// Draw is impossible.
	if tricksCount[starterTeam.Name] > tricksCount[opponentTeam.Name] {
		return starterTeam
	} else {
		return opponentTeam
	}
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
