package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(1)
	}

	// split into multiple lines
	raw := string(bytes)
	commands := strings.Split(raw, "\n")
	if len(commands) < 1 {
		// invalid STDIN
		os.Exit(1)
	}

	// parse total node number
	totalNodeNumber, err := parseNodeNumber(commands[0])
	if err != nil {
		os.Exit(1)
	}

	// parse edges
	edges, err := parseEdges(commands[1:])
	if err != nil {
		os.Exit(1)
	}

	result := hasCycle(totalNodeNumber, edges)
	fmt.Println(result)
}

const CmdSpace = " "

func hasCycle(totalNodeNumber int, edges []*edge) bool {
	union := NewUnion()

	for _, edge := range edges {
		from := edge.source
		to := edge.destination

		if union.Connected(from, to) {
			return true
		}

		union.Join(from, to)
	}

	return false
}

func parseNodeNumber(raw string) (int, error) {
	values, err := parseIntegers(raw)
	if err != nil {
		return 0, err
	}

	if len(values) != 1 {
		return 0, errors.New("Failed to parse node number")
	}

	return values[0], nil
}

func parseEdges(lines []string) ([]*edge, error) {

	edges := make([]*edge, 0)

	for i := range lines {
		line := lines[i]
		values, err := parseIntegers(line)

		if err != nil {
			return nil, err
		}

		if len(values) != 2 {
			return nil, errors.New("Edge has more than 2 nodes")
		}

		from := values[0]
		to := values[1]
		edges = append(edges, &edge{source: from, destination: to})
	}

	return edges, nil
}

func parseIntegers(raw string) ([]int, error) {
	// trim string to make sure about the input (STDIN)
	raw = strings.TrimSpace(raw)

	splited := strings.Split(raw, CmdSpace)
	integers := make([]int, 0)

	for i := range splited {
		stringified := splited[i]
		x, err := strconv.Atoi(stringified)
		if err != nil {
			return nil, err
		}

		integers = append(integers, x)
	}

	return integers, nil
}

// Edge represents the DIRECTED edge in a graph (connecting 2 nodes)
type edge struct {
	source      int
	destination int
}

// Use union-find structure to check whether graph has cycle or not.
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
