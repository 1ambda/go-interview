package dijkstra

// Edge struct hold a target vertex and it's weight
type Neighbor struct {
	Destination *Vertex
	Weight      float64
}

type Vertex struct {
	Neighbors   []*Neighbor
	Name        string
	Visited     bool
	Previous    *Vertex
	Distance    float64 // tentative distance
	RemoveIndex int     // index in priority queue
}

// implement priority queue for Vertex
type VertexPriorityQueue []*Vertex

func (q VertexPriorityQueue) Len() int {
	return len(q)
}

func (q VertexPriorityQueue) Less(i, j int) bool {
	return q[i].Distance < q[j].Distance
}

func (q VertexPriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].RemoveIndex = i
	q[j].RemoveIndex = j
}

func (q *VertexPriorityQueue) Push(x interface{}) {
	vertex := x.(*Vertex)
	vertex.RemoveIndex = len(*q)
	*q = append(*q, vertex)
}

func (q *VertexPriorityQueue) Pop() interface{} {
	old := *q
	last := len(old) - 1
	removed := old[last]
	removed.RemoveIndex = -1

	*q = old[:last]
	return removed
}
