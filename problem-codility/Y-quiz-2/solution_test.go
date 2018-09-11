package solution

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	input1 := []int{1, 0, 0, 1, 0, 0}
	output1 := Solution(input1)
	expected1 := 2

	assert.Equal(t, expected1, output1)

	input2 := []int{1, 1, 0, 1, 0, 0}
	output2 := Solution(input2)
	expected2 := 3

	assert.Equal(t, expected2, output2)

	input3 := []int{1, 1, 1, 1, 1, 0}
	output3 := Solution(input3)
	expected3 := 1

	assert.Equal(t, expected3, output3)

	input4 := []int{1, 1, 1, 1, 1, 1}
	output4 := Solution(input4)
	expected4 := 0

	assert.Equal(t, expected4, output4)
}
