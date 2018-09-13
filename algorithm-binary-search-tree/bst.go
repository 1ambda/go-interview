package bst

import (
	"math"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) Min() int {
	if n == nil {
		return math.MinInt64
	}

	if n.left == nil {
		return n.value
	}

	return n.left.Min()
}

func (n *Node) Max() int {
	if n == nil {
		return math.MaxInt64
	}

	if n.right == nil {
		return n.value
	}

	return n.right.Max()
}

func (n *Node) Find(target int) bool {
	if n == nil {
		return false
	}

	if target < n.value {
		return n.left.Find(target)
	} else if target == n.value {
		return true
	} else {
		return n.right.Find(target)
	}
}

// Valid returns true if it's valid binary search tree.
func (n *Node) Valid() bool {
	return n.validRecursive(math.MinInt64, math.MaxInt64)
}

func (n *Node) validRecursive(min, max int) bool {
	if n == nil {
		return true
	}

	return min < n.value && n.value < max &&
		n.left.validRecursive(min, n.value) &&
		n.right.validRecursive(n.value, max)
}

// Insert the passed value and returns true.
// If value already exists or tree is valid, will return false.
func (n *Node) Insert(value int) bool {
	if n == nil {
		return false
	}

	if value == n.value {
		return false
	}

	if value < n.value {
		if n.left == nil {
			n.left = &Node{
				value: value,
			}
			return true
		}

		return n.left.Insert(value)
	} else {
		if n.right == nil {
			n.right = &Node{
				value: value,
			}
			return true
		}

		return n.right.Insert(value)
	}
}

// Delete returns true if target node is removed successfully, Otherwise returns false.
func (n *Node) Delete(value int) (bool, *Node) {
	if n == nil {
		return false, nil
	}

	if n.value == value {
		// if both children nodes are empty, return nil
		if n.left == nil && n.right == nil {
			return true, nil
		}

		// return right node if left node is empty
		if n.left == nil && n.right != nil {
			return true, n.right
		}

		// return left node if right node is empty
		if n.left != nil && n.right == nil {
			return true, n.left
		}

		// if both children nodes are not empty
		// then, use Max(n.left) as the current node
		middle := n.left.Max()
		_, n.left = n.left.Delete(middle)
		n.value = middle

		return true, n
	}

	if value < n.value {
		deleted, left := n.left.Delete(value)
		n.left = left
		return deleted, n
	} else {
		deleted, right := n.right.Delete(value)
		n.right = right
		return deleted, n
	}
}
