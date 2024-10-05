package balancedbinarytree_test

import (
	"testing"

	"github.com/iktzdx/go-neetcode/trees"
	balancedbinarytree "github.com/iktzdx/go-neetcode/trees/balanced-binary-tree"
)

func Test_IsBalanced(t *testing.T) {
	tests := map[string]struct {
		root []int
		want bool
	}{
		"case #1 - empty tree": {
			root: []int{},
			want: true,
		},
		"case #2 - single vertex": {
			root: []int{1},
			want: true,
		},
		"case #3 - balanced": {
			root: []int{3, 9, 20, trees.NULL, trees.NULL, 15, 7},
			want: true,
		},
		"case #4 - not balanced": {
			root: []int{1, 2, 2, 3, 3, trees.NULL, trees.NULL, 4, 4},
			want: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			root := trees.BSTFromArrayInts(test.root)
			got := balancedbinarytree.IsBalanced(root)

			if test.want != got {
				t.Fatalf("IsBalanced(%v): expected = %v, got = %v", test.root, test.want, got)
			}
		})
	}
}
