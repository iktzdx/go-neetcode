// Copied from https://github.com/hidypeople/leetcode-go/blob/main/binaryTree/binaryTree.go
package trees

var NULL = -1 << 63 //nolint:gochecknoglobals

// binary tree with int values.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Transform array into binary tree, the empty children are NULL.
func BSTFromArrayInts(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}

		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}

		i++
	}

	return root
}

// Transform binary tree to array of ints, the empty children will be NULL.
func BSTToArrayInts(root *TreeNode) []int {
	res := make([]int, 0, 1024)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		for i := range size {
			nd := queue[i]
			if nd == nil {
				res = append(res, NULL)
			} else {
				res = append(res, nd.Val)
				queue = append(queue, nd.Left, nd.Right)
			}
		}

		queue = queue[size:]
	}

	i := len(res)
	for i > 0 && res[i-1] == NULL {
		i--
	}

	return res[:i]
}
