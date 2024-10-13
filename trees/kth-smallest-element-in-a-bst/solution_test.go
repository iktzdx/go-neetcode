package kthsmallestelementinabst_test

import (
	"testing"

	"github.com/iktzdx/go-neetcode/trees"
	solution "github.com/iktzdx/go-neetcode/trees/kth-smallest-element-in-a-bst"
)

func Test_KthSmallest(t *testing.T) {
	tests := map[string]struct {
		root []int
		k    int
		want int
	}{
		"case #1": {
			root: []int{3, 1, 4, trees.NULL, 2},
			k:    1,
			want: 1,
		},
		"case #2": {
			root: []int{5, 3, 6, 2, 4, trees.NULL, trees.NULL, 1},
			k:    3,
			want: 3,
		},
		"case #3": {
			root: []int{4, 3, 5, 2, trees.NULL},
			k:    4,
			want: 5,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			root := trees.BSTFromArrayInts(test.root)

			got := solution.KthSmallest(root, test.k)

			if test.want != got {
				t.Fatalf("KthSmallest(%v, %d): expected = %d, got = %d", test.root, test.k, test.want, got)
			}
		})
	}
}
