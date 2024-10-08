package lowestcommonancestorofabinarysearchtree

import "github.com/iktzdx/go-neetcode/trees"

func LowestCommonAncestor(root, p, q *trees.TreeNode) *trees.TreeNode {
	curr := root

	for curr != nil {
		if p.Val > curr.Val && q.Val > curr.Val {
			curr = curr.Right
		} else if p.Val < curr.Val && q.Val < curr.Val {
			curr = curr.Left
		} else {
			break
		}
	}

	return curr
}
