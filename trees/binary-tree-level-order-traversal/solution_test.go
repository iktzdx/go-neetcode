package binarytreelevelordertraversal_test

import (
	"reflect"
	"testing"

	"github.com/iktzdx/go-neetcode/trees"
	binarytreelevelordertraversal "github.com/iktzdx/go-neetcode/trees/binary-tree-level-order-traversal"
)

func Test_LevelOrder(t *testing.T) {
	tests := map[string]struct {
		root []int
		want [][]int
	}{
		"case #1 - zero levels": {
			root: []int{},
			want: [][]int{},
		},
		"case #2 - only one level": {
			root: []int{1},
			want: [][]int{{1}},
		},
		"case #3 - multiple levels": {
			root: []int{3, 9, 20, trees.NULL, trees.NULL, 15, 7},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			root := trees.BSTFromArrayInts(test.root)

			got := binarytreelevelordertraversal.LevelOrder(root)

			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("LevelOrder(%v): expected = %v, got = %v", test.root, test.want, got)
			}
		})
	}
}
