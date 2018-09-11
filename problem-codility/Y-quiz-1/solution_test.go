package solution

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	input1A := 53
	input1B := 1953786
	output1 := Solution(input1A, input1B)
	expected1 := 2

	assert.Equal(t, expected1, output1)

	input2A := 78
	input2B := 195378678
	output2 := Solution(input2A, input2B)
	expected2 := 4

	assert.Equal(t, expected2, output2)

	input3A := 57
	input3B := 153786
	output3 := Solution(input3A, input3B)
	expected3 := -1

	assert.Equal(t, expected3, output3)

	input4A := 0
	input4B := 0
	output4 := Solution(input4A, input4B)
	expected4 := 0

	assert.Equal(t, expected4, output4)

	input5A := 3786
	input5B := 153786
	output5 := Solution(input5A, input5B)
	expected5 := 2

	assert.Equal(t, expected5, output5)

	input6A := 1537
	input6B := 153786
	output6 := Solution(input6A, input6B)
	expected6 := 0

	assert.Equal(t, expected6, output6)

	input7A := 153
	input7B := 1527153
	output7 := Solution(input7A, input7B)
	expected7 := 4

	assert.Equal(t, expected7, output7)
}
