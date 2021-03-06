package graph

// https://golang.org/pkg/container/heap/
type FloatMinHeap []float64

func (h FloatMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h FloatMinHeap) Len() int {
	return len(h)
}

func (h FloatMinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h FloatMinHeap) Min() float64 {
	return h[0]
}

func (h *FloatMinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(float64))
}

func (h *FloatMinHeap) Pop() interface{} {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
