package mergetwosortedlists_test

import (
	"reflect"
	"testing"

	solution "github.com/iktzdx/go-neetcode/linked-list/merge-two-sorted-lists"
)

func Test_MergeTwoLists(t *testing.T) {
	tests := map[string]struct {
		list1 *solution.ListNode
		list2 *solution.ListNode
		want  *solution.ListNode
	}{
		"empty lists": {
			list1: nil,
			list2: nil,
			want:  nil,
		},
		"one of the lists has a single element": {
			list1: nil,
			list2: &solution.ListNode{Val: 1, Next: nil},
			want:  &solution.ListNode{Val: 1, Next: nil},
		},
		"both lists have a single element": {
			list1: &solution.ListNode{Val: 2, Next: nil},
			list2: &solution.ListNode{Val: 1, Next: nil},
			want:  &solution.ListNode{Val: 1, Next: &solution.ListNode{Val: 2, Next: nil}},
		},
		"there are duplicates in both lists": {
			list1: &solution.ListNode{Val: 1, Next: &solution.ListNode{Val: 2, Next: &solution.ListNode{Val: 4, Next: nil}}},
			list2: &solution.ListNode{Val: 1, Next: &solution.ListNode{Val: 3, Next: &solution.ListNode{Val: 4, Next: nil}}},
			want: &solution.ListNode{
				Val: 1, Next: &solution.ListNode{Val: 1, Next: &solution.ListNode{Val: 2, Next: &solution.ListNode{
					Val: 3, Next: &solution.ListNode{Val: 4, Next: &solution.ListNode{Val: 4, Next: nil}},
				}}},
			},
		},
		"contain negative values": {
			list1: &solution.ListNode{Val: -10, Next: &solution.ListNode{Val: -3, Next: &solution.ListNode{Val: 4, Next: nil}}},
			list2: &solution.ListNode{Val: -1, Next: &solution.ListNode{Val: 1, Next: &solution.ListNode{Val: 5, Next: nil}}},
			want: &solution.ListNode{
				Val: -10, Next: &solution.ListNode{Val: -3, Next: &solution.ListNode{Val: -1, Next: &solution.ListNode{
					Val: 1, Next: &solution.ListNode{Val: 4, Next: &solution.ListNode{Val: 5, Next: nil}},
				}}},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.MergeTwoLists(test.list1, test.list2)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("MergeTwoLists(%v, %v): expected = %v, got = %v", test.list1, test.list2, test.want, got)
			}
		})
	}
}
