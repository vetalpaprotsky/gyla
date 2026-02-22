package game

import "math/rand"

type deck struct {
	cards []Card
}

func newDeck() deck {
	shuffledIndexes := createSliceWithShuffledIndexes()
	cards := make([]Card, cardsCount)
	k := 0

	for _, r := range validRanks {
		for _, s := range validSuits {
			randInx := shuffledIndexes[k]
			card, _ := newCard(r, s) // safe to ignore error
			cards[randInx] = card
			k++
		}
	}

	return deck{cards: cards}
}

// TODO: If one hand has four 7, or four 6, and we need to re-deal the cards.
// It's not allowed by the game rules.
func (d deck) deal(t table) []hand {
	players := t.getAllPlayers()
	hands := make([]hand, playersCount)

	for i := range players {
		start := i * cardsInHandCount
		end := start + cardsInHandCount
		hands[i] = hand{player: players[i], cards: d.cards[start:end]}
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
