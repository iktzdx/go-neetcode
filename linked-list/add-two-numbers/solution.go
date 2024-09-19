package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

const base int = 10

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	curr := dummyNode

	var carry int
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / base
		curr.Next = &ListNode{Val: sum % base}
		curr = curr.Next
	}

	return dummyNode.Next
}
