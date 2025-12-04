package tui

import (
	"gyla/models"
	"sort"
	"strings"
)

var plainCardsOrder = map[models.Suit]int{
	models.ClubsSuit:    4,
	models.SpadesSuit:   3,
	models.HeartsSuit:   2,
	models.DiamondsSuit: 1,
}

func Card(c models.Card) string {
	var underline = "\033[4m"
	var red = "\033[31m"
	var black = "\033[30m"
	var reset = "\033[0m"
	var color string

	if !c.IsTrump {
		underline = ""
	}

	if c.Suit == models.HeartsSuit || c.Suit == models.DiamondsSuit {
		color = red
	} else {
		color = black
	}

	return underline + color + string(c.Rank) + c.Suit.Tui() + reset
}

func Suit(s models.Suit) string {
	var red = "\033[31m"
	var black = "\033[30m"
	var reset = "\033[0m"
	var color string

	if s == models.HeartsSuit || s == models.DiamondsSuit {
		color = red
	} else {
		color = black
	}

	var suitSymbol string
	switch s {
	case models.ClubsSuit:
		suitSymbol = "♣"
	case models.SpadesSuit:
		suitSymbol = "♠"
	case models.HeartsSuit:
		suitSymbol = "♥"
	case models.DiamondsSuit:
		suitSymbol = "♦"
	}

	return color + suitSymbol + reset
}

func Cards(cards []models.Card) string {
	copied := append([]models.Card{}, cards...)
	sort.Slice(copied, func(i, j int) bool {
		a := copied[i]
		b := copied[j]

		if !a.IsTrump && !b.IsTrump && a.Suit != b.Suit {
			return plainCardsOrder[a.Suit] > plainCardsOrder[b.Suit]
		} else {
			return a.Level() > b.Level()
		}
	})

	strCards := make([]string, len(copied))
	for i, c := range copied {
		strCards[i] = Card(c)
	}

	return strings.Join(strCards, " ")
}
