package reorderlist_test

import (
	"reflect"
	"testing"

	solution "github.com/iktzdx/go-neetcode/linked-list/reorder-list"
)

func Test_ReorderList(t *testing.T) {
	tests := map[string]struct {
		head *solution.ListNode
		want *solution.ListNode
	}{
		"example #1": {
			head: &solution.ListNode{
				Val: 1, Next: &solution.ListNode{
					Val: 2, Next: &solution.ListNode{
						Val: 3, Next: &solution.ListNode{
							Val: 4, Next: nil,
						},
					},
				},
			},
			want: &solution.ListNode{
				Val: 1, Next: &solution.ListNode{
					Val: 4, Next: &solution.ListNode{
						Val: 2, Next: &solution.ListNode{
							Val: 3, Next: nil,
						},
					},
				},
			},
		},
		"example #2": {
			head: &solution.ListNode{
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
			want: &solution.ListNode{
				Val: 1, Next: &solution.ListNode{
					Val: 5, Next: &solution.ListNode{
						Val: 2, Next: &solution.ListNode{
							Val: 4, Next: &solution.ListNode{
								Val: 3, Next: nil,
							},
						},
					},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			solution.ReorderList(test.head)
			if !reflect.DeepEqual(test.want, test.head) {
				t.Fatalf("ReorderList(...): expected = %v, got = %v", test.want, test.head)
			}
		})
	}
}
