package game

import (
	"math/rand/v2"
)

// NOTE: This AI is AI-generated :D

func applyAIActions(g *Game) {
	for {
		if !applyAIAction(g) {
			return
		}
	}
}

func applyAIAction(g *Game) bool {
	action := getAIAction(g)
	if action.isZero() {
		return false
	}

	err := g.apply(action)
	if err != nil {
		panic("ai action failed: " + err.Error())
	}

	return true
}

func getAIAction(g *Game) Action {
	next := g.nextAction()
	if next.isZero() || !g.getParticipant(next.Player).IsAI {
		return Action{}
	}

	curRound := g.currentRound()

	switch next.Name {
	case AssignTrumpAction:
		suit := chooseTrumpSuit(curRound.getHand(next.Player))
		return Action{
			Name:   AssignTrumpAction,
			Player: next.Player,
			Suit:   suit,
		}
	case PlayCardAction:
		return chooseCard(curRound, next.Player)
	default:
		return Action{}
	}
}

// chooseTrumpSuit picks the suit where the player has the most non-default-trump
// cards. Having more cards of a suit means the trump 6 and suit cards become
// powerful, and you are more likely to control tricks in that suit.
// Ties are broken by the total level sum.
func chooseTrumpSuit(h *hand) Suit {
	type suitScore struct {
		suit  Suit
		count int
		value int
	}

	scores := make([]suitScore, len(allSuits))
	for i, s := range allSuits {
		scores[i].suit = s
		for _, c := range h.cards {
			if c.IsDefaultTrump() {
				continue
			}
			if c.Suit == s {
				scores[i].count++
				scores[i].value += int(c.Rank)
			}
		}
	}

	best := scores[0]
	for _, sc := range scores[1:] {
		if sc.count > best.count || (sc.count == best.count && sc.value > best.value) {
			best = sc
		}
	}

	return best.suit
}

// --- Round analysis helpers ---

// allPlayedCards returns every card that has been played in the round so far,
// across all completed and in-progress tricks.
func allPlayedCards(r round) []PlayedCard {
	var played []PlayedCard
	for _, t := range r.tricks {
		played = append(played, t.playedCards...)
	}
	return played
}

// suitHasBeenPlayed checks if a specific non-trump card (by rank+suit) has
// already been played in the round.
func cardHasBeenPlayed(r round, rank Rank, suit Suit) bool {
	for _, pc := range allPlayedCards(r) {
		if pc.Card.Rank == rank && pc.Card.Suit == suit {
			return true
		}
	}
	return false
}

// countCardsOfSuitInHand counts non-trump cards of a given suit in a hand.
func countCardsOfSuitInHand(h *hand, suit Suit) int {
	count := 0
	for _, c := range h.cards {
		if !c.IsTrump && c.Suit == suit {
			count++
		}
	}
	return count
}

// isLastCardOfSuit checks whether a card is the only remaining non-trump card
// of its suit in the player's hand.
func isLastCardOfSuit(h *hand, c Card) bool {
	if c.IsTrump {
		return false
	}
	return countCardsOfSuitInHand(h, c.Suit) == 1
}

// teammateVoidSuits inspects completed tricks to find suits where the teammate
// signaled they are void. A signal is: teammate led (started) a trick with a
// high non-ace non-trump card (King, Queen, or 10), which hints it's their
// last card of that suit. We also detect when a teammate couldn't follow suit
// (played a different suit or a trump when they should have followed).
func teammateVoidSuits(r round, player Player) map[Suit]bool {
	teammate := player.teammate()
	voids := map[Suit]bool{}

	for _, t := range r.tricks {
		if !t.isCompleted() && len(t.playedCards) < 2 {
			continue
		}

		// Detect explicit signal: teammate led with a high non-ace card.
		// This means "I have no more of this suit, please lead it so I can trump."
		if t.starter == teammate && len(t.playedCards) > 0 {
			leadCard := t.playedCards[0].Card
			if !leadCard.IsTrump && leadCard.Rank != AceRank {
				// High card lead (King, Queen, 10) = signal
				if leadCard.Rank == KingRank || leadCard.Rank == QueenRank || leadCard.Rank == TenRank {
					voids[leadCard.Suit] = true
				}
			}
		}

		// Detect when teammate couldn't follow suit (played off-suit or trump
		// when the lead was a non-trump suit).
		if len(t.playedCards) >= 2 {
			leadCard := t.playedCards[0].Card
			if !leadCard.IsTrump {
				for _, pc := range t.playedCards[1:] {
					if pc.Player == teammate {
						// Teammate played a different suit or a trump instead of following
						if pc.Card.IsTrump || pc.Card.Suit != leadCard.Suit {
							voids[leadCard.Suit] = true
						}
					}
				}
			}
		}
	}

	return voids
}

// opponentVoidSuits detects suits where opponents showed they are void
// (couldn't follow suit).
func opponentVoidSuits(r round, player Player) map[Suit]bool {
	opp1 := player.leftOpponent()
	opp2 := player.rightOpponent()
	voids := map[Suit]bool{}

	for _, t := range r.tricks {
		if len(t.playedCards) < 2 {
			continue
		}

		leadCard := t.playedCards[0].Card
		if leadCard.IsTrump {
			continue
		}

		for _, pc := range t.playedCards[1:] {
			if pc.Player == opp1 || pc.Player == opp2 {
				if pc.Card.IsTrump || pc.Card.Suit != leadCard.Suit {
					voids[leadCard.Suit] = true
				}
			}
		}
	}

	return voids
}

// --- Card Play ---

func chooseCard(curRound round, player Player) Action {
	curTrick := curRound.currentTrick()
	h := curRound.getHand(player)
	playable := h.playableCardsFor(curTrick)
	if len(playable) == 0 {
		panic("AI: no playable cards for expected next player")
	}

	var chosen Card
	if curTrick.isEmpty() {
		chosen = chooseLead(h, playable, curRound, player)
	} else {
		chosen = chooseFollow(h, playable, curTrick, player, curRound)
	}

	return Action{
		Name:   PlayCardAction,
		Player: player,
		Rank:   chosen.Rank,
		Suit:   chosen.Suit,
	}
}

// chooseLead picks a card when the AI is leading the trick.
//
// Strategy priorities:
//  1. Lead a suit where teammate is void so they can trump in.
//  2. Lead with Aces (safe, high chance of winning).
//  3. Lead with Kings only if the Ace is already played OR it's the last card
//     of that suit (signaling teammate).
//  4. If we have lots of trumps, flush opponents by leading high trump.
//  5. Lead the last card of a suit as a signal to teammate.
//  6. Fall back to lowest non-trump, then lowest trump.
func chooseLead(h *hand, playable []Card, r round, player Player) Card {
	nonTrumps := filterNonTrumps(playable)
	trumps := filterTrumps(playable)
	tmVoids := teammateVoidSuits(r, player)
	oppVoids := opponentVoidSuits(r, player)

	// 1. Lead a suit where teammate is void (so they can trump in),
	//    but avoid suits where opponents are also void (they'd trump too).
	//    Pick our highest card in that suit to try to win or at least force
	//    opponents to use a trump.
	if len(tmVoids) > 0 && len(nonTrumps) > 0 {
		var candidates []Card
		for _, c := range nonTrumps {
			if tmVoids[c.Suit] && !oppVoids[c.Suit] {
				candidates = append(candidates, c)
			}
		}
		if len(candidates) > 0 {
			return highestCard(candidates)
		}
	}

	// 2. Lead with Aces — safest high card lead, very likely to win.
	for _, c := range nonTrumps {
		if c.Rank == AceRank {
			return c
		}
	}

	// 3. Lead with a King, but only if:
	//    a) The Ace of that suit has already been played (good chance to win), OR
	//    b) It's the last card of that suit (signal to teammate).
	for _, c := range nonTrumps {
		if c.Rank == KingRank {
			aceGone := cardHasBeenPlayed(r, AceRank, c.Suit)
			lastOfSuit := isLastCardOfSuit(h, c)
			if aceGone || lastOfSuit {
				return c
			}
		}
	}

	// 4. If we have 4+ trumps, lead with the highest trump to flush out
	//    opponent trumps.
	if len(trumps) >= 4 {
		return highestCard(trumps)
	}

	// 5. Signal: if we have a card that is the last of its suit, play it
	//    (high-ish card preferred) to tell teammate we're void after this.
	var signals []Card
	for _, c := range nonTrumps {
		if isLastCardOfSuit(h, c) && (c.Rank == QueenRank || c.Rank == TenRank || c.Rank == KingRank) {
			signals = append(signals, c)
		}
	}
	if len(signals) > 0 {
		return highestCard(signals)
	}

	// 6. Lead with the lowest non-trump to minimize risk.
	if len(nonTrumps) > 0 {
		return lowestCard(nonTrumps)
	}

	// 7. Only trumps left — play the lowest to conserve high ones.
	return lowestCard(playable)
}

// chooseFollow picks a card when the AI is following in a trick.
func chooseFollow(h *hand, playable []Card, t trick, player Player, r round) Card {
	teammate := player.teammate()
	currentWinner := currentTrickWinner(t)
	teammateIsWinning := currentWinner == teammate

	if teammateIsWinning {
		// Teammate is winning — play the lowest card to not waste resources.
		return lowestCard(playable)
	}

	// Try to win the trick with the cheapest winning card.
	winningCards := findWinningCards(playable, t)
	if len(winningCards) > 0 {
		return lowestCard(winningCards)
	}

	// Can't win — dump strategically.
	// Prefer dumping a card from a suit we want to get void in (especially
	// a suit where opponents have shown strength), so we can trump later.
	return dumpCard(h, playable, r, player)
}

// dumpCard chooses the best card to throw away when we can't win.
// Priorities:
// - Dump from a suit where we have few cards (getting void = can trump later).
// - Among those, dump the lowest card.
func dumpCard(h *hand, playable []Card, r round, player Player) Card {
	nonTrumps := filterNonTrumps(playable)

	// Prefer dumping non-trump cards to preserve trumps.
	if len(nonTrumps) > 0 {
		// Find the suit where we have the fewest cards — getting void faster.
		type suitGroup struct {
			suit  Suit
			count int
		}
		suitCounts := map[Suit]int{}
		for _, c := range nonTrumps {
			suitCounts[c.Suit] = countCardsOfSuitInHand(h, c.Suit)
		}

		// Find the minimum count.
		minCount := 999
		for _, cnt := range suitCounts {
			if cnt < minCount {
				minCount = cnt
			}
		}

		// Among suits with the fewest cards, pick the lowest card.
		var candidates []Card
		for _, c := range nonTrumps {
			if suitCounts[c.Suit] == minCount {
				candidates = append(candidates, c)
			}
		}
		if len(candidates) > 0 {
			return lowestCard(candidates)
		}

		return lowestCard(nonTrumps)
	}

	return lowestCard(playable)
}

// currentTrickWinner determines who is currently winning an in-progress trick.
func currentTrickWinner(t trick) Player {
	if len(t.playedCards) == 0 {
		return Player(0)
	}

	winPlayer := t.playedCards[0].Player
	winCard := t.playedCards[0].Card

	if hasAnyTrumpsInPlayed(t.playedCards) {
		for _, pc := range t.playedCards {
			if pc.Card.level() > winCard.level() {
				winPlayer = pc.Player
				winCard = pc.Card
			}
		}
	} else {
		leadingSuit := t.playedCards[0].Card.Suit
		for _, pc := range t.playedCards {
			if pc.Card.Suit == leadingSuit && pc.Card.level() > winCard.level() {
				winPlayer = pc.Player
				winCard = pc.Card
			}
		}
	}

	return winPlayer
}

// findWinningCards returns the subset of playable cards that would currently
// win the trick if played.
func findWinningCards(playable []Card, t trick) []Card {
	if len(t.playedCards) == 0 {
		return playable
	}

	// Figure out the current best card in the trick.
	bestCard := t.playedCards[0].Card
	hasTrumps := hasAnyTrumpsInPlayed(t.playedCards)

	if hasTrumps {
		for _, pc := range t.playedCards {
			if pc.Card.level() > bestCard.level() {
				bestCard = pc.Card
			}
		}
	} else {
		leadingSuit := t.playedCards[0].Card.Suit
		for _, pc := range t.playedCards {
			if pc.Card.Suit == leadingSuit && pc.Card.level() > bestCard.level() {
				bestCard = pc.Card
			}
		}
	}

	var winners []Card
	for _, c := range playable {
		if wouldBeat(c, bestCard, t) {
			winners = append(winners, c)
		}
	}

	return winners
}

// wouldBeat checks if candidate would beat current best in the context of the trick.
func wouldBeat(candidate, best Card, t trick) bool {
	leadingSuit := t.playedCards[0].Card.Suit
	hasTrumps := hasAnyTrumpsInPlayed(t.playedCards) || best.IsTrump

	if hasTrumps {
		if !candidate.IsTrump {
			return false
		}
		return candidate.level() > best.level()
	}

	// No trumps in trick so far.
	if candidate.IsTrump {
		return true
	}

	if candidate.Suit == leadingSuit && candidate.level() > best.level() {
		return true
	}

	return false
}

func hasAnyTrumpsInPlayed(played []PlayedCard) bool {
	for _, pc := range played {
		if pc.Card.IsTrump {
			return true
		}
	}
	return false
}

// --- Helpers ---

func filterTrumps(cards []Card) []Card {
	var result []Card
	for _, c := range cards {
		if c.IsTrump {
			result = append(result, c)
		}
	}
	return result
}

func filterNonTrumps(cards []Card) []Card {
	var result []Card
	for _, c := range cards {
		if !c.IsTrump {
			result = append(result, c)
		}
	}
	return result
}

func highestCard(cards []Card) Card {
	best := cards[0]
	for _, c := range cards[1:] {
		if c.level() > best.level() {
			best = c
		}
	}
	return best
}

func lowestCard(cards []Card) Card {
	worst := cards[0]
	for _, c := range cards[1:] {
		if c.level() < worst.level() {
			worst = c
		}
	}
	return worst
}

// randomSuit returns a random suit — kept for potential future use.
func randomSuit() Suit {
	return allSuits[rand.IntN(4)]
}
