package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Next: head}
	left, right := dummyNode, head

	// right := head + n
	for n > 0 && right != nil {
		right = right.Next
		n--
	}

	// shift both pointers
	for right != nil {
		left = left.Next
		right = right.Next
	}

	// remove nth node
	left.Next = left.Next.Next

	return dummyNode.Next
}
