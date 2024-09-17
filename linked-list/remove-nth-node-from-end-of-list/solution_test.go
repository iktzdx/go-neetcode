package removenthnodefromendoflist_test

import (
	"reflect"
	"testing"

	solution "github.com/iktzdx/go-neetcode/linked-list/remove-nth-node-from-end-of-list"
)

func Test_RemoveNthFromEnd(t *testing.T) {
	tests := map[string]struct {
		head *solution.ListNode
		n    int
		want *solution.ListNode
	}{
		"single node": {
			head: &solution.ListNode{Val: 1, Next: nil},
			n:    1,
			want: nil,
		},
		"only one node left": {
			head: &solution.ListNode{Val: 1, Next: &solution.ListNode{Val: 2, Next: nil}},
			n:    1,
			want: &solution.ListNode{Val: 1, Next: nil},
		},
		"remove 2nd from the end": {
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
			n: 2,
			want: &solution.ListNode{
				Val: 1, Next: &solution.ListNode{
					Val: 2, Next: &solution.ListNode{
						Val: 3, Next: &solution.ListNode{
							Val: 5, Next: nil,
						},
					},
				},
			},
		},
		"remove 3rd from the end": {
			head: &solution.ListNode{
				Val: 1, Next: &solution.ListNode{
					Val: 2, Next: &solution.ListNode{
						Val: 3, Next: &solution.ListNode{
							Val: 4, Next: &solution.ListNode{
								Val: 5, Next: &solution.ListNode{
									Val: 6, Next: nil,
								},
							},
						},
					},
				},
			},
			n: 3,
			want: &solution.ListNode{
				Val: 1, Next: &solution.ListNode{
					Val: 2, Next: &solution.ListNode{
						Val: 3, Next: &solution.ListNode{
							Val: 5, Next: &solution.ListNode{
								Val: 6, Next: nil,
							},
						},
					},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.RemoveNthFromEnd(test.head, test.n)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("RemoveNthFromEnd(%v, %d): expected = %v, got = %v", test.head, test.n, test.want, got)
			}
		})
	}
}
