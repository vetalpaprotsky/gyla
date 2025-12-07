package models

import "math/rand"

type Deck struct {
	cards []Card
}

func newDeck() Deck {
	shuffledIndexes := createSliceWithShuffledIndexes()
	cards := make([]Card, cardsCount)
	k := 0

	for i := range ranksCount {
		for j := range suitsCount {
			randInx := shuffledIndexes[k]
			card, _ := newCard(ValidRanks[i], ValidSuits[j]) // safe to ignore error
			cards[randInx] = card
			k++
		}
	}

	return Deck{cards: cards}
}

// TODO: If one hand has four 7, or four 6, and we need to re-deal the cards.
// It's not allowed by the game rules.
func (d Deck) deal(pr PlayersRelation) []Hand {
	players := pr.allPlayers()
	hands := make([]Hand, playersCount)

	for i := range players {
		start := i * cardsInHandCount
		end := start + cardsInHandCount
		hands[i] = Hand{Player: players[i], Cards: d.cards[start:end]}
	}

	return hands
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
