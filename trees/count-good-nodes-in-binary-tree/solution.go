package countgoodnodesinbinarytree

import "github.com/iktzdx/go-neetcode/trees"

func GoodNodes(root *trees.TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(node *trees.TreeNode, maxVal int) int {
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
