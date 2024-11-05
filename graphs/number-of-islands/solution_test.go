package numberofislands_test

import (
	"testing"

	numberofislands "github.com/iktzdx/go-neetcode/graphs/number-of-islands"
)

func Test_NumIslands(t *testing.T) {
	tests := map[string]struct {
		grid [][]byte
		want int
	}{
		"case #1": {
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		"case #2": {
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
		"case #3": {
			grid: [][]byte{
				{'1', '1', '0', '0', '1'},
				{'1', '1', '0', '0', '1'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 4,
		},
		"case #4": {
			grid: [][]byte{
				{'0', '1', '1', '1', '0'},
				{'0', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := numberofislands.NumIslands(test.grid)

			if test.want != got {
				t.Fatalf("NumIslands(%v): expected = %d, got = %d", test.grid, test.want, got)
			}
		})
	}
}
