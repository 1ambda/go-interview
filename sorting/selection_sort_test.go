package sorting

import (
	"testing"
	"sort"
	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	arr := RandomArray(30)
	sorted := SelectionSort(arr)

	assert.True(t, sort.IsSorted(sort.IntSlice(sorted)))
}
