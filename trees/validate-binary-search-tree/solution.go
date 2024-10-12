package validatebinarysearchtree

import (
	"math"

	"github.com/iktzdx/go-neetcode/trees"
)

func IsValidBST(root *trees.TreeNode) bool {
	return dfs(root, math.MinInt, math.MaxInt)
}

func dfs(node *trees.TreeNode, left, right int) bool {
	if node == nil {
		return true
	}

	if node.Val >= right || node.Val <= left {
		return false
	}

	return dfs(node.Left, left, node.Val) && dfs(node.Right, node.Val, right)
}
