package graphvalidtree_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/graphs/graph-valid-tree"
)

func Test_ValidTree(t *testing.T) {
	tests := map[string]struct {
		n     int
		edges [][]int
		want  bool
	}{
		"case #1": {
			n: 5,
			edges: [][]int{
				{0, 1}, {0, 2}, {0, 3}, {1, 4},
			},
			want: true,
		},
		"case #2": {
			n: 5,
			edges: [][]int{
				{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4},
			},
			want: false,
		},
		"case #3": {
			n:     0,
			edges: [][]int{},
			want:  true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.ValidTree(test.n, test.edges)

			if test.want != got {
				t.Fatalf("ValidTree(%d, %v): expected = %v, got = %v", test.n, test.edges, test.want, got)
			}
		})
	}
}
