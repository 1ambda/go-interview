package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	root1 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	assert.True(t, isValidBST(root1))

	root2 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	assert.False(t, isValidBST(root2))

	//    10
	//   5  15
	//     6 20
	root3 := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}
	assert.False(t, isValidBST(root3))

	root4 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	assert.True(t, isValidBST(root4))
}
