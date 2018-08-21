package sorting

import (
	"testing"
	"sort"
	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	arr := RandomArray(30)
	sorted := InsertionSort(arr)

	assert.True(t, sort.IsSorted(sort.IntSlice(sorted)))
}
