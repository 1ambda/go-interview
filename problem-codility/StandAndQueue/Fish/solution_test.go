package codility

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {

	fishes := []int{4, 3, 2, 1, 5}
	directions := []int{0, 1, 0, 0, 0}

	actual := Solution(fishes, directions)
	expected := 2
	assert.Equal(t, expected, actual)
}
