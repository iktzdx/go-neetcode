package diameterofbinarytree_test

import (
	"testing"

	"github.com/iktzdx/go-neetcode/trees"
	diameterofbinarytree "github.com/iktzdx/go-neetcode/trees/diameter-of-binary-tree"
)

func Test_DiameterOfBinaryTree(t *testing.T) {
	tests := map[string]struct {
		root []int
		want int
	}{
		"case #1 - an empty tree": {
			root: []int{},
			want: 0,
		},
		"case #2 - a single vertex": {
			root: []int{1},
			want: 0,
		},
		"case #3 - only one leaf": {
			root: []int{1, 2},
			want: 1,
		},
		"case #4 - contains sub-tree": {
			root: []int{1, 2, 3, 4, 5},
			want: 3, // 3 is the length of the path [4,2,1,3] or [5,2,1,3]
		},
		"case #5 - the longest path does not pass through the root": {
			root: []int{1, trees.NULL, 3, 4, 5, 6, 7, 8, 9},
			want: 4, // 4 is the length of the path [6,4,3,5,9] or [7,4,3,5,8]
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			root := trees.BSTFromArrayInts(test.root)
			got := diameterofbinarytree.DiameterOfBinaryTree(root)

			if test.want != got {
				t.Fatalf("DiameterOfBinaryTree(%v): expected = %d, got = %d", root, test.want, got)
			}
		})
	}
}
