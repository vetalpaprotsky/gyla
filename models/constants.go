package models

const (
	ClubsSuit    = "C"
	SpadesSuit   = "S"
	HeartsSuit   = "H"
	DiamondsSuit = "D"
)

const (
	SixRank   = "6"
	SevenRank = "7"
	EightRank = "8"
	NineRank  = "9"
	TenRank   = "10"
	JackRank  = "J"
	QueenRank = "Q"
	KingRank  = "K"
	AceRank   = "A"
)

var ValidSuits = [suitsCount]string{
	ClubsSuit,
	SpadesSuit,
	HeartsSuit,
	DiamondsSuit,
}

var ValidRanks = [ranksCount]string{
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
const handsCount = playersCount
const cardsInHandCount = 9
const tricksPerRoundCount = cardsInHandCount
const maxPossibleNumberOfRounds = 60/6*2 - 1
