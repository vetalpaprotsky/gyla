package models

const (
	ClubsSuit    = Suit("C")
	SpadesSuit   = Suit("S")
	HeartsSuit   = Suit("H")
	DiamondsSuit = Suit("D")
)

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

var ValidSuits = [suitsCount]Suit{
	ClubsSuit,
	SpadesSuit,
	HeartsSuit,
	DiamondsSuit,
}

var ValidRanks = [ranksCount]Rank{
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

const suitsCount = 4
const ranksCount = 9
const cardsCount = 36
const playersCount = 4
const movesInTrickCount = playersCount
const cardsInHandCount = 9
const tricksPerRoundCount = cardsInHandCount
const maxPossibleNumberOfRounds = 60/6*2 - 1
