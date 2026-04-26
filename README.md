# Gyla

A card game engine for a traditional trick-taking card game.

## Game Rules

### Players & Teams

- 4 players split into 2 teams: Player 1 + Player 3 = Team 1, Player 2 + Player 4 = Team 2.
- Players sit in a circle: P1 → P2 → P3 → P4 → P1 (left opponent order).
- Players can be human or AI.

### Cards

- **9 ranks:** 6, 7, 8, 9, 10, Jack, Queen, King, Ace.
- **4 suits:** Clubs, Spades, Hearts, Diamonds.
- 36 cards total, 9 dealt to each player.

### Trumps

There are two kinds of trumps:

1. **Default trumps** — all 7s and all Jacks are always trumps regardless of the chosen suit (8 cards total).
2. **Suit trumps** — when a player chooses a trump suit, all remaining cards of that suit also become trumps.

### Card Hierarchy

**Trump cards (highest to lowest):**

| # | Card |
|---|------|
| 1 | 6 of trump suit *(highest card in the game)* |
| 2 | 7 of Clubs |
| 3 | 7 of Spades |
| 4 | 7 of Hearts |
| 5 | 7 of Diamonds |
| 6 | Jack of Clubs |
| 7 | Jack of Spades |
| 8 | Jack of Hearts |
| 9 | Jack of Diamonds |
| 10 | Ace of trump suit |
| 11 | King of trump suit |
| 12 | Queen of trump suit |
| 13 | 10 of trump suit |
| 14 | 9 of trump suit |
| 15 | 8 of trump suit |

**Non-trump cards (within their suit, highest to lowest):**

Ace → King → Queen → 10 → 9 → 8 → 6

### Round Lifecycle

1. Cards are dealt (9 per player). If any player receives all four 7s or all four 6s, the hand is re-dealt.
2. The **starter** (trumper) assigns a trump suit.
3. 9 tricks are played (4 cards each, using all 36 cards).

### Determining the Starter

- **First round:** The player holding the 9 of Diamonds starts.
- **Subsequent rounds:**
  - If the team that assigned trump **won** the round, the same player assigns trump again.
  - If the team that assigned trump **lost**, the left opponent of the previous starter becomes the new starter.

### Trick Rules

1. The trick starter plays any card.
2. Other players must **follow suit:**
   - If the lead card is a **trump** → you must play a trump if you have one.
   - If the lead card is a **non-trump suit** → you must play a card of the same suit (non-trump only). If you don't have any, you can play anything.
3. **Trick winner:**
   - If any trumps were played → the highest trump wins.
   - If no trumps were played → the highest card of the leading suit wins.
4. The trick winner starts the next trick.

### Scoring

The team that wins more tricks (5+ out of 9) wins the round. Draws are impossible.

**Points awarded to the winning team per round:**

| Tricks Won | Points |
|-----------|--------|
| All 9 | 24 |
| 8 out of 9 | 18 |
| 5 to 7 | 6 |

**Bonus:** If the winning team is *not* the team that assigned trump, they receive an additional **+6 points**.

**Game end:** The game ends when a team reaches **60 points**.

## Usage

```go
package main

import "github.com/vetalpaprotsky/gyla/game"

func main() {
	g := game.NewGame(
		"Alice", "Bob", "Charlie", "Diana",
		"Team Alpha", "Team Beta",
		false, true, false, true, // players 2 and 4 are AI
	)

	events, err := g.Start()
	// handle events...

	events, err = g.Apply(game.Action{
		Name:   game.AssignTrumpAction,
		Player: game.Player1,
		Suit:   game.HeartsSuit,
	})
	// handle events...

	events, err = g.Apply(game.Action{
		Name:   game.PlayCardAction,
		Player: game.Player1,
		Rank:   game.AceRank,
		Suit:   game.HeartsSuit,
	})
	// handle events...
}
```

## Running Tests

```sh
cd game && go test -v ./...
```
