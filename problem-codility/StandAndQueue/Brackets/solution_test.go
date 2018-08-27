package codility

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {

	assert.Equal(t, Solution(""), 1)
	assert.Equal(t, Solution("(()(())())"), 1)
	assert.Equal(t, Solution("()()"), 1)

	assert.Equal(t, Solution("("), 0)
	assert.Equal(t, Solution(")"), 0)
	assert.Equal(t, Solution("()("), 0)
	assert.Equal(t, Solution("())"), 0)
	assert.Equal(t, Solution("(()()"), 0)

	assert.Equal(t, Solution("(()(())())"), 1)
	assert.Equal(t, Solution("()()"), 1)

	assert.Equal(t, Solution("{[()()]}"), 1)
	assert.Equal(t, Solution("([)()]"), 0)
}
