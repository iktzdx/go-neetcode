package reorderlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReorderList(head *ListNode) {
	mid := findMidNode(head)
	newHead := mid.Next
	mid.Next = nil

	newHead = reverseLinkedList(newHead)

	c1, c2 := head, newHead
	var f1, f2 *ListNode

	for c1 != nil && c2 != nil {
		f1 = c1.Next
		f2 = c2.Next

		c1.Next = c2
		c2.Next = f1

		c1 = f1
		c2 = f2
	}
}

func findMidNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverseLinkedList(head *ListNode) *ListNode {
	var prev, frw *ListNode
	curr := head

	for curr != nil {
		frw = curr.Next
		curr.Next = prev
		prev = curr
		curr = frw
	}

	return prev
}
