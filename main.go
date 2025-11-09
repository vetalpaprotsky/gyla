package main

import (
	"fmt"
	"gyla/models"
	"strings"
)

func stateChangeCallback(g *models.Game) {
	fmt.Printf("Player 1: %s", g.Player1.Name)
	fmt.Printf("Player 2: %s", g.Player2.Name)
	fmt.Printf("Player 3: %s", g.Player3.Name)
	fmt.Printf("Player 4: %s", g.Player4.Name)

	round := g.CurrentRound()

	for i, hand := range round.Hands {
		var cards []string

		for _, card := range hand.Cards {
			cards = append(cards, card.ID())
		}

		fmt.Printf("Player %d, %s - %s", i+1, hand.Player.Name, strings.Join(cards, ", "))
	}

	// TODO Current trick
}

// TODO: if suit is invalid, ask for one more attempt.
func playerTrumpAssignmentCallback(player string, cards []models.Card) string {
	for {
		var suit string

		fmt.Printf("%s - Enter trump suit: ", player)
		fmt.Scan(&suit)

		for _, validSuit := range models.ValidSuits {
			if suit == validSuit {
				return suit
			}
		}

		fmt.Printf("%s - Invalid suit entered", player)
	}
}

func playerMoveCallback(player string, cards []models.Card) models.Card {
	for {
		var rank, suit string

		fmt.Printf("%s - Enter rank and suit: ", player)
		fmt.Scan(&rank, &suit)

		for _, card := range cards {
			if rank == card.Rank && suit == card.Suit {
				return card
			}
		}

		fmt.Printf("%s - Invalid rank and suit entered", player)
	}
}

func main() {
	game := models.NewGame("Homer", "Ned", "Bart", "Lisa")

	// TODO: Check for errors
	game.StartGameLoop(
		stateChangeCallback,
		playerTrumpAssignmentCallback,
		playerMoveCallback,
	)

	// How to restore a game at any point if let's same something goes wrong?
	// We can save game after each move, but we need to create some "restore"
	// algorithm, which checks store unfinished state of game and knows how to
	// continue it.
}
