package containerwithmostwater_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/two-pointers/container-with-most-water"
)

func Test_MaxArea(t *testing.T) {
	tests := map[string]struct {
		height []int
		want   int
	}{
		"last height is the 2nd endpoint": {
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		"first height is the 1st endpoint": {
			height: []int{7, 1, 2, 5, 4, 7, 6, 3},
			want:   36,
		},
		"max height in the middle": {
			height: []int{1, 3, 42, 69, 2, 1},
			want:   42,
		},
		"min square": {
			height: []int{1, 1},
			want:   1,
		},
		"same heights": {
			height: []int{3, 3, 3, 3},
			want:   9,
		},
		"zero height": {
			height: []int{0, 987654321},
			want:   0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.MaxArea(test.height)
			if test.want != got {
				t.Fatalf("MaxArea(%v): expected = %d, got = %d", test.height, test.want, got)
			}
		})
	}
}
