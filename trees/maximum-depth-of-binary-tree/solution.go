package maximumdepthofbinarytree

import "github.com/iktzdx/go-neetcode/trees"

func MaxDepth(root *trees.TreeNode) int {
	if root == nil {
		return 0
	}

	level := 0
	queue := []*trees.TreeNode{root}

	for len(queue) > 0 {
		for range len(queue) {
			var node *trees.TreeNode
			node, queue = queue[0], queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		level++
	}

	return level
}
