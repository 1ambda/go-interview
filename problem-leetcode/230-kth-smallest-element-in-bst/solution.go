package leetcode

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	values []interface{}
}

func NewStack() *Stack {
	return &Stack{
		values: make([]interface{}, 0),
	}
}

func (s *Stack) Push(value interface{}) int {
	s.values = append(s.values, value)
	return len(s.values)
}

func (s *Stack) Pop() interface{} {
	value := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]

	return value
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack) Peek() interface{} {
	if len(s.values) == 0 {
		panic("Can't peek an empty stack")
	}

	return s.values[len(s.values)-1]
}

func kthSmallest(root *TreeNode, k int) int {

	stack := NewStack()

	for root != nil || !stack.IsEmpty() {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}

		root = stack.Pop().(*TreeNode)

		// traversing BST in-order gives a sorted array
		k--
		if k == 0 {
			return root.Val
		}

		root = root.Right

	}

	return math.MinInt64 // shouldn't be here
}
