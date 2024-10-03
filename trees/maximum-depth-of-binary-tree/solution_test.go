package maximumdepthofbinarytree_test

import (
	"testing"

	"github.com/iktzdx/go-neetcode/trees"
	maximumdepthofbinarytree "github.com/iktzdx/go-neetcode/trees/maximum-depth-of-binary-tree"
)

func Test_MaxDepth(t *testing.T) {
	tests := map[string]struct {
		root []int
		want int
	}{
		"case #1 - empty tree": {
			root: []int{},
			want: 0,
		},
		"case #2 - single vertex": {
			root: []int{1},
			want: 1,
		},
		"case #3 - two children": {
			root: []int{1, trees.NULL, 2},
			want: 2,
		},
		"case #4 - contains sub-trees": {
			root: []int{3, 9, 20, trees.NULL, trees.NULL, 15, 7},
			want: 3,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := maximumdepthofbinarytree.MaxDepth(trees.BSTFromArrayInts(test.root))
			if test.want != got {
				t.Fatalf("MaxDepth(%v): expected = %v, got = %v", test.root, test.want, got)
			}
		})
	}
}
