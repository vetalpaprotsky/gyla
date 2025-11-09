package models

import (
	"errors"
	"gyla/testutils"
	"testing"
)

func TestNewCard(t *testing.T) {
	tests := []struct {
		Rank, Suit string
		want       *Card
		wantErr    error
	}{
		{"10", "H", &Card{Rank: "10", Suit: "H", isTrump: false}, nil},
		{"K", "S", &Card{Rank: "K", Suit: "S", isTrump: false}, nil},
		{"7", "C", &Card{Rank: "7", Suit: "C", isTrump: true}, nil},
		{"J", "D", &Card{Rank: "J", Suit: "D", isTrump: true}, nil},
		{"5", "S", nil, errors.New("Invalid Rank: 5")},
		{"Q", "K", nil, errors.New("Invalid Suit: K")},
	}

	for _, tt := range tests {
		name := testutils.FunctionCallName("newCard", tt.Rank, tt.Suit)
		t.Run(name, func(t *testing.T) {
			got, err := newCard(tt.Rank, tt.Suit)

			if errMsg := testutils.TestErr(err, tt.wantErr); errMsg != "" {
				t.Error(errMsg)
			} else if errMsg := testutils.TestGotPtr(got, tt.want); errMsg != "" {
				t.Error(errMsg)
			}
		})
	}
}

func TestLevel(t *testing.T) {
	tests := []struct {
		card Card
		want int
	}{
		{Card{Rank: "Q", Suit: "H"}, 5},
		{Card{Rank: "Q", Suit: "H", isTrump: true}, 11},
		{Card{Rank: "9", Suit: "C"}, 3},
		{Card{Rank: "9", Suit: "C", isTrump: true}, 9},
		{Card{Rank: "6", Suit: "S"}, 1},
		{Card{Rank: "6", Suit: "S", isTrump: true}, 22},
		{Card{Rank: "7", Suit: "H"}, 19},
		{Card{Rank: "7", Suit: "H", isTrump: true}, 19},
		{Card{Rank: "J", Suit: "D"}, 14},
		{Card{Rank: "J", Suit: "D", isTrump: true}, 14},
	}

	for _, tt := range tests {
		name := testutils.MethodCallName(tt.card, "level")
		t.Run(name, func(t *testing.T) {
			got := tt.card.level()
			if errMsg := testutils.TestGot(got, tt.want); errMsg != "" {
				t.Error(errMsg)
			}
		})
	}
}
