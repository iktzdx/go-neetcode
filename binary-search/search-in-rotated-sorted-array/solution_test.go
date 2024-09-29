package searchinrotatedsortedarray_test

import (
	"testing"

	solutioni "github.com/iktzdx/go-neetcode/binary-search/search-in-rotated-sorted-array"
)

func Test_Search(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		target int
		want   int
	}{
		"case #1 - rotated 4 times, contains target": {
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		"case #2 - rotated 4 times, no target": {
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		"case #3 - single element, no target": {
			nums:   []int{1},
			target: 0,
			want:   -1,
		},
		"case #4 - rotated 1 time, target is the first element": {
			nums:   []int{5, 0, 1, 2, 3, 4},
			target: 5,
			want:   0,
		},
		"case #5 - rotated 2 times, target is the last element": {
			nums:   []int{4, 5, 1, 2, 3},
			target: 3,
			want:   4,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solutioni.Search(test.nums, test.target)
			if test.want != got {
				t.Fatalf("Search(%v, %d): expected = %d, got = %d", test.nums, test.target, test.want, got)
			}
		})
	}
}
