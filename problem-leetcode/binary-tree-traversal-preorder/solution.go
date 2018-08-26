package leetcode

/**

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,2,3]

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	nodes := make([]int, 0)
	nodes = append(nodes, root.Val)
	left := preorderTraversal(root.Left)
	for _, v := range left {
		nodes = append(nodes, v)
	}

	right := preorderTraversal(root.Right)
	for _, v := range right {
		nodes = append(nodes, v)
	}

	return nodes
}

// to DFS iteratively, we need stack.
func preorderTraversalIteratively(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)
	result := make([]int, 0)

	stack = append(stack, root)

	for len(stack) > 0 {
		// pop
		node := stack[0]
		stack = stack[1:]

		// pre-order
		result = append(result, node.Val)

		right := node.Right
		if right != nil {
			// push
			stack = append([]*TreeNode{right}, stack...)
		}

		left := node.Left
		if left != nil {
			// push
			stack = append([]*TreeNode{left}, stack...)
		}
	}

	return result
}
