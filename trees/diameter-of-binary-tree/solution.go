package diameterofbinarytree

import "github.com/iktzdx/go-neetcode/trees"

func DiameterOfBinaryTree(root *trees.TreeNode) int {
	tree := NewBST()

	tree.DFS(root)

	return tree.GetDiameter()
}

type bst struct {
	diameter int
}

func NewBST() *bst {
	return &bst{diameter: 0}
}

func (t *bst) DFS(curr *trees.TreeNode) int {
	if curr == nil {
		return 0
	}

	left := t.DFS(curr.Left)
	right := t.DFS(curr.Right)

	t.diameter = max(t.diameter, left+right)

	return 1 + max(left, right)
}

func (t *bst) GetDiameter() int {
	return t.diameter
}
