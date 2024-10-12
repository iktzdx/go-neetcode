package validatebinarysearchtree_test

import (
	"testing"

	"github.com/iktzdx/go-neetcode/trees"
	solution "github.com/iktzdx/go-neetcode/trees/validate-binary-search-tree"
)

func Test_IsValidBST(t *testing.T) {
	tests := map[string]struct {
		root []int
		want bool
	}{
		"case #1 - single vertex is always valid": {
			root: []int{1},
			want: true,
		},
		"case #2 - only left side, valid tree": {
			root: []int{2, 1, trees.NULL},
			want: true,
		},
		"case #3 - only right side, valid tree": {
			root: []int{1, trees.NULL, 2},
			want: true,
		},
		"case #4 - only left side, invalid tree": {
			root: []int{1, 2, trees.NULL},
			want: false,
		},
		"case #5 - only right side, invalid tree": {
			root: []int{2, trees.NULL, 1},
			want: false,
		},
		"case #6 - has both sides, valid tree": {
			root: []int{2, 1, 3},
			want: true,
		},
		"case #7 - multiple sub-trees, invalid tree": {
			root: []int{5, 1, 4, trees.NULL, trees.NULL, 3, 6},
			want: false, // 5 -> 4
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			root := trees.BSTFromArrayInts(test.root)

			got := solution.IsValidBST(root)

			if test.want != got {
				t.Fatalf("IsValidBST(%v): expected = %v, got = %v", test.root, test.want, got)
			}
		})
	}
}
