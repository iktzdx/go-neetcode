package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	reversed := head
	if head.Next != nil {
		reversed = ReverseList(head.Next)
		head.Next.Next = head
	}

	head.Next = nil

	return reversed
}
