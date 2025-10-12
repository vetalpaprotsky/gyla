package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSuitValid(t *testing.T) {
	assert.Equal(t, true, isSuitValid("C"), "valid")
	assert.Equal(t, false, isSuitValid("G"), "invalid")
}

func TestIsRankValid(t *testing.T) {
	assert.Equal(t, true, isRankValid("6"), "valid")
	assert.Equal(t, true, isRankValid("K"), "valid")
	assert.Equal(t, true, isRankValid("10"), "valid")
	assert.Equal(t, false, isRankValid("5"), "invalid")
}

// TODO add more tests
