package countgoodnodesinbinarytree

import "github.com/iktzdx/go-neetcode/trees"

type dfsFunc func(node *trees.TreeNode, maxVal int) int

func GoodNodes(root *trees.TreeNode) int {
	var dfs dfsFunc

	dfs = func(node *trees.TreeNode, maxVal int) int {
		var count int

		if node == nil {
			return count
		}

		if node.Val >= maxVal {
			count++
		}

		maxVal = max(maxVal, node.Val)

		return count + dfs(node.Left, maxVal) + dfs(node.Right, maxVal)
	}

	return dfs(root, root.Val)
}
