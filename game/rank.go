package game

type Rank int

const (
	SixRank   = Rank(1)
	SevenRank = Rank(2)
	EightRank = Rank(3)
	NineRank  = Rank(4)
	TenRank   = Rank(5)
	JackRank  = Rank(6)
	QueenRank = Rank(7)
	KingRank  = Rank(8)
	AceRank   = Rank(9)
)

var allRanks = [9]Rank{
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

func (s Rank) isZero() bool {
	return s == Rank(0)
}

func (r Rank) isValid() bool {
	for _, rank := range allRanks {
		if r == rank {
			return true
		}
	}

	return false
}
