package game

const (
	clubsSuit    = Suit("C")
	spadesSuit   = Suit("S")
	heartsSuit   = Suit("H")
	diamondsSuit = Suit("D")
)

const (
	sixRank   = Rank("6")
	sevenRank = Rank("7")
	eightRank = Rank("8")
	nineRank  = Rank("9")
	tenRank   = Rank("10")
	jackRank  = Rank("J")
	queenRank = Rank("Q")
	kingRank  = Rank("K")
	aceRank   = Rank("A")
)

var validSuits = [suitsCount]Suit{
	clubsSuit,
	spadesSuit,
	heartsSuit,
	diamondsSuit,
}

var validRanks = [ranksCount]Rank{
	sixRank,
	sevenRank,
	eightRank,
	nineRank,
	tenRank,
	jackRank,
	queenRank,
	kingRank,
	aceRank,
}

const suitsCount = 4
const ranksCount = 9
const cardsCount = 36
const playersCount = 4
const movesPerTrickCount = playersCount
const cardsInHandCount = 9
const tricksPerRoundCount = cardsInHandCount
const maxPossibleNumberOfRounds = 60/6*2 - 1
