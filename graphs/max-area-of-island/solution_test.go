package maxareaofisland_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/graphs/max-area-of-island"
)

func Test_MaxAreaOfIsland(t *testing.T) {
	tests := map[string]struct {
		grid [][]int
		want int
	}{
		"case #1": {
			grid: [][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			},
			want: 6,
		},
		"case #2": {
			grid: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: 0,
		},
		"case #3": {
			grid: [][]int{
				{0, 1, 1, 0, 1},
				{1, 0, 1, 0, 1},
				{0, 1, 1, 0, 1},
				{0, 1, 0, 0, 1},
			},
			want: 6,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.MaxAreaOfIsland(test.grid)

			if test.want != got {
				t.Fatalf("MaxAreaOfIsland(%v): expected = %d, got = %d", test.grid, test.want, got)
			}
		})
	}
}
