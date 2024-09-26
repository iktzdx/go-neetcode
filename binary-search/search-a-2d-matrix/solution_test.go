package searcha2dmatrix_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/binary-search/search-a-2d-matrix"
)

func Test_SearchMatrix(t *testing.T) {
	tests := map[string]struct {
		matrix [][]int
		target int
		want   bool
	}{
		"target is in the left part of a matrix": {
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
			want:   true,
		},
		"target is in the right part of a matrix": {
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 30,
			want:   true,
		},
		"target is in bounds, but not in matrix": {
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 13,
			want:   false,
		},
		"target is greater than the last element of a matrix": {
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 88,
			want:   false,
		},
		"each row contains only one integer": {
			matrix: [][]int{{1}, {10}, {123}, {4321}, {91823}},
			target: 10,
			want:   true,
		},
		"matrix contains only one row with a single element": {
			matrix: [][]int{{1}},
			target: 1,
			want:   true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.SearchMatrix(test.matrix, test.target)
			if test.want != got {
				t.Fatalf("SearchMatrix(%v, %d): expected = %v, got = %v", test.matrix, test.target, test.want, got)
			}
		})
	}
}
