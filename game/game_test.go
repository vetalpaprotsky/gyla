package game

import (
	"testing"
)

func TestGameWorkflow(t *testing.T) {
	t.Run("cannot apply actions before starting", func(t *testing.T) {
		g := NewGame("Alice", "Bob", "Charlie", "Diana", "Team A", "Team B", false, false, false, false)

		// Apply should fail before Start is called — game has no round yet
		_, err := g.Apply(Action{Name: AssignTrumpAction, Player: Player1, Suit: ClubsSuit})
		if err == nil {
			t.Fatal("expected error when applying action before game starts")
		}
	})

	t.Run("cannot start game twice", func(t *testing.T) {
		g := NewGame("Alice", "Bob", "Charlie", "Diana", "Team A", "Team B", false, false, false, false)

		_, err := g.Start()
		if err != nil {
			t.Fatalf("expected no error on first start, got: %v", err)
		}

		_, err = g.Start()
		if err == nil {
			t.Fatal("expected error when starting game twice")
		}
	})

	t.Run("start emits correct initial events", func(t *testing.T) {
		g := NewGame("Alice", "Bob", "Charlie", "Diana", "Team A", "Team B", false, false, false, false)

		events, err := g.Start()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(events) < 2 {
			t.Fatalf("expected at least 2 events (GameStarted, RoundStarted), got %d", len(events))
		}

		if events[0].EventType != GameStartedEvent {
			t.Errorf("expected first event to be %s, got %s", GameStartedEvent, events[0].EventType)
		}

		if events[1].EventType != RoundStartedEvent {
			t.Errorf("expected second event to be %s, got %s", RoundStartedEvent, events[1].EventType)
		}

		// Verify game state has player info
		state := events[len(events)-1].GameState
		if len(state.PlayersInfo) != 4 {
			t.Errorf("expected 4 players, got %d", len(state.PlayersInfo))
		}

		// Round 1 should be started
		if state.Round.Number != 1 {
			t.Errorf("expected round 1, got %d", state.Round.Number)
		}

		// Each player should have 9 cards
		for _, h := range state.Round.Hands {
			if len(h.Cards) != cardsPerPlayerCount {
				t.Errorf("player %v should have %d cards, got %d", h.Player, cardsPerPlayerCount, len(h.Cards))
			}
		}
	})

	t.Run("trump assignment validation", func(t *testing.T) {
		g := NewGame("Alice", "Bob", "Charlie", "Diana", "Team A", "Team B", false, false, false, false)

		events, _ := g.Start()
		state := events[len(events)-1].GameState
		starter := state.Round.Trumper

		// Wrong player tries to assign trump
		wrongPlayer := starter.leftOpponent()
		_, err := g.Apply(Action{Name: AssignTrumpAction, Player: wrongPlayer, Suit: ClubsSuit})
		if err == nil {
			t.Fatal("expected error when wrong player assigns trump")
		}

		// Invalid suit
		_, err = g.Apply(Action{Name: AssignTrumpAction, Player: starter, Suit: Suit(99)})
		if err == nil {
			t.Fatal("expected error for invalid trump suit")
		}

		// Correct player assigns trump
		events, err = g.Apply(Action{Name: AssignTrumpAction, Player: starter, Suit: HeartsSuit})
		if err != nil {
			t.Fatalf("unexpected error assigning trump: %v", err)
		}

		foundTrumpAssigned := false
		for _, e := range events {
			if e.EventType == TrumpAssignedEvent {
				foundTrumpAssigned = true
				if e.GameState.Round.Trump != HeartsSuit {
					t.Errorf("expected trump to be Hearts, got %v", e.GameState.Round.Trump)
				}
			}
		}
		if !foundTrumpAssigned {
			t.Error("expected TrumpAssignedEvent in events")
		}

		// Cannot assign trump again
		_, err = g.Apply(Action{Name: AssignTrumpAction, Player: starter, Suit: SpadesSuit})
		if err == nil {
			t.Fatal("expected error when assigning trump twice")
		}
	})

	t.Run("card play validation", func(t *testing.T) {
		g := NewGame("Alice", "Bob", "Charlie", "Diana", "Team A", "Team B", false, false, false, false)

		events, _ := g.Start()
		state := events[len(events)-1].GameState
		starter := state.Round.Trumper

		// Cannot play card before trump is assigned
		hand := state.Round.getHand(starter)
		if len(hand.Cards) > 0 {
			c := hand.Cards[0]
			_, err := g.Apply(Action{Name: PlayCardAction, Player: starter, Rank: c.Card.Rank, Suit: c.Card.Suit})
			if err == nil {
				t.Fatal("expected error when playing card before trump is assigned")
			}
		}

		// Assign trump
		g.Apply(Action{Name: AssignTrumpAction, Player: starter, Suit: DiamondsSuit})

		// Get updated state
		curState := g.State()
		curTrick := curState.Round.Tricks
		if len(curTrick) == 0 {
			t.Fatal("expected at least one trick after trump assignment")
		}

		nextPlayer := curTrick[len(curTrick)-1].Next

		// Wrong player tries to play
		wrongPlayer := nextPlayer.leftOpponent()
		wrongHand := curState.Round.getHand(wrongPlayer)
		if len(wrongHand.Cards) > 0 {
			c := wrongHand.Cards[0]
			_, err := g.Apply(Action{Name: PlayCardAction, Player: wrongPlayer, Rank: c.Card.Rank, Suit: c.Card.Suit})
			if err == nil {
				t.Fatal("expected error when wrong player plays a card")
			}
		}

		// Correct player plays a playable card
		correctHand := curState.Round.getHand(nextPlayer)
		var playableCard HandCard
		for _, c := range correctHand.Cards {
			if c.IsPlayable {
				playableCard = c
				break
			}
		}

		events, err := g.Apply(Action{Name: PlayCardAction, Player: nextPlayer, Rank: playableCard.Card.Rank, Suit: playableCard.Card.Suit})
		if err != nil {
			t.Fatalf("unexpected error playing card: %v", err)
		}

		foundCardPlayed := false
		for _, e := range events {
			if e.EventType == CardPlayedEvent ||
				e.EventType == CardPlayedAndTrickCompletedEvent ||
				e.EventType == CardPlayedAndRoundCompletedEvent ||
				e.EventType == CardPlayedAndGameCompletedEvent {
				foundCardPlayed = true
			}
		}
		if !foundCardPlayed {
			t.Error("expected a card played event")
		}
	})

	t.Run("play full round", func(t *testing.T) {
		g := NewGame("Alice", "Bob", "Charlie", "Diana", "Team A", "Team B", false, false, false, false)

		events, _ := g.Start()
		state := events[len(events)-1].GameState
		starter := state.Round.Trumper

		// Assign trump
		events, err := g.Apply(Action{Name: AssignTrumpAction, Player: starter, Suit: ClubsSuit})
		if err != nil {
			t.Fatalf("unexpected error assigning trump: %v", err)
		}

		// Play 9 tricks (4 cards each = 36 cards total)
		cardsPlayed := 0
		for cardsPlayed < 36 {
			curState := g.State()
			tricks := curState.Round.Tricks
			if len(tricks) == 0 {
				t.Fatal("expected at least one trick")
			}

			curTrick := tricks[len(tricks)-1]
			nextPlayer := curTrick.Next

			if nextPlayer.isZero() {
				t.Fatal("no next player but round not complete")
			}

			hand := curState.Round.getHand(nextPlayer)
			var playable HandCard
			found := false
			for _, c := range hand.Cards {
				if c.IsPlayable {
					playable = c
					found = true
					break
				}
			}

			if !found {
				t.Fatalf("no playable cards for player %v at card %d", nextPlayer, cardsPlayed)
			}

			events, err = g.Apply(Action{
				Name:   PlayCardAction,
				Player: nextPlayer,
				Rank:   playable.Card.Rank,
				Suit:   playable.Card.Suit,
			})
			if err != nil {
				t.Fatalf("unexpected error playing card %d: %v", cardsPlayed, err)
			}

			cardsPlayed++

			// Check for round completion
			for _, e := range events {
				if e.EventType == CardPlayedAndRoundCompletedEvent ||
					e.EventType == CardPlayedAndGameCompletedEvent {
					if cardsPlayed != 36 {
						t.Fatalf("round/game completed after only %d cards", cardsPlayed)
					}
				}
			}
		}

		// After 36 cards, the round should be completed
		lastEvent := events[len(events)-1]
		finalState := lastEvent.GameState

		if lastEvent.EventType != CardPlayedAndGameCompletedEvent {
			// Round completed but game continues — new round should have started
			if finalState.Round.Number < 2 {
				// Check the event was at least round completed
				foundRoundComplete := false
				for _, e := range events {
					if e.EventType == CardPlayedAndRoundCompletedEvent {
						foundRoundComplete = true
					}
				}
				if !foundRoundComplete {
					t.Error("expected round completed or game completed event after 36 cards")
				}
			}
		}
	})

	t.Run("play full game until completion", func(t *testing.T) {
		g := NewGame("Alice", "Bob", "Charlie", "Diana", "Team A", "Team B", false, false, false, false)

		events, err := g.Start()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		maxRounds := 20
		roundsPlayed := 0
		gameCompleted := false

		for !gameCompleted && roundsPlayed < maxRounds {
			curState := g.State()

			// Assign trump if needed (check if tricks have started)
			if len(curState.Round.Tricks) == 0 {
				starter := curState.Round.Trumper
				events, err = g.Apply(Action{
					Name:   AssignTrumpAction,
					Player: starter,
					Suit:   allSuits[roundsPlayed%len(allSuits)],
				})
				if err != nil {
					t.Fatalf("error assigning trump in round %d: %v", roundsPlayed+1, err)
				}
			}

			// Play cards until round ends
			for {
				curState = g.State()
				tricks := curState.Round.Tricks
				if len(tricks) == 0 {
					break
				}

				curTrick := tricks[len(tricks)-1]
				nextPlayer := curTrick.Next
				if nextPlayer.isZero() {
					break
				}

				hand := curState.Round.getHand(nextPlayer)
				var playable HandCard
				found := false
				for _, c := range hand.Cards {
					if c.IsPlayable {
						playable = c
						found = true
						break
					}
				}

				if !found {
					t.Fatalf("no playable cards for player %v in round %d", nextPlayer, roundsPlayed+1)
				}

				events, err = g.Apply(Action{
					Name:   PlayCardAction,
					Player: nextPlayer,
					Rank:   playable.Card.Rank,
					Suit:   playable.Card.Suit,
				})
				if err != nil {
					t.Fatalf("error playing card in round %d: %v", roundsPlayed+1, err)
				}

				for _, e := range events {
					if e.EventType == CardPlayedAndGameCompletedEvent {
						gameCompleted = true
						finalState := e.GameState
						if finalState.Stats.WinTeam.isZero() {
							t.Error("game completed but no winning team")
						}
						if finalState.Stats.Points[finalState.Stats.WinTeam] < 60 {
							t.Errorf("winning team has %d points, expected >= 60",
								finalState.Stats.Points[finalState.Stats.WinTeam])
						}
						t.Logf("Game completed! %v wins with %d points after %d rounds",
							finalState.Stats.WinTeam,
							finalState.Stats.Points[finalState.Stats.WinTeam],
							roundsPlayed+1)
					}
					if e.EventType == CardPlayedAndRoundCompletedEvent {
						roundsPlayed++
					}
				}

				if gameCompleted {
					break
				}
			}

			if gameCompleted {
				break
			}
		}

		if !gameCompleted {
			t.Fatalf("game did not complete within %d rounds", maxRounds)
		}

		// Verify final state
		finalState := g.State()
		if finalState.Stats.WinTeam.isZero() {
			t.Error("expected a winning team")
		}

		// Verify stats have round entries
		if len(finalState.Stats.Rounds) == 0 {
			t.Error("expected at least one round in stats")
		}

		for i, rs := range finalState.Stats.Rounds {
			if rs.WinTeam.isZero() {
				t.Errorf("round %d stats has no winning team", i+1)
			}
			if rs.Tricks[Team1]+rs.Tricks[Team2] != tricksPerRoundCount {
				t.Errorf("round %d: total tricks = %d, expected %d",
					i+1, rs.Tricks[Team1]+rs.Tricks[Team2], tricksPerRoundCount)
			}
		}
	})

	t.Run("event states are consistent snapshots", func(t *testing.T) {
		g := NewGame("Alice", "Bob", "Charlie", "Diana", "Team A", "Team B", false, false, false, false)

		events, _ := g.Start()

		// Each event should carry a valid GameState
		for i, e := range events {
			if len(e.GameState.PlayersInfo) != 4 {
				t.Errorf("event %d (%s): expected 4 players, got %d",
					i, e.EventType, len(e.GameState.PlayersInfo))
			}
			// GameStartedEvent is emitted before the first round starts, so round number is 0
			if e.EventType == GameStartedEvent {
				if e.GameState.Round.Number != 0 {
					t.Errorf("event %d (%s): round number should be 0, got %d",
						i, e.EventType, e.GameState.Round.Number)
				}
			} else {
				if e.GameState.Round.Number < 1 {
					t.Errorf("event %d (%s): round number should be >= 1, got %d",
						i, e.EventType, e.GameState.Round.Number)
				}
			}
		}
	})

	t.Run("GameView hides opponent hands", func(t *testing.T) {
		g := NewGame("Alice", "Bob", "Charlie", "Diana", "Team A", "Team B", false, false, false, false)

		events, _ := g.Start()
		state := events[len(events)-1].GameState

		view := state.ViewFor(Player1)

		if view.You != Player1 {
			t.Errorf("expected You to be Player1, got %v", view.You)
		}
		if view.LeftOpponent != Player2 {
			t.Errorf("expected LeftOpponent to be Player2, got %v", view.LeftOpponent)
		}
		if view.Teammate != Player3 {
			t.Errorf("expected Teammate to be Player3, got %v", view.Teammate)
		}
		if view.RightOpponent != Player4 {
			t.Errorf("expected RightOpponent to be Player4, got %v", view.RightOpponent)
		}

		// View should show own hand cards but only counts for others
		if len(view.Round.Hand.Cards) != cardsPerPlayerCount {
			t.Errorf("expected %d cards in own hand, got %d", cardsPerPlayerCount, len(view.Round.Hand.Cards))
		}
		if view.Round.LeftOpponentHand != cardsPerPlayerCount {
			t.Errorf("expected left opponent to have %d cards, got %d", cardsPerPlayerCount, view.Round.LeftOpponentHand)
		}
		if view.Round.TeammateHand != cardsPerPlayerCount {
			t.Errorf("expected teammate to have %d cards, got %d", cardsPerPlayerCount, view.Round.TeammateHand)
		}
		if view.Round.RightOpponentHand != cardsPerPlayerCount {
			t.Errorf("expected right opponent to have %d cards, got %d", cardsPerPlayerCount, view.Round.RightOpponentHand)
		}
	})

	t.Run("playing non-existent card fails", func(t *testing.T) {
		g := NewGame("Alice", "Bob", "Charlie", "Diana", "Team A", "Team B", false, false, false, false)

		events, _ := g.Start()
		state := events[len(events)-1].GameState
		starter := state.Round.Trumper

		g.Apply(Action{Name: AssignTrumpAction, Player: starter, Suit: SpadesSuit})

		curState := g.State()
		tricks := curState.Round.Tricks
		nextPlayer := tricks[len(tricks)-1].Next

		// Try to play a card with invalid rank
		_, err := g.Apply(Action{
			Name:   PlayCardAction,
			Player: nextPlayer,
			Rank:   Rank(99),
			Suit:   ClubsSuit,
		})
		if err == nil {
			t.Fatal("expected error when playing non-existent card")
		}
	})

	t.Run("follow suit rule is enforced", func(t *testing.T) {
		g := NewGame("Alice", "Bob", "Charlie", "Diana", "Team A", "Team B", false, false, false, false)

		events, _ := g.Start()
		state := events[len(events)-1].GameState
		starter := state.Round.Trumper

		// Assign trump
		g.Apply(Action{Name: AssignTrumpAction, Player: starter, Suit: HeartsSuit})

		curState := g.State()
		tricks := curState.Round.Tricks
		firstPlayer := tricks[len(tricks)-1].Next

		// First player plays any card (first playable)
		hand1 := curState.Round.getHand(firstPlayer)
		var firstCard HandCard
		for _, c := range hand1.Cards {
			if c.IsPlayable {
				firstCard = c
				break
			}
		}

		g.Apply(Action{
			Name:   PlayCardAction,
			Player: firstPlayer,
			Rank:   firstCard.Card.Rank,
			Suit:   firstCard.Card.Suit,
		})

		// Now check second player can only play playable cards
		curState = g.State()
		tricks = curState.Round.Tricks
		secondPlayer := tricks[len(tricks)-1].Next
		hand2 := curState.Round.getHand(secondPlayer)

		for _, c := range hand2.Cards {
			if !c.IsPlayable {
				// Try to play a non-playable card — should fail
				_, err := g.Apply(Action{
					Name:   PlayCardAction,
					Player: secondPlayer,
					Rank:   c.Card.Rank,
					Suit:   c.Card.Suit,
				})
				if err == nil {
					t.Fatalf("expected error when playing non-playable card (rank=%v suit=%v)", c.Card.Rank, c.Card.Suit)
				}
				break
			}
		}
	})

	t.Run("round stats have correct points", func(t *testing.T) {
		g := NewGame("Alice", "Bob", "Charlie", "Diana", "Team A", "Team B", false, false, false, false)

		g.Start()

		// Play one full round
		curState := g.State()
		starter := curState.Round.Trumper
		g.Apply(Action{Name: AssignTrumpAction, Player: starter, Suit: ClubsSuit})

		for {
			curState = g.State()
			tricks := curState.Round.Tricks
			if len(tricks) == 0 {
				break
			}

			curTrick := tricks[len(tricks)-1]
			nextPlayer := curTrick.Next
			if nextPlayer.isZero() {
				break
			}

			hand := curState.Round.getHand(nextPlayer)
			var playable HandCard
			for _, c := range hand.Cards {
				if c.IsPlayable {
					playable = c
					break
				}
			}

			events, err := g.Apply(Action{
				Name:   PlayCardAction,
				Player: nextPlayer,
				Rank:   playable.Card.Rank,
				Suit:   playable.Card.Suit,
			})
			if err != nil {
				t.Fatalf("error playing card: %v", err)
			}

			for _, e := range events {
				if e.EventType == CardPlayedAndRoundCompletedEvent ||
					e.EventType == CardPlayedAndGameCompletedEvent {
					// Check round stats
					stats := e.GameState.Stats
					if len(stats.Rounds) == 0 {
						t.Fatal("expected at least one round in stats after round completion")
					}
					rs := stats.Rounds[0]

					// Winner must have more tricks
					winTeam := rs.WinTeam
					losingTeam := winTeam.opponent()
					if rs.Tricks[winTeam] <= rs.Tricks[losingTeam] {
						t.Errorf("winning team %v has %d tricks, losing team has %d",
							winTeam, rs.Tricks[winTeam], rs.Tricks[losingTeam])
					}

					// Points must be positive for winner
					if rs.Points[winTeam] <= 0 {
						t.Errorf("expected positive points for winning team, got %d", rs.Points[winTeam])
					}

					// Losing team gets 0 points
					if rs.Points[losingTeam] != 0 {
						t.Errorf("expected 0 points for losing team, got %d", rs.Points[losingTeam])
					}

					return
				}
			}
		}
	})

	t.Run("all AI players complete game automatically on start", func(t *testing.T) {
		g := NewGame("Bot1", "Bot2", "Bot3", "Bot4", "Team A", "Team B", true, true, true, true)

		events, err := g.Start()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// With all AI players, the game should complete entirely during Start()
		lastEvent := events[len(events)-1]
		if lastEvent.EventType != CardPlayedAndGameCompletedEvent {
			t.Fatalf("expected game to complete automatically with all AI players, last event was %s", lastEvent.EventType)
		}

		finalState := lastEvent.GameState
		if finalState.Stats.WinTeam.isZero() {
			t.Error("game completed but no winning team")
		}
		if finalState.Stats.Points[finalState.Stats.WinTeam] < 60 {
			t.Errorf("winning team has %d points, expected >= 60",
				finalState.Stats.Points[finalState.Stats.WinTeam])
		}

		t.Logf("All-AI game completed! %v wins with %d points in %d rounds",
			finalState.Stats.WinTeam,
			finalState.Stats.Points[finalState.Stats.WinTeam],
			len(finalState.Stats.Rounds))
	})

	t.Run("mixed human and AI players", func(t *testing.T) {
		// Player1 (human), Player2 (AI), Player3 (human), Player4 (AI)
		g := NewGame("Alice", "Bot2", "Charlie", "Bot4", "Team A", "Team B", false, true, false, true)

		events, err := g.Start()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Game should not be complete yet — human players need to act
		lastEvent := events[len(events)-1]
		if lastEvent.EventType == CardPlayedAndGameCompletedEvent {
			t.Fatal("game should not complete on start with human players present")
		}

		// Play the game to completion, only providing actions for human players
		maxActions := 1000
		actionCount := 0
		gameCompleted := false

		for !gameCompleted && actionCount < maxActions {
			curState := g.State()

			// If no tricks yet, the starter needs to assign trump
			if len(curState.Round.Tricks) == 0 {
				starter := curState.Round.Trumper
				playerInfo := curState.PlayersInfo[starter]

				if playerInfo.IsAI {
					t.Fatal("AI starter should have assigned trump automatically")
				}

				events, err = g.Apply(Action{
					Name:   AssignTrumpAction,
					Player: starter,
					Suit:   allSuits[actionCount%len(allSuits)],
				})
				if err != nil {
					t.Fatalf("error assigning trump: %v", err)
				}
				actionCount++
			} else {
				// Find next player to act
				tricks := curState.Round.Tricks
				curTrick := tricks[len(tricks)-1]
				nextPlayer := curTrick.Next

				if nextPlayer.isZero() {
					t.Fatal("no next player but game not completed")
				}

				playerInfo := curState.PlayersInfo[nextPlayer]
				if playerInfo.IsAI {
					t.Fatalf("AI player %v should have played automatically", nextPlayer)
				}

				hand := curState.Round.getHand(nextPlayer)
				var playable HandCard
				for _, c := range hand.Cards {
					if c.IsPlayable {
						playable = c
						break
					}
				}

				events, err = g.Apply(Action{
					Name:   PlayCardAction,
					Player: nextPlayer,
					Rank:   playable.Card.Rank,
					Suit:   playable.Card.Suit,
				})
				if err != nil {
					t.Fatalf("error playing card: %v", err)
				}
				actionCount++
			}

			for _, e := range events {
				if e.EventType == CardPlayedAndGameCompletedEvent {
					gameCompleted = true
					finalState := e.GameState
					if finalState.Stats.WinTeam.isZero() {
						t.Error("game completed but no winning team")
					}
					t.Logf("Mixed game completed! %v wins with %d points in %d rounds (human actions: %d)",
						finalState.Stats.WinTeam,
						finalState.Stats.Points[finalState.Stats.WinTeam],
						len(finalState.Stats.Rounds),
						actionCount)
				}
			}
		}

		if !gameCompleted {
			t.Fatalf("game did not complete within %d actions", maxActions)
		}
	})

	t.Run("AI actions are included in event batch", func(t *testing.T) {
		// Player1 human, rest are AI — after human plays, AI should chain actions
		g := NewGame("Alice", "Bot2", "Bot3", "Bot4", "Team A", "Team B", false, true, true, true)

		events, _ := g.Start()
		state := events[len(events)-1].GameState
		starter := state.Round.Trumper

		if g.playersInfo[starter].IsAI {
			// AI starter should have already assigned trump during Start
			// and potentially played several cards
			if len(events) < 3 {
				t.Errorf("expected AI to generate multiple events during Start, got %d", len(events))
			}
		} else {
			// Human is starter, assign trump — AI players should chain
			events, err := g.Apply(Action{
				Name:   AssignTrumpAction,
				Player: starter,
				Suit:   DiamondsSuit,
			})
			if err != nil {
				t.Fatalf("error assigning trump: %v", err)
			}

			// After human assigns trump, AI players (2,3,4) should auto-play
			// We expect TrumpAssigned + TrickStarted + at least some CardPlayed events
			if len(events) < 2 {
				t.Errorf("expected AI to chain actions after trump assignment, got %d events", len(events))
			}

			// Verify the next player waiting for input is a human
			lastState := events[len(events)-1].GameState
			tricks := lastState.Round.Tricks
			if len(tricks) > 0 {
				nextPlayer := tricks[len(tricks)-1].Next
				if !nextPlayer.isZero() {
					playerInfo := lastState.PlayersInfo[nextPlayer]
					if playerInfo.IsAI {
						t.Errorf("expected next player to be human, but got AI player %v", nextPlayer)
					}
				}
			}
		}
	})
}
