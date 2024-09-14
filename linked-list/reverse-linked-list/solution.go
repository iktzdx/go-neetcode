package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var reversed *ListNode

	for head != nil {
		next := head.Next

		head.Next = reversed
		reversed = head
		head = next
	}

	return reversed
}
