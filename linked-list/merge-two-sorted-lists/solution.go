package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = MergeTwoLists(list1.Next, list2)
		return list1
	}

	list2.Next = MergeTwoLists(list1, list2.Next)
	return list2
}
