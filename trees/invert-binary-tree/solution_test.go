package invertbinarytree_test

import (
	"reflect"
	"testing"

	helper "github.com/iktzdx/go-neetcode/trees"
	solution "github.com/iktzdx/go-neetcode/trees/invert-binary-tree"
)

func Test_InvertTree(t *testing.T) {
	tests := map[string]struct {
		root []int
		want []int
	}{
		"case #1 - only two nodes after root": {
			root: []int{2, 1, 3},
			want: []int{2, 3, 1},
		},
		"case #2 - multiple nodes after root": {
			root: []int{4, 2, 7, 1, 3, 6, 9},
			want: []int{4, 7, 2, 9, 6, 3, 1},
		},
		"case #3 - empty tree": {
			root: []int{},
			want: []int{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			in := toSolutionTreeNode(helper.BSTFromArrayInts(test.root))
			expected := toSolutionTreeNode(helper.BSTFromArrayInts(test.want))

			got := solution.InvertTree(in)

			if !reflect.DeepEqual(expected, got) {
				gotArray := helper.BSTToArrayInts(toHelperTreeNode(got))
				t.Fatalf("InvertTree(%v): expected = %v, got = %v", test.root, test.want, gotArray)
			}
		})
	}
}

func toSolutionTreeNode(in *helper.TreeNode) *solution.TreeNode {
	if in == nil {
		return nil
	}

	return &solution.TreeNode{
		Val:   in.Val,
		Left:  toSolutionTreeNode(in.Left),
		Right: toSolutionTreeNode(in.Right),
	}
}

func toHelperTreeNode(in *solution.TreeNode) *helper.TreeNode {
	if in == nil {
		return nil
	}

	return &helper.TreeNode{
		Val:   in.Val,
		Left:  toHelperTreeNode(in.Left),
		Right: toHelperTreeNode(in.Right),
	}
}
