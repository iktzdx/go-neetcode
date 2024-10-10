package binarytreerightsideview

import "github.com/iktzdx/go-neetcode/trees"

func RightSideView(root *trees.TreeNode) []int {
	result := make([]int, 0)

	q := []*trees.TreeNode{root}

	// For each level of the tree we want
	// the right most node.
	for len(q) > 0 {
		var rightNode *trees.TreeNode

		for range len(q) {
			var node *trees.TreeNode

			node, q = q[0], q[1:]

			if node != nil {
				rightNode = node

				q = append(q, node.Left)
				q = append(q, node.Right)
			}
		}

		if rightNode != nil {
			result = append(result, rightNode.Val)
		}
	}

	return result
}
