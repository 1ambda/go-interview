package dijkstra

import (
	"container/heap"
	"math"
)

type Path struct {
	Names  [] string
	length float64
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
				p= append(p, current.Name)
			}

			// the reverse list `p`
			for i := (len(p) + 1) / 2; i > 0; i-- {
				p[i-1], p[len(p) - i] = p[len(p)-1], p[i-1]
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
