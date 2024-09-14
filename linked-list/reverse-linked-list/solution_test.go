package reverselinkedlist_test

import (
	"reflect"
	"testing"

	solution "github.com/iktzdx/go-neetcode/linked-list/reverse-linked-list"
)

func Test_ReverseList(t *testing.T) {
	tests := map[string]struct {
		head *solution.ListNode
		want *solution.ListNode
	}{
		"long list": {
			head: &solution.ListNode{ // 1,2,3,4,5
				Val: 1, Next: &solution.ListNode{
					Val: 2, Next: &solution.ListNode{
						Val: 3, Next: &solution.ListNode{
							Val: 4, Next: &solution.ListNode{
								Val: 5, Next: nil,
							},
						},
					},
				},
			},
			want: &solution.ListNode{ // 5,4,3,2,1
				Val: 5, Next: &solution.ListNode{
					Val: 4, Next: &solution.ListNode{
						Val: 3, Next: &solution.ListNode{
							Val: 2, Next: &solution.ListNode{
								Val: 1, Next: nil,
							},
						},
					},
				},
			},
		},
		"short list": {
			head: &solution.ListNode{ // 2,1
				Val: 2, Next: &solution.ListNode{
					Val: 1, Next: nil,
				},
			},
			want: &solution.ListNode{ // 1,2
				Val: 1, Next: &solution.ListNode{
					Val: 2, Next: nil,
				},
			},
		},
		"empty list": {
			head: &solution.ListNode{},
			want: &solution.ListNode{},
		},
		"head is nil": {
			head: nil,
			want: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.ReverseList(test.head)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("ReverseList(%v): expected = %v, got = %v", test.head, test.want, got)
			}
		})
	}
}
