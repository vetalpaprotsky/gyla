package models

import (
	"errors"
	"gyla/testutils"
	"testing"
)

func TestNewCard(t *testing.T) {
	tests := []struct {
		rank    Rank
		suit    Suit
		want    Card
		wantErr error
	}{
		{"10", "H", Card{Rank: "10", Suit: "H", IsTrump: false}, nil},
		{"K", "S", Card{Rank: "K", Suit: "S", IsTrump: false}, nil},
		{"7", "C", Card{Rank: "7", Suit: "C", IsTrump: true}, nil},
		{"J", "D", Card{Rank: "J", Suit: "D", IsTrump: true}, nil},
		{"5", "S", Card{}, errors.New("Invalid Rank: 5")},
		{"Q", "K", Card{}, errors.New("Invalid Suit: K")},
	}

	for _, tt := range tests {
		name := testutils.FunctionCallName("newCard", tt.rank, tt.suit)
		t.Run(name, func(t *testing.T) {
			got, err := newCard(tt.rank, tt.suit)

			if errMsg := testutils.TestErr(err, tt.wantErr); errMsg != "" {
				t.Error(errMsg)
			} else if errMsg := testutils.TestGot(got, tt.want); errMsg != "" {
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
		{Card{Rank: "Q", Suit: "H", IsTrump: true}, 11},
		{Card{Rank: "9", Suit: "C"}, 3},
		{Card{Rank: "9", Suit: "C", IsTrump: true}, 9},
		{Card{Rank: "6", Suit: "S"}, 1},
		{Card{Rank: "6", Suit: "S", IsTrump: true}, 22},
		{Card{Rank: "7", Suit: "H"}, 19},
		{Card{Rank: "7", Suit: "H", IsTrump: true}, 19},
		{Card{Rank: "J", Suit: "D"}, 14},
		{Card{Rank: "J", Suit: "D", IsTrump: true}, 14},
	}

	for _, tt := range tests {
		name := testutils.MethodCallName(tt.card, "level")
		t.Run(name, func(t *testing.T) {
			got := tt.card.Level()
			if errMsg := testutils.TestGot(got, tt.want); errMsg != "" {
				t.Error(errMsg)
			}
		})
	}
}
