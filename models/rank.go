package models

type Rank string

func (r Rank) IsValid() bool {
	for _, validRank := range ValidRanks {
		if r == validRank {
			return true
		}
	}

	return false
}
