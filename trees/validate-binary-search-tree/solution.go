package validatebinarysearchtree

import (
	"math"

	"github.com/iktzdx/go-neetcode/trees"
)

func IsValidBST(root *trees.TreeNode) bool {
	prev := math.MinInt

	return dfs(root, &prev)
}

func dfs(node *trees.TreeNode, prev *int) bool {
	if node == nil {
		return true
	}

	isValid := dfs(node.Left, prev)
	if !isValid || node.Val <= *prev {
		return false
	}

	*prev = node.Val

	return dfs(node.Right, prev)
}
