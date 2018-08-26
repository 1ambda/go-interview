package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInorderTraversal(t *testing.T) {

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

	//// [1,4,3,2]
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

	expected1 := []int{1, 3, 2}
	assert.Equal(t, expected1, inorderTraversal(&root1))

	expected2 := []int{1, 4, 3, 2}
	assert.Equal(t, expected2, inorderTraversal(&root2))
}
