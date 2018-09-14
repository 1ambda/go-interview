package dijkstra

import (
	"fmt"
	"testing"
)

// Edge struct hold the bare data needed to define a graph
type Edge struct {
	vertex1, vertex2 string
	distance         float64
}

// https://rosettacode.org/wiki/Dijkstra%27s_algorithm#Go
func buildGraph(graph []Edge, directed bool, start, end string) (vertice []*Vertex, startVertex, endVertex *Vertex) {

	all := make(map[string]*Vertex)

	for _, edge := range graph {
		v1 := all[edge.vertex1]
		v2 := all[edge.vertex2]

		// add previously unseen vertices
		if v1 == nil {
			v1 = &Vertex{Name: edge.vertex1}
			all[edge.vertex1] = v1
		}
		if v2 == nil {
			v2 = &Vertex{Name: edge.vertex2}
			all[edge.vertex2] = v2
		}

		v1.Neighbors = append(v1.Neighbors, &Neighbor{Destination: v2, Weight: edge.distance})
		if !directed {
			v2.Neighbors = append(v2.Neighbors, &Neighbor{Destination: v1, Weight: edge.distance})
		}
	}

	vertice = make([]*Vertex, len(all))

	n := 0
	for _, vertex := range all {
		vertice[n] = vertex
		n++
	}

	return vertice, all[start], all[end]
}

func TestDijkstra(t *testing.T) {
	graph := []Edge{
		{"a", "b", 7.0},
		{"a", "c", 9.0},
		{"a", "f", 14.0},
		{"b", "c", 10.0},
		{"b", "d", 15.0},
		{"c", "d", 11.0},
		{"c", "f", 2.0},
		{"d", "e", 6.0},
		{"e", "f", 9.0},
	}

	directed := true
	start := "a"
	end := "e"

	vertices, startVertex, endVertex := buildGraph(graph, directed, start, end)
	findAll := false

	if directed {
		fmt.Print("Directed")
	} else {
		fmt.Print("Undirected")
	}

	fmt.Printf(" graph with %d nodes, %d edges\n", len(vertices), len(graph))
	if startVertex == nil {
		fmt.Printf("start node %q not found in graph\n", start)
		return
	}

	if findAll {
		endVertex = nil
	} else if endVertex == nil {
		fmt.Printf("end node %q not found in graph\n", end)
		return
	}

	// run Dijkstra's shortest path algorithm
	paths := dijkstra(vertices, startVertex, endVertex)
	fmt.Println("Shortest path(s):")
	for _, p := range paths {
		fmt.Println(p.Names, "length", p.length)
	}

}
