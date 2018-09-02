package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	input1 := "ADOBECODEBANC"
	pattern1 := "ABC"
	expected := "BANC"

	assert.Equal(t, expected, minWindow(input1, pattern1))
}
