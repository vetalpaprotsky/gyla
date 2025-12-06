package main

import (
	"fmt"
	"gyla/models"
	"gyla/tui"
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
		winner, winnerOk := trick.Winner()
		if !winnerOk {
			continue
		}

		fmt.Printf("\tNumber: %d\n", trick.Number)
		for _, move := range trick.Moves {
			fmt.Printf("\t%s: %s", move.Player, tui.Card(move.Card))
			if winner == move.Player {
				fmt.Println(" -> Winner")
			} else {
				fmt.Println()
			}
		}
	}

	tricksPerTeam := round.TricksPerTeam()
	fmt.Printf("%s tricks: %d\n", tricksPerTeam.Team1, tricksPerTeam.Tricks1)
	fmt.Printf("%s tricks: %d\n", tricksPerTeam.Team2, tricksPerTeam.Tricks2)

	fmt.Println("Hands:")
	for _, hand := range round.Hands {
		fmt.Printf("\t%s: %s\n", hand.Player, tui.Cards(hand.Cards))
	}

	trick := round.CurrentTrick()
	if trick == nil {
		return
	}

	fmt.Println("Trick number:", round.CurrentTrick().Number)
	fmt.Println("Round trump:", tui.Suit(round.Trump))
	fmt.Println("Moves:")
	for _, move := range round.CurrentTrick().Moves {
		fmt.Printf("\t%s: %s\n", move.Player, tui.Card(move.Card))
	}
}

func playerTrumpAssignmentCallback(p models.Player, cards []models.Card) models.Suit {
	for {
		var suitStr string

		fmt.Printf("Enter trump suit <%s>:", p)
		fmt.Scan(&suitStr)
		suit := models.Suit(strings.ToUpper(suitStr))

		if suit.IsValid() {
			return suit
		}

		fmt.Printf("Invalid suit entered <%s>\n", p)
	}
}

func playerMoveCallback(p models.Player, cards []models.Card) models.Card {
	for {
		var rankAndSuit string

		fmt.Printf("Enter rank and suit (%s) <%s>:", tui.Cards(cards), p)
		fmt.Scan(&rankAndSuit)
		card, err := models.NewCardFromRankAndSuit(rankAndSuit)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		for _, c := range cards {
			if c.Rank == card.Rank && c.Suit == card.Suit {
				return c
			}
		}

		fmt.Printf("Invalid rank and suit entered <%s>\n", p)
	}
}

func main() {
	game := models.NewGame("Old", "Homer", "Ned", "Young", "Bart", "Lisa")
	err := game.StartGameLoop(
		stateChangeCallback,
		playerTrumpAssignmentCallback,
		playerMoveCallback,
	)

	if err != nil {
		fmt.Println("Error: ", err)
	}
}
