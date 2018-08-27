package codility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {

	input := []int{0, 1, 0, 1, 1}
	expected := 5
	assert.Equal(t, expected, Solution(input))
}
