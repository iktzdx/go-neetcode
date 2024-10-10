package binarytreerightsideview_test

import (
	"reflect"
	"testing"

	"github.com/iktzdx/go-neetcode/trees"
	solution "github.com/iktzdx/go-neetcode/trees/binary-tree-right-side-view"
)

func Test_RightSideView(t *testing.T) {
	tests := map[string]struct {
		root []int
		want []int
	}{
		"case #1 - empty tree": {
			root: []int{},
			want: []int{},
		},
		"case #2 - single vertex": {
			root: []int{1},
			want: []int{1},
		},
		"case #3 - single child on the right": {
			root: []int{1, trees.NULL, 3},
			want: []int{1, 3},
		},
		"case #4 - single child on the left": {
			root: []int{1, 2, trees.NULL},
			want: []int{1, 2},
		},
		"case #5 - multiple nodes": {
			root: []int{1, 2, 3, trees.NULL, 5, trees.NULL, 4},
			want: []int{1, 3, 4},
		},
		"case #6 - multiple nodes": {
			root: []int{1, 2, 3, 4},
			want: []int{1, 3, 4},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			root := trees.BSTFromArrayInts(test.root)

			got := solution.RightSideView(root)

			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("RightSideView(%v): expected = %v, got = %v", test.root, test.want, got)
			}
		})
	}
}
