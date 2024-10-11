package countgoodnodesinbinarytree_test

import (
	"testing"

	"github.com/iktzdx/go-neetcode/trees"
	solution "github.com/iktzdx/go-neetcode/trees/count-good-nodes-in-binary-tree"
)

func Test_GoodNodes(t *testing.T) {
	tests := map[string]struct {
		root []int
		want int
	}{
		"case #1 - root is always a good node": {
			root: []int{1},
			want: 1,
		},
		"case #2 - all nodes are good": {
			root: []int{1, 2, 3},
			want: 3,
		},
		"case #3 - only one bad node": {
			root: []int{3, 4, trees.NULL, 5, 2},
			want: 3, // 3 -> 4 -> 5
		},
		"case #4 - equal values are good": {
			root: []int{3, 3, trees.NULL, 4, 2},
			want: 3, // 3 -> 3 -> 4
		},
		"case #5 - there are good and bad nodes": {
			root: []int{3, 1, 4, 3, trees.NULL, 1, 5},
			want: 4, // 3 -> 4 -> 5 (3); 3 -> 1 -> 3 (1)
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			root := trees.BSTFromArrayInts(test.root)

			got := solution.GoodNodes(root)

			if test.want != got {
				t.Fatalf("GoodNodes(%v): expected = %d, got = %d", test.root, test.want, got)
			}
		})
	}
}
