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

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	nodes := make([]int, 0)

	left := inorderTraversal(root.Left)
	for _, v := range left {
		nodes = append(nodes, v)
	}

	nodes = append(nodes, root.Val)

	right := inorderTraversal(root.Right)
	for _, v := range right {
		nodes = append(nodes, v)
	}

	return nodes
}

