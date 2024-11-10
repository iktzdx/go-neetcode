package rottingoranges_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/graphs/rotting-oranges"
)

func Test_OrangesRotting(t *testing.T) {
	tests := map[string]struct {
		grid [][]int
		want int
	}{
		"case #1": {
			grid: [][]int{
				{2, 1, 1}, {1, 1, 0}, {0, 1, 1},
			},
			want: 4,
		},
		"case #2": {
			grid: [][]int{
				{2, 1, 1}, {0, 1, 1}, {1, 0, 1},
			},
			want: -1,
		},
		"case #3": {
			grid: [][]int{
				{0, 2},
			},
			want: 0,
		},
		"case #4": {
			grid: [][]int{
				{1, 1, 0}, {0, 1, 1}, {0, 1, 2},
			},
			want: 4,
		},
		"case #5": {
			grid: [][]int{
				{1, 0, 1}, {0, 2, 0}, {1, 0, 1},
			},
			want: -1,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.OrangesRotting(test.grid)

			if test.want != got {
				t.Fatalf("OrangesRotting(%v): expected = %d, got = %d", test.grid, test.want, got)
			}
		})
	}
}
