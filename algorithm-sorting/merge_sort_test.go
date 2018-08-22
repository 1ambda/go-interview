package sorting

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"sort"
	"fmt"
)

func TestMergeSort(t *testing.T) {
	arr := RandomArray(10000000)
	sorted := MergeSort(arr)

	assert.True(t, sort.IsSorted(sort.IntSlice(sorted)))

	fmt.Printf("before length: %d\n", len(arr))
	fmt.Printf("after  length: %d\n", len(sorted))
}