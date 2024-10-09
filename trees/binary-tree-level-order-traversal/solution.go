package binarytreelevelordertraversal

import (
	"github.com/iktzdx/go-neetcode/trees"
)

func LevelOrder(root *trees.TreeNode) [][]int {
	levels := make([][]int, 0)

	if root == nil {
		return levels
	}

	q := []*trees.TreeNode{root}

	for len(q) > 0 {
		level := make([]int, 0)

		for range len(q) {
			var node *trees.TreeNode

			node, q = q[0], q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}

			level = append(level, node.Val)
		}

		levels = append(levels, level)
	}

	return levels
}
