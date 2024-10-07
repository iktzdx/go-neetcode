package subtreeofanothertree

import "github.com/iktzdx/go-neetcode/trees"

func IsSubtree(root *trees.TreeNode, subRoot *trees.TreeNode) bool {
	if subRoot == nil {
		return true
	}

	// If the previous condition does not execute, then
	// sub-tree is not empty.
	if root == nil {
		return false
	}

	if isSameTree(root, subRoot) {
		return true
	}

	return IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}

func isSameTree(p, q *trees.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p == nil || q == nil) || (p.Val != q.Val) {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
