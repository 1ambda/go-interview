package sorting

import (
	"testing"
	"sort"
	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	arr := RandomArray(30)
	sorted := BubbleSort(arr)

	assert.True(t, sort.IsSorted(sort.IntSlice(sorted)))
}
