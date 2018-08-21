package sorting

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"sort"
)

func TestQuickPartition(t *testing.T) {
	arr := RandomArray(5)
	QuickPartition(arr, 0, len(arr) - 1)
}

func TestQuickSort(t *testing.T) {
	arr := RandomArray(10000000)
	sorted := QuickSort(arr, 0, len(arr) - 1)

	assert.True(t, sort.IsSorted(sort.IntSlice(sorted)))
}