package binarytreemaximumpathsum_test

import (
	"testing"

	"github.com/iktzdx/go-neetcode/trees"
	solution "github.com/iktzdx/go-neetcode/trees/binary-tree-maximum-path-sum"
)

func Test_MaxPathSum(t *testing.T) {
	tests := map[string]struct {
		root []int
		want int
	}{
		"case #1": {
			root: []int{1, 2, 3},
			want: 6, // 2 -> 1 -> 3
		},
		"case #2": {
			root: []int{-10, 9, 20, trees.NULL, trees.NULL, 15, 7},
			want: 42, // 15 -> 20 -> 7
		},
		"case #3": {
			root: []int{-15, 10, 20, trees.NULL, trees.NULL, 15, 5, -5},
			want: 40,
		},
		"case #4": {
			root: []int{1, -2, 3},
			want: 4,
		},
		"case #5": {
			root: []int{-1, 2, 3},
			want: 4,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			root := trees.BSTFromArrayInts(test.root)

			got := solution.MaxPathSum(root)

			if test.want != got {
				t.Fatalf("MaxPathSum(%v): expected = %d, got = %d", test.root, test.want, got)
			}
		})
	}
}
