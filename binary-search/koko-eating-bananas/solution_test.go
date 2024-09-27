package kokoeatingbananas_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/binary-search/koko-eating-bananas"
)

func Test_MinEatingSpeed(t *testing.T) {
	tests := map[string]struct {
		piles []int
		h     int
		want  int
	}{
		"case #1": {
			piles: []int{3, 6, 7, 11},
			h:     8,
			want:  4,
		},
		"case #2": {
			piles: []int{30, 11, 23, 4, 20},
			h:     5,
			want:  30,
		},
		"case #3": {
			piles: []int{30, 11, 23, 4, 20},
			h:     6,
			want:  23,
		},
		"case #4": {
			piles: []int{2, 2},
			h:     2,
			want:  2,
		},
		"case #5": {
			piles: []int{312884470},
			h:     312884469,
			want:  2,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.MinEatingSpeed(test.piles, test.h)
			if test.want != got {
				t.Fatalf("MinEatingSpeed(%v, %d): expected = %d, got = %d", test.piles, test.h, test.want, got)
			}
		})
	}
}
