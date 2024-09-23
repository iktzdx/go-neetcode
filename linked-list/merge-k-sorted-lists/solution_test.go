package mergeksortedlists_test

import (
	"reflect"
	"testing"

	solution "github.com/iktzdx/go-neetcode/linked-list/merge-k-sorted-lists"
)

func Test_MergeKLists(t *testing.T) {
	tests := map[string]struct {
		lists []*solution.ListNode
		want  *solution.ListNode
	}{
		"each list contains a single element": {
			lists: []*solution.ListNode{
				{Val: 5, Next: nil},
				{Val: 7, Next: nil},
				{Val: 3, Next: nil},
				{Val: 8, Next: nil},
			},
			want: &solution.ListNode{
				Val: 3, Next: &solution.ListNode{
					Val: 5, Next: &solution.ListNode{
						Val: 7, Next: &solution.ListNode{
							Val: 8, Next: nil,
						},
					},
				},
			},
		},
		"contains duplicates": {
			lists: []*solution.ListNode{
				{
					Val: 1, Next: &solution.ListNode{
						Val: 4, Next: &solution.ListNode{
							Val: 5, Next: nil,
						},
					},
				},
				{
					Val: 1, Next: &solution.ListNode{
						Val: 3, Next: &solution.ListNode{
							Val: 4, Next: nil,
						},
					},
				},
				{
					Val: 2, Next: &solution.ListNode{
						Val: 6, Next: nil,
					},
				},
			},
			want: &solution.ListNode{ //[1,1,2,3,4,4,5,6]
				Val: 1, Next: &solution.ListNode{
					Val: 1, Next: &solution.ListNode{
						Val: 2, Next: &solution.ListNode{
							Val: 3, Next: &solution.ListNode{
								Val: 4, Next: &solution.ListNode{
									Val: 4, Next: &solution.ListNode{
										Val: 5, Next: &solution.ListNode{
											Val: 6, Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"empty input": {
			lists: nil,
			want:  nil,
		},
		"contains empty list": {
			lists: []*solution.ListNode{},
			want:  nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.MergeKLists(test.lists)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("MergeKLists(%v): expected = %v, got = %v", test.lists, test.want, got)
			}
		})
	}
}
