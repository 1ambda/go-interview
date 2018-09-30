package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {

	candidates1 := []int{2, 3, 6, 7}
	target1 := 7
	expected1 := [][]int{
		{7,},
		{2, 2, 3},
	}

	actual1 := combinationSum(candidates1, target1)
	assert.ElementsMatch(t, expected1, actual1)
}
