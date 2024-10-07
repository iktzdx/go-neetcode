package subtreeofanothertree_test

import (
	"testing"

	"github.com/iktzdx/go-neetcode/trees"
	solution "github.com/iktzdx/go-neetcode/trees/subtree-of-another-tree"
)

func Test_IsSubtree(t *testing.T) {
	tests := map[string]struct {
		root, subRoot []int
		want          bool
	}{
		"case #1 - single node, is sub-tree": {
			root:    []int{1},
			subRoot: []int{1},
			want:    true,
		},
		"case #2 - single node, is not sub-tree": {
			root:    []int{1},
			subRoot: []int{2},
			want:    false,
		},
		"case #3 - with children, is sub-tree": {
			root:    []int{3, 4, 5, 1, 2},
			subRoot: []int{4, 1, 2},
			want:    true,
		},
		"case #4 - same nodes, different order, is not sub-tree": {
			root:    []int{3, 4, 5, 1, 2},
			subRoot: []int{4, 2, 1},
			want:    false,
		},
		"case #5 - missing children, is not sub-tree": {
			root:    []int{3, 4, 5, 1, 2, trees.NULL, trees.NULL, trees.NULL, trees.NULL, 0, trees.NULL},
			subRoot: []int{4, 1, 2},
			want:    false,
		},
		"case #6 - empty sub-tree is a sub-tree": {
			root:    []int{1, 2, 3},
			subRoot: []int{},
			want:    true,
		},
		"case #7 - sub-tree of an empty tree": {
			root:    []int{},
			subRoot: []int{1, 2, 3},
			want:    false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			root := trees.BSTFromArrayInts(test.root)
			subRoot := trees.BSTFromArrayInts(test.subRoot)

			got := solution.IsSubtree(root, subRoot)

			if test.want != got {
				t.Fatalf("IsSubtree(%v, %v): expected = %v, got = %v", test.root, test.subRoot, test.want, got)
			}
		})
	}
}
