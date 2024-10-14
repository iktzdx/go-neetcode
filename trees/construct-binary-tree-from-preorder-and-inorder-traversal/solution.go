package constructbinarytreefrompreorderandinordertraversal

import (
	"slices"

	"github.com/iktzdx/go-neetcode/trees"
)

func BuildTree(preorder []int, inorder []int) *trees.TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &trees.TreeNode{Val: preorder[0]}
	mid := slices.Index(inorder, preorder[0])

	root.Left = BuildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = BuildTree(preorder[mid+1:], inorder[mid+1:])

	return root
}
