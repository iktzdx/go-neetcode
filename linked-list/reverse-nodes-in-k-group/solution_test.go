package reversenodesinkgroup_test

import (
	"reflect"
	"testing"

	solution "github.com/iktzdx/go-neetcode/linked-list/reverse-nodes-in-k-group"
)

func Test_ReverseKGroup(t *testing.T) {
	tests := map[string]struct {
		head *solution.ListNode
		k    int
		want *solution.ListNode
	}{
		"k is equal to the length of the list": {
			head: &solution.ListNode{
				Val: 1, Next: &solution.ListNode{
					Val: 2, Next: &solution.ListNode{
						Val: 3, Next: nil,
					},
				},
			},
			k: 3,
			want: &solution.ListNode{
				Val: 3, Next: &solution.ListNode{
					Val: 2, Next: &solution.ListNode{
						Val: 1, Next: nil,
					},
				},
			},
		},
		"reverse the entire list one element at a time": {
			head: &solution.ListNode{
				Val: 1, Next: &solution.ListNode{
					Val: 2, Next: &solution.ListNode{
						Val: 3, Next: nil,
					},
				},
			},
			k: 1,
			want: &solution.ListNode{
				Val: 1, Next: &solution.ListNode{
					Val: 2, Next: &solution.ListNode{
						Val: 3, Next: nil,
					},
				},
			},
		},
		"the number of nodes is a multiple of k": {
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
			k: 2,
			want: &solution.ListNode{
				Val: 2, Next: &solution.ListNode{
					Val: 1, Next: &solution.ListNode{
						Val: 4, Next: &solution.ListNode{
							Val: 3, Next: &solution.ListNode{
								Val: 6, Next: &solution.ListNode{
									Val: 5, Next: nil,
								},
							},
						},
					},
				},
			},
		},
		"the number of nodes is not a multiple of k": {
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
			k: 3,
			want: &solution.ListNode{
				Val: 3, Next: &solution.ListNode{
					Val: 2, Next: &solution.ListNode{
						Val: 1, Next: &solution.ListNode{
							Val: 4, Next: &solution.ListNode{
								Val: 5, Next: nil,
							},
						},
					},
				},
			},
		},
		"the list consists of a single element": {
			head: &solution.ListNode{Val: 5, Next: nil},
			k:    1,
			want: &solution.ListNode{Val: 5, Next: nil},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.ReverseKGroup(test.head, test.k)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("ReverseKGroup(%v, %d): expected = %v, got = %v", test.head, test.k, test.want, got)
			}
		})
	}
}
