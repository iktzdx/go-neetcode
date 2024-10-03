package maximumdepthofbinarytree

import "github.com/iktzdx/go-neetcode/trees"

type value struct {
	depth int
	node  *trees.TreeNode
}

func MaxDepth(root *trees.TreeNode) int {
	var result int

	if root == nil {
		return result
	}

	s := []*value{{depth: 1, node: root}}
	result = len(s)

	for len(s) > 0 {
		node, depth := s[len(s)-1].node, s[len(s)-1].depth
		s = s[:len(s)-1]

		if node != nil {
			result = max(result, depth)
			s = append(s,
				&value{node: node.Left, depth: depth + 1},
				&value{node: node.Right, depth: depth + 1},
			)
		}
	}

	return result
}
