package solution

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// write efficient
func TestSolution(t *testing.T) {
	input1 := []int{7, 3, 7, 3, 1, 3, 4, 1}
	output1 := Solution(input1)
	expected1 := 5

	assert.Equal(t, expected1, output1)

	input2 := []int{2, 1, 1, 3, 2, 1, 1, 3}
	output2 := Solution(input2)
	expected2 := 3

	assert.Equal(t, expected2, output2)

	input3 := []int{7, 5, 2, 7, 2, 7, 4, 7}
	output3 := Solution(input3)
	expected3 := 6

	assert.Equal(t, expected3, output3)
}
