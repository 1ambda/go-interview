package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {

	input1 := [][]int {
		{1,2},
		{1,3},
		{2,3},
	}

	expected1 := []int {2, 3}

	assert.Equal(t, expected1, findRedundantConnection(input1))

	input2 := [][]int {
		{1,2},
		{2,3},
		{3,4},
		{1,4},
		{1,5},
	}

	expected2 := []int {1, 4}

	assert.Equal(t, expected2, findRedundantConnection(input2))
}
