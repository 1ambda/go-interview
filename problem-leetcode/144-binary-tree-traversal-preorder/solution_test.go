package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreorderTraversal(t *testing.T) {

	// [1,nil,2,3]
	root1 := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	root2 := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	// [1,4,3,2]

	expected1 := []int{1, 2, 3}
	assert.Equal(t, expected1, preorderTraversal(&root1))
	assert.Equal(t, expected1, preorderTraversalIteratively(&root1))

	expected2 := []int{1, 4, 2, 3}
	assert.Equal(t, expected2, preorderTraversal(&root2))
	assert.Equal(t, expected2, preorderTraversalIteratively(&root2))
}
