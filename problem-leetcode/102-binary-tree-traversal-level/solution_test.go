package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {

	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	expected1 := [][]int{
		[]int{3},
		[]int{9, 20},
		[]int{15, 7},
	}
	assert.Equal(t, expected1, levelOrder(root1))

	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
		},
	}
	expected2 := [][]int{
		[]int{1},
		[]int{2},
		[]int{3},
		[]int{4},
		[]int{5},
	}
	assert.Equal(t, expected2, levelOrder(root2))
}
