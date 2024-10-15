package binarytreemaximumpathsum

import "github.com/iktzdx/go-neetcode/trees"

type dfsFunc func(*trees.TreeNode) int

func MaxPathSum(root *trees.TreeNode) int {
	var (
		res int
		dfs dfsFunc
	)

	// Return the max path sum without split.
	dfs = func(node *trees.TreeNode) int {
		if node == nil {
			return 0
		}

		left := max(dfs(node.Left), 0)
		right := max(dfs(node.Right), 0)

		// Calculate max path sum with split.
		res = max(res, (node.Val + left + right))

		return node.Val + max(left, right)
	}

	dfs(root)

	return res
}
