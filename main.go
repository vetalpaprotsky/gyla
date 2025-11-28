package main

import (
	"fmt"
	"gyla/models"
	"strings"
)

// TODO: Once GameState model is created, we could make almost all fields
// in all structs start with a lower letter (private fields).
//
// GameState - stores game state without diving into details. Very useful for
// rendering a view.
func stateChangeCallback(g *models.Game) {
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Teams:")
	fmt.Printf("\t%s: %s, %s\n", g.Relation.Team1, g.Relation.Player1, g.Relation.Player3)
	fmt.Printf("\t%s: %s, %s\n", g.Relation.Team2, g.Relation.Player2, g.Relation.Player4)

	round := g.CurrentRound()
	if round == nil {
		return
	}

	score := g.Score()
	fmt.Printf(
		"Score:\n\t%s: %d\n\t%s: %d\n",
		score.Team1,
		score.Points1,
		score.Team2,
		score.Points2,
	)

	fmt.Println("Round number:", round.Number)
	fmt.Println("Tricks:")
	for _, trick := range round.Tricks {
		if !trick.IsCompleted() {
			continue
		}

		fmt.Printf("\tNumber: %d\n", trick.Number)
		for _, move := range trick.Moves {
			fmt.Printf("\t%s: %s", move.Player, move.Card.TuiID())
			if trick.Winner() == move.Player {
				fmt.Println(" -> Winner")
			} else {
				fmt.Println()
			}
		}
	}

	fmt.Println("Hands:")
	for _, hand := range round.Hands {
		var cards []string

		for _, card := range hand.Cards {
			cards = append(cards, card.TuiID())
		}

		fmt.Printf("\t%s: %s\n", hand.Player, strings.Join(cards, " "))
	}

	trick := round.CurrentTrick()
	if trick == nil {
		return
	}

	fmt.Println("Trick number: ", round.CurrentTrick().Number)
	fmt.Println("Round trump:", round.Trump)
	fmt.Println("Moves:")
	for _, move := range round.CurrentTrick().Moves {
		fmt.Printf("\t%s: %s\n", move.Player, move.Card.TuiID())
	}
}

func playerTrumpAssignmentCallback(p models.Player, cards []models.Card) string {
	for {
		var suit string

		fmt.Printf("Enter trump suit <%s>:", p)
		fmt.Scan(&suit)
		suit = strings.ToUpper(suit)

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
		var cardID string
		var cardIDs []string

		for _, card := range cards {
			cardIDs = append(cardIDs, card.TuiID())
		}

		fmt.Printf("Enter rank and suit (%s) <%s>:", strings.Join(cardIDs, " "), p)
		fmt.Scan(&cardID)
		rank := strings.ToUpper(cardID[:len(cardID)-1])
		suit := strings.ToUpper(cardID[len(cardID)-1:])

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
