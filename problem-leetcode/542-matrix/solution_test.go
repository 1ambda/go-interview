package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {

	input1 := [][]int {
		[]int{0,0,0},
		[]int{0,1,0},
		[]int{0,0,0},
	}

	expected1 := [][]int {
		[]int{0,0,0},
		[]int{0,1,0},
		[]int{0,0,0},
	}

	assert.Equal(t, expected1, updateMatrix(input1))

	input2 := [][]int {
		[]int{0,0,0},
		[]int{0,1,0},
		[]int{1,1,1},
	}

	expected2 := [][]int {
		[]int{0,0,0},
		[]int{0,1,0},
		[]int{1,2,1},
	}

	assert.Equal(t, expected2, updateMatrix(input2))
}
