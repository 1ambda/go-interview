package codility

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	arr1 := []int{3, 8, 9, 7, 6}
	exp1 := []int{9, 7, 6, 3, 8}
	ret1 := Solution(arr1, 3)
	assert.Equal(t, exp1, ret1)

	var arr2 []int
	var exp2 []int
	ret2 := Solution(arr2, 3)
	assert.Equal(t, exp2, ret2)
}
