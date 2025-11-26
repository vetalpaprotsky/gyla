package main

import (
	"fmt"
	"gyla/models"
	"strings"
)

// TODO: once GameState model is created, we could make almost all fields
// in all structs start with a lower letter (private fields).
//
// GameState - stores game state without diving into details. Very useful for
// rendering a view.
func stateChangeCallback(g *models.Game) {
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Player 1:", g.Relation.Player1)
	fmt.Println("Player 2:", g.Relation.Player2)
	fmt.Println("Player 3:", g.Relation.Player3)
	fmt.Println("Player 4:", g.Relation.Player4)

	round := g.CurrentRound()
	if round == nil {
		return
	}

	fmt.Println("Round number: ", round.Number)
	fmt.Println("Round trump: ", round.Trump)
	fmt.Println("Tricks: ")
	for _, trick := range round.Tricks {
		if !trick.IsCompleted() {
			continue
		}

		fmt.Printf("\nNumber: %d\n", trick.Number)
		for _, move := range trick.Moves {
			fmt.Printf("\t %s: %s", move.Player, move.Card.ID())
			if trick.Winner() == move.Player {
				fmt.Println(" -> Winner")
			} else {
				fmt.Println()
			}
		}
	}

	fmt.Println("Hands: ")
	for _, hand := range round.Hands {
		var cards []string

		for _, card := range hand.Cards {
			cards = append(cards, card.ID())
		}

		fmt.Printf("\t %s -> %s\n", hand.Player, strings.Join(cards, ", "))
	}

	trick := round.CurrentTrick()
	if trick == nil {
		return
	}

	fmt.Println("Trick number: ", round.CurrentTrick().Number)
	fmt.Println("Moves: ")

	for _, move := range round.CurrentTrick().Moves {
		fmt.Printf("\t %s -> %s\n", move.Player, move.Card.ID())
	}
}

func playerTrumpAssignmentCallback(p models.Player, cards []models.Card) string {
	for {
		var suit string

		fmt.Printf("Enter trump suit <%s>:", p)
		fmt.Scan(&suit)

		for _, validSuit := range models.ValidSuits {
			if suit == validSuit {
				return suit
			}
		}

		fmt.Printf("Invalid suit entered <%s>", p)
	}
}

func playerMoveCallback(p models.Player, cards []models.Card) models.Card {
	for {
		var rank, suit string
		var cardIDs []string

		for _, card := range cards {
			cardIDs = append(cardIDs, card.ID())
		}

		fmt.Printf("Enter rank and suit (%s) <%s>:", strings.Join(cardIDs, ", "), p)
		fmt.Scan(&rank, &suit)

		for _, card := range cards {
			if rank == card.Rank && suit == card.Suit {
				return card
			}
		}

		fmt.Printf("Invalid rank and suit entered <%s>", p)
	}
}

func main() {
	game := models.NewGame("Old", "Homer", "Ned", "Young", "Bart", "Lisa")

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
