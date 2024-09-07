package trappingrainwater_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/two-pointers/trapping-rain-water"
)

func Test_Trap(t *testing.T) {
	tests := map[string]struct {
		height []int
		want   int
	}{
		"distributed over the entire map": {
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		"concentrated in the middle": {
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
		"flat map": {
			height: []int{0, 0, 0, 0},
			want:   0,
		},
		"only one bar": {
			height: []int{9},
			want:   0,
		},
		"shallow slope": {
			height: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			want:   0,
		},
		"deep pit": {
			height: []int{9, 0, 0, 0, 0, 0, 9},
			want:   45,
		},
		"two sections": {
			height: []int{9, 0, 0, 9, 0, 0, 9},
			want:   36,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.Trap(test.height)
			if test.want != got {
				t.Fatalf("Trap(%v): expected = %d, got %d", test.height, test.want, got)
			}
		})
	}
}
