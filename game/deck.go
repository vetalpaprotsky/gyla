package game

import "math/rand"

const cardsCount = len(allRanks) * len(allSuits)
const cardsPerPlayerCount = cardsCount / len(allPlayers)

type deck struct {
	cards []card
}

func newDeck() deck {
	shuffledIndexes := createSliceWithShuffledIndexes()
	cards := make([]card, cardsCount)
	k := 0

	for _, r := range allRanks {
		for _, s := range allSuits {
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
func (d deck) deal() []hand {
	hands := make([]hand, len(allPlayers))

	for i := range allPlayers {
		start := i * cardsPerPlayerCount
		end := start + cardsPerPlayerCount
		hands[i] = hand{player: allPlayers[i], cards: d.cards[start:end]}
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
