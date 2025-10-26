package models

import (
	"errors"
	"gyla/testutils"
	"testing"
)

func TestNewCard(t *testing.T) {
	tests := []struct {
		rank, suit string
		want       *Card
		wantErr    error
	}{
		{"10", "H", &Card{rank: "10", suit: "H", isTrump: false}, nil},
		{"K", "S", &Card{rank: "K", suit: "S", isTrump: false}, nil},
		{"7", "C", &Card{rank: "7", suit: "C", isTrump: true}, nil},
		{"J", "D", &Card{rank: "J", suit: "D", isTrump: true}, nil},
		{"5", "S", nil, errors.New("Invalid rank: 5")},
		{"Q", "K", nil, errors.New("Invalid suit: K")},
	}

	for _, tt := range tests {
		name := testutils.FunctionCallName("newCard", tt.rank, tt.suit)
		t.Run(name, func(t *testing.T) {
			got, err := newCard(tt.rank, tt.suit)

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
		{Card{rank: "Q", suit: "H"}, 5},
		{Card{rank: "Q", suit: "H", isTrump: true}, 11},
		{Card{rank: "9", suit: "C"}, 3},
		{Card{rank: "9", suit: "C", isTrump: true}, 9},
		{Card{rank: "6", suit: "S"}, 1},
		{Card{rank: "6", suit: "S", isTrump: true}, 22},
		{Card{rank: "7", suit: "H"}, 19},
		{Card{rank: "7", suit: "H", isTrump: true}, 19},
		{Card{rank: "J", suit: "D"}, 14},
		{Card{rank: "J", suit: "D", isTrump: true}, 13},
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
