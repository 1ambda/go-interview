package leetcode

type stack []int

func (s *stack) Push(v int) {
	*s = append(*s, v)
}

func (s *stack) Pop() int {
	if len(*s) == 0 {
		panic("Can't pop an empty stack")
	}

	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func NewStack() stack {
	s := make([]int, 0)
	return s
}

// DFS version
func checkCycle(edges [][]int) []int {
	if len(edges) == 0 {
		return nil
	}

	// map[from]to
	graph := make(map[int][]int)

	for _, edge := range edges {
		from := edge[0]
		to := edge[1]

		if graph[from] == nil {
			graph[from] = make([]int, 0)
		}

		graph[from] = append(graph[from], to)
	}

	visited := make(map[int]bool)
	unvisited := NewStack()

	start := edges[0][0]
	unvisited.Push(start)

	answer := make([]int, 0)

	for {
		current := unvisited.Pop()
		visited[current] = true

		neighbors := graph[current]
		for _, dest := range neighbors {
			// push dest into the stack
			if _, ok := visited[dest]; !ok {
				unvisited.Push(dest)
			}

			// check dest has back edge toward current node
			if destNeighbors, ok := graph[dest]; ok {
				for _, destNeighbor := range destNeighbors {
					if destNeighbor == current {
						answer = append(answer, current)
						answer = append(answer, dest)
						return answer
					}
				}
			}
		}


		if unvisited.isEmpty() {
			break
		}
	}

	return answer
}
