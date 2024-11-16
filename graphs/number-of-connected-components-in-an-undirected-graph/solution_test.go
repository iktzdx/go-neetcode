package numberofconnectedcomponentsinanundirectedgraph_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/graphs/number-of-connected-components-in-an-undirected-graph"
)

func Test_CountComponents(t *testing.T) {
	tests := map[string]struct {
		n     int
		edges [][]int
		want  int
	}{
		"case #1": {
			n:     3,
			edges: [][]int{{0, 1}, {0, 2}},
			want:  1,
		},
		"case #2": {
			n:     6,
			edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {4, 5}},
			want:  2,
		},
		"case #3": {
			n:     5,
			edges: [][]int{{0, 1}, {1, 2}, {3, 4}, {0, 2}},
			want:  2,
		},
		"case #4": {
			n:     6,
			edges: [][]int{{0, 1}, {2, 3}, {3, 4}, {3, 5}, {0, 2}},
			want:  1,
		},
		"case #5": {
			n:     5,
			edges: [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}},
			want:  2,
		},
		"case #6": {
			n:     7,
			edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {4, 5}, {5, 6}},
			want:  2,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.CountComponents(test.n, test.edges)

			if test.want != got {
				t.Fatalf("CountComponents(%d, %v): expected = %d, got = %d", test.n, test.edges, test.want, got)
			}
		})
	}
}
