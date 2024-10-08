package lowestcommonancestorofabinarysearchtree_test

import (
	"reflect"
	"testing"

	"github.com/iktzdx/go-neetcode/trees"
	solution "github.com/iktzdx/go-neetcode/trees/lowest-common-ancestor-of-a-binary-search-tree"
)

func Test_LowestCommonAncestor(t *testing.T) {
	tests := map[string]struct {
		root []int
		p, q int
		want int
	}{
		"case #1 - root node is LCA": {
			root: []int{6, 2, 8, 0, 4, 7, 9, trees.NULL, trees.NULL, 3, 5},
			p:    2,
			q:    8,
			want: 6,
		},
		"case #2 - a node is a descendant of itself": {
			root: []int{6, 2, 8, 0, 4, 7, 9, trees.NULL, trees.NULL, 3, 5},
			p:    2,
			q:    4,
			want: 2,
		},
		"case #3 - the tree consists only of p and q": {
			root: []int{2, 1},
			p:    2,
			q:    1,
			want: 2,
		},
		"case #4 - search only in the right sub-tree": {
			root: []int{6, 2, 8, 0, 4, 7, 9, trees.NULL, trees.NULL, 3, 5},
			p:    7,
			q:    9,
			want: 8,
		},
		"case #5 - search only in the left sub-tree": {
			root: []int{6, 2, 8, 0, 4, 7, 9, trees.NULL, trees.NULL, 3, 5},
			p:    3,
			q:    5,
			want: 4,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			root := trees.BSTFromArrayInts(test.root)
			p, q := trees.BSTFromArrayInts([]int{test.p}), trees.BSTFromArrayInts([]int{test.q})

			got := trees.BSTToArrayInts(solution.LowestCommonAncestor(root, p, q))[0]

			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("LowestCommonAncestor(%v, %v, %v): expected = %v, got = %v", test.root, test.p, test.q, test.want, got)
			}
		})
	}
}
