package balancedbinarytree

import (
	"github.com/iktzdx/go-neetcode/trees"
)

func IsBalanced(root *trees.TreeNode) bool {
	balance, _ := dfs(root)

	return balance
}

func dfs(root *trees.TreeNode) (bool, int) {
	var height int

	if root == nil {
		return true, height
	}

	balanceLeft, heightLeft := dfs(root.Left)
	balanceRight, heightRight := dfs(root.Right)
	balance := (abs(heightLeft-heightRight) <= 1) && balanceLeft && balanceRight

	return balance, 1 + max(heightLeft, heightRight)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
