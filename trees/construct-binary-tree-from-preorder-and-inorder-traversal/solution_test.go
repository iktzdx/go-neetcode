package constructbinarytreefrompreorderandinordertraversal_test

import (
	"reflect"
	"testing"

	"github.com/iktzdx/go-neetcode/trees"
	solution "github.com/iktzdx/go-neetcode/trees/construct-binary-tree-from-preorder-and-inorder-traversal"
)

func Test_BuildTree(t *testing.T) {
	tests := map[string]struct {
		preorder, inorder []int
		want              []int
	}{
		"case #1": {
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			want:     []int{3, 9, 20, trees.NULL, trees.NULL, 15, 7},
		},
		"case #2": {
			preorder: []int{-1},
			inorder:  []int{-1},
			want:     []int{-1},
		},
		"case #3": {
			preorder: []int{1, 2, 3, 4},
			inorder:  []int{2, 1, 3, 4},
			want:     []int{1, 2, 3, trees.NULL, trees.NULL, trees.NULL, 4},
		},
		"case #4": {
			preorder: []int{1},
			inorder:  []int{1},
			want:     []int{1},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.BuildTree(test.preorder, test.inorder)
			actual := trees.BSTToArrayInts(got)

			if !reflect.DeepEqual(test.want, actual) {
				t.Fatalf("BuildTree(%v, %v): expected = %v, got = %v", test.preorder, test.inorder, test.want, actual)
			}
		})
	}
}
