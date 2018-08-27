package codility

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	arr1 := []int{9, 3, 9, 3, 9, 7, 9}
	ret1 := Solution(arr1)
	assert.Equal(t, ret1, 7)
}
