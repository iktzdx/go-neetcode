package validatebinarysearchtree

import (
	"math"

	"github.com/iktzdx/go-neetcode/trees"
)

type dfsFunc func(*trees.TreeNode) bool

func IsValidBST(root *trees.TreeNode) bool {
	var dfs dfsFunc

	prev := math.MinInt

	dfs = func(node *trees.TreeNode) bool {
		if node == nil {
			return true
		}

		isValid := dfs(node.Left)
		if !isValid || node.Val <= prev {
			return false
		}

		prev = node.Val

		return dfs(node.Right)
	}

	return dfs(root)
}
