package longestconsecutivesequence_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/arrays-and-hashing/longest-consecutive-sequence"
)

func Test_LongestConsecutive(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want int
	}{
		"empty input": {
			nums: []int{},
			want: 0,
		},
		"single element": {
			nums: []int{0},
			want: 1,
		},
		"no sequences": {
			nums: []int{0, 3, 9, 6, -5},
			want: 1,
		},
		"only positive numbers": {
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		"only negative numbers": {
			nums: []int{-100, -4, -200, -1, -3, -2},
			want: 4,
		},
		"signed and unsigned integers": {
			nums: []int{100, 0, -200, -1, 1, 2},
			want: 4,
		},
		"contains duplicates": {
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
		"the same integers": {
			nums: []int{1, 1, 1, 1, 1, 1, 1, 1},
			want: 1,
		},
		"from n to n+1 - sorted": {
			nums: []int{-3, -2, -1, 0, 1, 2, 3},
			want: 7,
		},
		"have multiple sequences #1": {
			nums: []int{0, 3, 2, 1, 100, 200, 5, 6, 7},
			want: 4,
		},
		"have multiple sequences #2": {
			nums: []int{3, 2, 1, 100, 200, 8, 5, 6, 7},
			want: 4,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.LongestConsecutive(test.nums)
			if test.want != got {
				t.Fatalf("LongestConsecutive(%v): expected = %d, got = %d", test.nums, test.want, got)
			}
		})
	}
}
