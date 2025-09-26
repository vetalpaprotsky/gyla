package main

import (
	"fmt"
	"gyla/models"
	"log"
)

func main() {
	player1 := models.Player{Name: "homer"}
	player2 := models.Player{Name: "ned"}
	player3 := models.Player{Name: "bart"}
	player4 := models.Player{Name: "lisa"}

	player1.LeftOpponent = &player2
	player1.Teammate = &player3
	player1.RightOpponent = &player4

	player2.LeftOpponent = &player3
	player2.Teammate = &player4
	player2.RightOpponent = &player1

	player3.LeftOpponent = &player4
	player3.Teammate = &player1
	player3.RightOpponent = &player2

	player4.LeftOpponent = &player1
	player4.Teammate = &player2
	player4.RightOpponent = &player3

	team1 := models.Team{Name: "team1", Player1: &player1, Player2: &player3}
	team2 := models.Team{Name: "team2", Player1: &player2, Player2: &player4}

	player1.Team = &team1
	player2.Team = &team2
	player3.Team = &team1
	player4.Team = &team2

	players := []models.Player{player1, player2, player3, player4}

	roundPointer, err := models.NewRound(players, nil)

	if err != nil {
		log.Printf("Error: %s", err)
	}

	round := *roundPointer
	round.AssignTrump(models.HeartsSuit)
	fmt.Println(round)
}
