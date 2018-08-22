package sorting

import (
	"testing"
	"sort"
	"github.com/stretchr/testify/assert"
	"github.com/1ambda/go-interview/util"
)

func TestSelectionSort(t *testing.T) {
	arr := util.RandomArray(30)
	sorted := SelectionSort(arr)

	assert.True(t, sort.IsSorted(sort.IntSlice(sorted)))
}
