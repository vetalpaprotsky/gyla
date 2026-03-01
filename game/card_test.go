package game

import (
	"testing"

	"github.com/vetalpaprotsky/gyla/testutils"
)

func TestNewCard(t *testing.T) {
	tests := []struct {
		rank    Rank
		suit    Suit
		want    Card
		wantErr error
	}{
		{TenRank, HeartsSuit, Card{Rank: TenRank, Suit: HeartsSuit, IsTrump: false}, nil},
		{KingRank, SpadesSuit, Card{Rank: KingRank, Suit: SpadesSuit, IsTrump: false}, nil},
		{SevenRank, ClubsSuit, Card{Rank: SevenRank, Suit: ClubsSuit, IsTrump: true}, nil},
		{JackRank, DiamondsSuit, Card{Rank: JackRank, Suit: DiamondsSuit, IsTrump: true}, nil},
		{Rank(99), SpadesSuit, Card{}, newInvalidRankError(Rank(99))},
		{QueenRank, Suit(99), Card{}, newInvalidSuitError(Suit(99))},
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
		{Card{Rank: QueenRank, Suit: HeartsSuit}, 5},
		{Card{Rank: QueenRank, Suit: HeartsSuit, IsTrump: true}, 11},
		{Card{Rank: NineRank, Suit: ClubsSuit}, 3},
		{Card{Rank: NineRank, Suit: ClubsSuit, IsTrump: true}, 9},
		{Card{Rank: SixRank, Suit: SpadesSuit}, 1},
		{Card{Rank: SixRank, Suit: SpadesSuit, IsTrump: true}, 22},
		{Card{Rank: SevenRank, Suit: HeartsSuit}, 19},
		{Card{Rank: SevenRank, Suit: HeartsSuit, IsTrump: true}, 19},
		{Card{Rank: JackRank, Suit: DiamondsSuit}, 14},
		{Card{Rank: JackRank, Suit: DiamondsSuit, IsTrump: true}, 14},
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
