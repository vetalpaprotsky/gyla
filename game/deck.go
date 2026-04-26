package game

import "math/rand"

const cardsCount = len(allRanks) * len(allSuits)
const cardsPerPlayerCount = cardsCount / len(allPlayers)

func dealHands() []hand {
outer:
	for {
		deck := shuffledDeck()
		hands := make([]hand, len(allPlayers))

		for i := range allPlayers {
			start := i * cardsPerPlayerCount
			end := start + cardsPerPlayerCount
			cards := deck[start:end]

			if includeAllSuitsOf(cards, SixRank) || includeAllSuitsOf(cards, SevenRank) {
				continue outer
			}

			hands[i] = hand{player: allPlayers[i], cards: cards}
		}

		return hands
	}
}

func includeAllSuitsOf(cards []card, rank Rank) bool {
	count := 0

	for _, c := range cards {
		if c.rank == rank {
			count += 1
		}
	}

	return count == len(allSuits)
}

func shuffledDeck() []card {
	shuffledIndexes := createSliceWithShuffledIndexes()
	cards := make([]card, cardsCount)
	k := 0

	for _, r := range allRanks {
		for _, s := range allSuits {
			randInx := shuffledIndexes[k]
			card, err := newCard(r, s)
			if err != nil {
				panic(err)
			}
			cards[randInx] = card
			k++
		}
	}

	return cards
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
