package clonegraph_test

import (
	"testing"

	"github.com/iktzdx/go-neetcode/graphs"
	solution "github.com/iktzdx/go-neetcode/graphs/clone-graph"
)

func Test_CloneGraph(t *testing.T) {
	tests := map[string]struct {
		node [][]int
		want [][]int
	}{
		"case #1": {
			node: nil,
			want: nil,
		},
		"case #2": {
			node: [][]int{{}},
			want: [][]int{},
		},
		"case #3": {
			node: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
			want: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			input := graphs.GraphFromNeighborIdxArray(test.node)
			got := solution.CloneGraph(input)
			output := graphs.GraphToNeighborIdxArray(got)

			if !graphs.GraphArrayEqual(test.want, output) {
				t.Fatalf("CloneGraph(%v): expected = %v, got = %v", test.node, test.want, output)
			}
		})
	}
}
