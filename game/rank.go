package game

type Rank string

// TODO: We could store them as bites(chars)
const (
	SixRank   = Rank("6")
	SevenRank = Rank("7")
	EightRank = Rank("8")
	NineRank  = Rank("9")
	TenRank   = Rank("10")
	JackRank  = Rank("J")
	QueenRank = Rank("Q")
	KingRank  = Rank("K")
	AceRank   = Rank("A")
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

func (r Rank) isValid() bool {
	for _, validRank := range validRanks {
		if r == validRank {
			return true
		}
	}

	return false
}
