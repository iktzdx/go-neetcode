package binarytreemaximumpathsum

import (
	"math"

	"github.com/iktzdx/go-neetcode/trees"
)

func MaxPathSum(root *trees.TreeNode) int {
	result := math.MinInt
	dfs(root, &result)

	return result
}

func dfs(node *trees.TreeNode, result *int) int {
	if node == nil {
		return 0
	}

	left := dfs(node.Left, result)
	right := dfs(node.Right, result)

	// Calculate max path sum with split.
	*result = max(*result, (node.Val + left + right))

	return max(0, max(left+node.Val, right+node.Val))
}
