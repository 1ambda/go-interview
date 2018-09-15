package leetcode

func findRedundantConnection(edges [][]int) []int {

	union := NewUnion()

	answers := make([][]int, 0)

	for _, edge := range edges {
		from := edge[0]
		to := edge[1]

		if union.Connected(from, to) {
			answers = append(answers, []int{from, to})
			continue
		}

		union.Join(from, to)
	}

	// return last cyclic-edge
	if len(answers) > 0 {
		return answers[len(answers)-1]
	}

	return nil
}

type Union interface {
	// Find finds the parent of a given element
	Find(p int) int

	// Join adds a connection between two elements
	Join(p int, q int)

	// Connected checks if two elements are connected
	Connected(p int, q int) bool
}

type union struct {
	parent map[int]int
	rank   map[int]int
}

func NewUnion() Union {
	return &union{
		parent: make(map[int]int),
		rank:   make(map[int]int),
	}
}

func (u union) Connected(p int, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u union) Find(p int) int {
	root := p

	// find	the root
	for {
		if found, ok := u.parent[root]; ok {
			root = found
		} else {
			break
		}
	}

	// compress
	for p != root {
		found := u.parent[p]
		u.parent[p] = root
		p = found
	}

	return root
}

func (u *union) Join(p int, q int) {
	pRoot := u.Find(p)
	qRoot := u.Find(q)

	if pRoot == qRoot {
		return
	}

	// merge the lower-ranking set into the larger-rank set
	if u.rank[pRoot] < u.rank[qRoot] {
		u.parent[pRoot] = qRoot
	} else {
		u.parent[qRoot] = pRoot
	}

	// increase the rank of the merged component
	// if joining two components have the same rank
	if u.rank[pRoot] == u.rank[qRoot] {
		u.rank[pRoot]++
	}
}
