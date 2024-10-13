package kthsmallestelementinabst

import "github.com/iktzdx/go-neetcode/trees"

func KthSmallest(root *trees.TreeNode, k int) int {
	var processed int // Amount of processed nodes.

	stack := make([]*trees.TreeNode, 0)

	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil { // Go left as far as we can.
			stack = append(stack, curr)
			curr = curr.Left
		}

		node := stack[len(stack)-1]
		processed++

		if k == processed {
			return node.Val
		}

		stack = stack[:len(stack)-1]
		curr = node.Right // There are no more nodes on the left side.
	}

	return -1
}
