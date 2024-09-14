package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head.Next, head, head.Next.Next = nil, ReverseList(head.Next), head

	return head
}
