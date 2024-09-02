package largestrectangleinhistogram_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/stack/largest-rectangle-in-histogram"
)

func Test_LargestRectangleArea(t *testing.T) {
	tests := map[string]struct {
		heights []int
		want    int
	}{
		"mixed order - largest at the end": {
			heights: []int{7, 1, 7, 2, 2, 4},
			want:    8,
		},
		"increasing order": {
			heights: []int{1, 3, 7},
			want:    7,
		},
		"decreasing order": {
			heights: []int{7, 3, 1},
			want:    7,
		},
		"mixed order - largest in the middle": {
			heights: []int{2, 1, 5, 6, 2, 3},
			want:    10,
		},
		"equal areas": {
			heights: []int{4, 2},
			want:    4,
		},
		"single bin": {
			heights: []int{13},
			want:    13,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.LargestRectangleArea(test.heights)
			if test.want != got {
				t.Fatalf("LargestRectangleArea(%v): expected = %v, got = %v", test.heights, test.want, got)
			}
		})
	}
}
