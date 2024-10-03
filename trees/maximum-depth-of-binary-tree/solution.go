package maximumdepthofbinarytree

import "github.com/iktzdx/go-neetcode/trees"

func MaxDepth(root *trees.TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(MaxDepth(root.Left), MaxDepth(root.Right))
}
