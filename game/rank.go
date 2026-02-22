package game

type Rank string

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

var validRanks = [ranksCount]Rank{
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
