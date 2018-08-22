package sorting

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"sort"
	"github.com/1ambda/go-interview/util"
)

func TestQuickPartition(t *testing.T) {
	arr := util.RandomArray(5)
	QuickPartition(arr, 0, len(arr) - 1)
}

func TestQuickSort(t *testing.T) {
	arr := util.RandomArray(10000000)
	sorted := QuickSort(arr, 0, len(arr) - 1)

	assert.True(t, sort.IsSorted(sort.IntSlice(sorted)))
}