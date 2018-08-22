package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	arr1 := []int{2, 7, 11, 15}

	assert.Equal(t, []int{0, 1}, twoSum(arr1, 9))
	assert.Equal(t, []int{0, 1}, twoSumHashTable(arr1, 9))
}
