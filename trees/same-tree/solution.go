package sametree

import "github.com/iktzdx/go-neetcode/trees"

func IsSameTree(p *trees.TreeNode, q *trees.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	// If the previous condition does not execute, then
	// one of the nodes still can be nil, but not both at the same time.
	if (p == nil || q == nil) || (p.Val != q.Val) {
		return false
	}

	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
