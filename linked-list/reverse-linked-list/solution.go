package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// https://go.dev/ref/spec#Assignment_statements
	// The assignment proceeds in two phases:
	//   - First, the operands on the left and the expressions on the right are all evaluated in the usual order.
	//   - Second, the assignments are carried out in left-to-right order.
	head.Next, head, head.Next.Next = nil, ReverseList(head.Next), head

	return head
}
