package reorderlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReorderList(head *ListNode) {
	// Find middle
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse second part
	var prev, second *ListNode
	second, slow.Next = slow.Next, nil

	for second != nil {
		second.Next, prev, second = prev, second, second.Next
	}

	second = prev

	// Merge
	first := head
	for second != nil {
		first.Next, second.Next, first, second = second, first.Next, first.Next, second.Next
	}
}
