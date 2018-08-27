package leetcode

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBSTRecursive(node *TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}

	return (min < node.Val && node.Val < max) &&
		isValidBSTRecursive(node.Left, min, node.Val) &&
		isValidBSTRecursive(node.Right, node.Val, max)
}

// recursion version
func isValidBST(root *TreeNode) bool {
	if nil == root {
		return true
	}

	return isValidBSTRecursive(root, math.MinInt64, math.MaxInt64)
}
