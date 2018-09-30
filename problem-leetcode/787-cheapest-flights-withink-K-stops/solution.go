package leetcode

import (
	"container/heap"
	"fmt"
	"math"
)

func networkDelayTime(times [][]int, N int, K int) int {
	vertices, start := buildGraph(times, K)
	var end *Vertex = nil // visit all nodes

	paths := dijkstra(vertices, start, end)

	if len(vertices) != N {
		return -1
	}

	for _, vertex := range vertices {
		// not all vertices were visited
		if !vertex.Visited {
			return -1
		}
	}

	// no proper solution found
	if len(paths) == 0 {
		return -1
	}

	return paths[len(paths)-1].length
}

func buildGraph(edges [][]int, start int) ([]*Vertex, *Vertex) {

	startName := fmt.Sprintf("%d", start)

	all := make(map[string]*Vertex)
	for _, edge := range edges {
		from := fmt.Sprintf("%d", edge[0])
		to := fmt.Sprintf("%d", edge[1])
		weight := edge[2]

		v1 := all[from]
		v2 := all[to]

		// add previously unseen vertices
		if v1 == nil {
			v1 = &Vertex{Name: from}
			all[from] = v1
		}
		if v2 == nil {
			v2 = &Vertex{Name: to}
			all[to] = v2
		}

		v1.Neighbors = append(v1.Neighbors, &Neighbor{Destination: v2, Weight: weight})
	}

	vertices := make([]*Vertex, len(all))
	n := 0
	for _, vertex := range all {
		vertices[n] = vertex
		n++
	}

	return vertices, all[startName]
}

// Neighbor struct hold a target vertex and it's weight
type Neighbor struct {
	Destination *Vertex
	Weight      int
}

type Vertex struct {
	Neighbors   []*Neighbor
	Name        string
	Visited     bool
	Previous    *Vertex
	Distance    int // tentative distance
	RemoveIndex int // index in priority queue
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

type Path struct {
	Names  [] string
	length int
}

func dijkstra(vertices []*Vertex, start, end *Vertex) (paths []Path) {
	// initialize
	for _, v := range vertices {
		v.Distance = math.MaxInt64
		v.Visited = false
		v.Previous = nil
		v.RemoveIndex = - 1
	}

	current := start
	current.Distance = 0

	var unvisited VertexPriorityQueue

	for {

		// update tentative distance to neighbors
		for _, neighbor := range current.Neighbors {
			dest := neighbor.Destination
			if !dest.Visited {

				d := current.Distance + neighbor.Weight
				if d < dest.Distance {
					dest.Distance = d
					dest.Previous = current

					// put dest vertex into the priority queue
					if dest.RemoveIndex < 0 {
						heap.Push(&unvisited, dest)
					} else {
						// less expensive than `Remove` + `Push`
						heap.Fix(&unvisited, dest.RemoveIndex)
					}
				}
			}
		}

		// mark current node visited, record path and distance
		current.Visited = true
		if end == nil || current == end {
			distance := current.Distance

			var p []string

			// recover path by tracing prev links
			for ; current != nil; current = current.Previous {
				p = append(p, current.Name)
			}

			// the reverse list `p`
			for i := (len(p) + 1) / 2; i > 0; i-- {
				p[i-1], p[len(p)-i] = p[len(p)-1], p[i-1]
			}
			paths = append(paths, Path{p, distance})

			// if end vertex is nil, visit all vertices and record the shortest distance.
			if end != nil {
				return paths
			}
		}

		if len(unvisited) == 0 {
			break
		}

		// visit the next vertex which has shortest distance from current vertex
		current = heap.Pop(&unvisited).(*Vertex)
	}

	return paths
}
