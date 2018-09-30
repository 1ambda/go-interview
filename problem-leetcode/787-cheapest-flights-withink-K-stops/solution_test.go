package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {

	input1 := [][]int{
		{2, 1, 1,},
		{2, 3, 1,},
		{3, 4, 1,},
	}

	n1 := 4
	k1 := 2

	expected1 := 2
	action1 := networkDelayTime(input1, n1, k1)
	assert.Equal(t, expected1, action1)

	input2 := [][]int{
		{1, 2, 1,},
		{2, 3, 7,},
		{1, 3, 4,},
		{2, 1, 2,},
	}

	n2 := 4
	k2 := 1

	expected2 := -1 // N == 4, but len(nodes) == 3
	action2 := networkDelayTime(input2, n2, k2)
	assert.Equal(t, expected2, action2)
}
