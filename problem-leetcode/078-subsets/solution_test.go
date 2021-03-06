package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {

	input1 := []int{1, 2, 3}
	expected1 := [][]int{
		{3},
		{1},
		{2},
		{1, 2, 3},
		{1, 3,},
		{2, 3,},
		{1, 2,},
		{},
	}

	actual1 := subsets(input1)
	assert.ElementsMatch(t, expected1, actual1)
}
