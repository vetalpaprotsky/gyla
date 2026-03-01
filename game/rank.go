package game

type Rank int

const (
	SixRank   = 1
	SevenRank = 2
	EightRank = 3
	NineRank  = 4
	TenRank   = 5
	JackRank  = 6
	QueenRank = 7
	KingRank  = 8
	AceRank   = 9
)

var validRanks = [9]Rank{
	SixRank,
	SevenRank,
	EightRank,
	NineRank,
	TenRank,
	JackRank,
	QueenRank,
	KingRank,
	AceRank,
}

func (s Rank) IsZero() bool {
	return s == Rank(0)
}

func (r Rank) isValid() bool {
	for _, validRank := range validRanks {
		if r == validRank {
			return true
		}
	}

	return false
}
