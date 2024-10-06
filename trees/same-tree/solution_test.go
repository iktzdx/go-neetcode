package sametree_test

import (
	"testing"

	"github.com/iktzdx/go-neetcode/trees"
	solution "github.com/iktzdx/go-neetcode/trees/same-tree"
)

func Test_IsSameTree(t *testing.T) {
	tests := map[string]struct {
		p, q []int
		want bool
	}{
		"case #1 - empty trees": {
			p:    []int{},
			q:    []int{},
			want: true,
		},
		"case #2 - single vertex, identical": {
			p:    []int{1},
			q:    []int{1},
			want: true,
		},
		"case #3 - single vertex, not identical": {
			p:    []int{},
			q:    []int{1},
			want: false,
		},
		"case #4 - single vertex, not identical": {
			p:    []int{1},
			q:    []int{},
			want: false,
		},
		"case #5 - structurally identical": {
			p:    []int{1, 2, 3},
			q:    []int{1, 2, 3},
			want: true,
		},
		"case #6 - structurally not identical": {
			p:    []int{1, 2, trees.NULL},
			q:    []int{1, trees.NULL, 2},
			want: false,
		},
		"case #7 - identical nodes, different order": {
			p:    []int{1, 2, 1},
			q:    []int{1, 1, 2},
			want: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p, q := trees.BSTFromArrayInts(test.p), trees.BSTFromArrayInts(test.q)
			got := solution.IsSameTree(p, q)

			if test.want != got {
				t.Fatalf("IsSameTree(%v, %v): expected = %v, got = %v", test.p, test.q, test.want, got)
			}
		})
	}
}
