package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	result := make([][]int, 0)

	if root == nil {
		return result
	}

	parentQueue := make([]*TreeNode, 0)
	parentQueue = append(parentQueue, root)

	for len(parentQueue) > 0 {
		childrenQueue := make([]*TreeNode, 0)
		currentLevel := make([]int, 0)

		// iterate current level and put children
		for len(parentQueue) > 0 {
			// dequeue
			currentNode := parentQueue[0]
			parentQueue = parentQueue[1:]
			currentLevel = append(currentLevel, currentNode.Val)

			if currentNode.Left != nil {
				childrenQueue = append(childrenQueue, currentNode.Left)
			}
			if currentNode.Right != nil {
				childrenQueue = append(childrenQueue, currentNode.Right)
			}
		}

		result = append(result, currentLevel)
		parentQueue = childrenQueue
	}

	return result
}
