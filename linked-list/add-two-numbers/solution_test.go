package addtwonumbers_test

import (
	"reflect"
	"testing"

	solution "github.com/iktzdx/go-neetcode/linked-list/add-two-numbers"
)

func Test_AddTwoNumbers(t *testing.T) {
	tests := map[string]struct {
		l1   *solution.ListNode
		l2   *solution.ListNode
		want *solution.ListNode
	}{
		"same amount of digits": {
			l1: &solution.ListNode{
				Val: 2, Next: &solution.ListNode{
					Val: 4, Next: &solution.ListNode{
						Val: 3, Next: nil,
					},
				},
			},
			l2: &solution.ListNode{
				Val: 5, Next: &solution.ListNode{
					Val: 6, Next: &solution.ListNode{
						Val: 4, Next: nil,
					},
				},
			},
			want: &solution.ListNode{
				Val: 7, Next: &solution.ListNode{
					Val: 0, Next: &solution.ListNode{
						Val: 8, Next: nil,
					},
				},
			},
		},
		"sum of two zeroes": {
			l1:   &solution.ListNode{Val: 0, Next: nil},
			l2:   &solution.ListNode{Val: 0, Next: nil},
			want: &solution.ListNode{Val: 0, Next: nil},
		},
		"amount of digits increasing": {
			l1: &solution.ListNode{
				Val: 9, Next: &solution.ListNode{
					Val: 9, Next: &solution.ListNode{
						Val: 9, Next: &solution.ListNode{
							Val: 9, Next: &solution.ListNode{
								Val: 9, Next: &solution.ListNode{
									Val: 9, Next: &solution.ListNode{
										Val: 9, Next: nil,
									},
								},
							},
						},
					},
				},
			},
			l2: &solution.ListNode{
				Val: 9, Next: &solution.ListNode{
					Val: 9, Next: &solution.ListNode{
						Val: 9, Next: &solution.ListNode{
							Val: 9, Next: nil,
						},
					},
				},
			},
			want: &solution.ListNode{
				Val: 8, Next: &solution.ListNode{
					Val: 9, Next: &solution.ListNode{
						Val: 9, Next: &solution.ListNode{
							Val: 9, Next: &solution.ListNode{
								Val: 0, Next: &solution.ListNode{
									Val: 0, Next: &solution.ListNode{
										Val: 0, Next: &solution.ListNode{
											Val: 1, Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"carry creates a new node": {
			l1:   &solution.ListNode{Val: 7, Next: nil},
			l2:   &solution.ListNode{Val: 8, Next: nil},
			want: &solution.ListNode{Val: 5, Next: &solution.ListNode{Val: 1, Next: nil}},
		},
		"operands have different number lengths": {
			l1: &solution.ListNode{
				Val: 2, Next: &solution.ListNode{
					Val: 4, Next: &solution.ListNode{
						Val: 3, Next: &solution.ListNode{
							Val: 3, Next: nil,
						},
					},
				},
			},
			l2: &solution.ListNode{
				Val: 5, Next: &solution.ListNode{
					Val: 6, Next: &solution.ListNode{
						Val: 4, Next: nil,
					},
				},
			},
			want: &solution.ListNode{
				Val: 7, Next: &solution.ListNode{
					Val: 0, Next: &solution.ListNode{
						Val: 8, Next: &solution.ListNode{
							Val: 3, Next: nil,
						},
					},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.AddTwoNumbers(test.l1, test.l2)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("AddTwoNumbers(%v, %v): expected = %v, got = %v", test.l1, test.l2, test.want, got)
			}
		})
	}
}
