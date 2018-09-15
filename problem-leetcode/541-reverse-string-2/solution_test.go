package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, "bacdfeg", reverseStr("abcdefg", 2))
	assert.Equal(t, "a", reverseStr("a", 2))
	assert.Equal(t, "cbadefg", reverseStr("abcdefg", 3))
}
