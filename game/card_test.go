package game

import (
	"testing"

	"github.com/vetalpaprotsky/gyla/testutils"
)

func TestNewCard(t *testing.T) {
	tests := []struct {
		rank    Rank
		suit    Suit
		want    card
		wantErr error
	}{
		{TenRank, HeartsSuit, card{rank: TenRank, suit: HeartsSuit, isTrump: false}, nil},
		{KingRank, SpadesSuit, card{rank: KingRank, suit: SpadesSuit, isTrump: false}, nil},
		{SevenRank, ClubsSuit, card{rank: SevenRank, suit: ClubsSuit, isTrump: true}, nil},
		{JackRank, DiamondsSuit, card{rank: JackRank, suit: DiamondsSuit, isTrump: true}, nil},
		{Rank(99), SpadesSuit, card{}, newInvalidRankError(Rank(99))},
		{QueenRank, Suit(99), card{}, newInvalidSuitError(Suit(99))},
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
		card card
		want int
	}{
		{card{rank: QueenRank, suit: HeartsSuit}, 5},
		{card{rank: QueenRank, suit: HeartsSuit, isTrump: true}, 11},
		{card{rank: NineRank, suit: ClubsSuit}, 3},
		{card{rank: NineRank, suit: ClubsSuit, isTrump: true}, 9},
		{card{rank: SixRank, suit: SpadesSuit}, 1},
		{card{rank: SixRank, suit: SpadesSuit, isTrump: true}, 22},
		{card{rank: SevenRank, suit: HeartsSuit}, 19},
		{card{rank: SevenRank, suit: HeartsSuit, isTrump: true}, 19},
		{card{rank: JackRank, suit: DiamondsSuit}, 14},
		{card{rank: JackRank, suit: DiamondsSuit, isTrump: true}, 14},
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
