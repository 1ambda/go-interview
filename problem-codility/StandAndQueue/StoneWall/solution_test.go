package codility

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {

	input := []int{8, 8, 5, 7, 9, 8, 7, 4, 8}

	actual := Solution(input)
	expected := 7
	assert.Equal(t, expected, actual)
}
