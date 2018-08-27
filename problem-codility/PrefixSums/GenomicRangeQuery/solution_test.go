package codility

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	p := []int{2, 5, 0}
	q := []int{4, 5, 6}
	result := Solution("CAGCCTA", p, q)
	assert.Equal(t, []int{2, 4, 1}, result)
}
