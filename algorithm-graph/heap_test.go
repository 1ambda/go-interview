package graph

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloatHeap(t *testing.T) {
	h := &FloatMinHeap{2.0, 1.0, 5.0}

	heap.Init(h)
	assert.Equal(t, 1.0, h.Min())

	heap.Push(h, 3.0)
	assert.Equal(t, 1.0, h.Min())

	min := heap.Pop(h)
	assert.Equal(t, 1.0, min)

	assert.Equal(t, 2.0, h.Min())
}
