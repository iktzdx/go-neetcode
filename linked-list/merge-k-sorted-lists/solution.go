package mergeksortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		var mergedLists []*ListNode

		for i := 0; i < len(lists); i += 2 {
			var l1, l2 *ListNode

			l1 = lists[i]

			if (i + 1) < len(lists) {
				l2 = lists[i+1]
			}

			mergedLists = append(mergedLists, mergeLists(l1, l2))
		}

		lists = mergedLists
	}

	return lists[0]
}

func mergeLists(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}

		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else if list2 != nil {
		tail.Next = list2
	}

	return dummy.Next
}
