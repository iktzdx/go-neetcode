package reversenodesinkgroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	newHead := &ListNode{Next: head}
	groupPrev := newHead

	for {
		kth := getKth(groupPrev, k)

		if kth == nil {
			break
		}

		groupNext := kth.Next
		prev, curr := kth.Next, groupPrev.Next

		for curr != groupNext {
			curr.Next, prev, curr = prev, curr, curr.Next
		}

		groupPrev.Next, groupPrev = kth, groupPrev.Next
	}

	return newHead.Next
}

func getKth(curr *ListNode, k int) *ListNode {
	for curr != nil && k > 0 {
		curr = curr.Next
		k--
	}

	return curr
}
