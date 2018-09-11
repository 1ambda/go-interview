package solution

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// write efficient
func TestSolution(t *testing.T) {
	input1 := []int{1, 5, 3, 3, 7}
	output1 := Solution(input1)

	assert.True(t, output1)

	input2 := []int{1, 3, 5, 3, 4}
	output2 := Solution(input2)

	assert.False(t, output2)

	// already sorted
	input3 := []int{1, 3, 5}
	output3 := Solution(input3)

	assert.True(t, output3)

	input4 := []int{1, 3, 2, 4}
	output4 := Solution(input4)

	assert.True(t, output4)

	input5 := []int{1, 4, 2, 3}
	output5 := Solution(input5)

	assert.False(t, output5)
}
