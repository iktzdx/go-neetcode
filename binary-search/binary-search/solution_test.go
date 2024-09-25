package binarysearch_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/binary-search/binary-search"
)

func Test_Search(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		target int
		want   int
	}{
		"target exists in nums": {
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		"target does not exist in nums": {
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   -1,
		},
		"target is the first element": {
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: -1,
			want:   0,
		},
		"target is the last element": {
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 12,
			want:   5,
		},
		"nums contains a single element": {
			nums:   []int{12},
			target: 12,
			want:   0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.Search(test.nums, test.target)
			if test.want != got {
				t.Fatalf("Search(%v, %d): expected = %d, got = %d", test.nums, test.target, test.want, got)
			}
		})
	}
}
