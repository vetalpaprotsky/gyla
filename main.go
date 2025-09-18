package main

import (
	"fmt"
	"gyla/models"
	"log"
)

func main() {
	players := []string{"homer", "ned", "bart", "lisa"}
	roundPointer, err := models.NewRound(players)

	if err != nil {
		log.Printf("Error: %s", err)
	}

	round := *roundPointer
	round.AssignTrump(models.HeartsSuit)
	fmt.Println(round)
}
