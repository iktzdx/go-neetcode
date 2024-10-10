package binarytreerightsideview

import (
	"github.com/iktzdx/go-neetcode/trees"
)

type dfsFunc func(root *trees.TreeNode, lvl int)

func RightSideView(root *trees.TreeNode) []int {
	var dfs dfsFunc

	result := make([]int, 0)

	dfs = func(root *trees.TreeNode, lvl int) {
		if root == nil {
			return
		}

		if lvl > len(result) {
			result = append(result, root.Val)
		}

		lvl++

		dfs(root.Right, lvl)
		dfs(root.Left, lvl)
	}

	dfs(root, 1)

	return result
}
