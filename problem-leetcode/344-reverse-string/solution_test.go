package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, "fdsa", reverseString("asdf"))
	assert.Equal(t, "f̅ds⃝a", reverseString("as⃝df̅"))

	assert.Equal(t, "fdsa", reverseStringSimple("asdf"))
	assert.NotEqual(t, "f̅ds⃝a", reverseStringSimple("as⃝df̅"))
}
