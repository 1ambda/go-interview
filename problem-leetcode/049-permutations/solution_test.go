package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {

	input1 := []int{1, 2, 3}
	expected1 := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}

	actual1 := permute(input1)
	assert.Equal(t, expected1, actual1)

	input2 := []int{4, 5, 6}
	expected2 := [][]int{
		{4, 5, 6},
		{4, 6, 5},
		{5, 4, 6},
		{5, 6, 4},
		{6, 4, 5},
		{6, 5, 4},
	}

	actual2 := permute(input2)
	assert.Equal(t, expected2, actual2)
}
