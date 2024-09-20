package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	visited := make(map[*ListNode]bool, 0)

	for head != nil {
		if visited[head] {
			return true
		}

		visited[head] = true
		head = head.Next
	}

	return false
}
