package pacificatlanticwaterflow_test

import (
	"reflect"
	"testing"

	pacificatlanticwaterflow "github.com/iktzdx/go-neetcode/graphs/pacific-atlantic-water-flow"
)

func Test_PacificAtlantic(t *testing.T) {
	tests := map[string]struct {
		heights [][]int
		want    [][]int
	}{
		"case #1": {
			heights: [][]int{{1}},
			want:    [][]int{{0, 0}},
		},
		"case #2": {
			heights: [][]int{
				{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4},
			},
			want: [][]int{
				{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0},
			},
		},
		"case #3": {
			heights: [][]int{
				{1}, {1},
			},
			want: [][]int{
				{0, 0}, {1, 0},
			},
		},
		"case #4": {
			heights: [][]int{
				{4, 2, 7, 3, 4}, {7, 4, 6, 4, 7}, {6, 3, 5, 3, 6},
			},
			want: [][]int{
				{0, 2}, {0, 4}, {1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4}, {2, 0},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := pacificatlanticwaterflow.PacificAtlantic(test.heights)

			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("PacificAtlantic(%v): expected = %v, got = %v", test.heights, test.want, got)
			}
		})
	}
}
