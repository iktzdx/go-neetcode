package copylistwithrandompointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CopyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	oldToCopy := make(map[*Node]*Node, 0)

	// First pass:
	//  - create copies of each node,
	//  - map the original node to the new copy.
	curr := head
	for curr != nil {
		oldToCopy[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}

	// Second pass:
	//  - leverage the map to link every old node to the new node.
	curr = head
	for curr != nil {
		oldToCopy[curr].Next = oldToCopy[curr.Next]
		oldToCopy[curr].Random = oldToCopy[curr.Random]
		curr = curr.Next
	}

	return oldToCopy[head]
}
