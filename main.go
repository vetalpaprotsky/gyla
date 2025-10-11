package main

import (
	"fmt"
	"gyla/models"
	"log"
)

func main() {
	game := models.NewGame("Homer", "Ned", "Bart", "Lisa")

	round, err := game.StartRound()

	if err != nil {
		log.Println("Error: ", err)
	}

	starter := round.Starter
	fmt.Println(starter.Name)

	// Infinite rounds loop (in theory, we can cound max number of rounds)
	// // 9 tricks loop
	// // // 4 turns loop

	// How to restore a game at any point if let's same something goes wrong?
	// We can save game after each turn, but we need to create some "restore"
	// algorithm, which checks store unfinished state of game and knows how to
	// continue it.
}
