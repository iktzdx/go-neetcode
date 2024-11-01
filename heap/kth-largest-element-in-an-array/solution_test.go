package kthlargestelementinanarray_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/heap/kth-largest-element-in-an-array"
)

func Test_FindKthLargest(t *testing.T) {
	tests := map[string]struct {
		nums []int
		k    int
		want int
	}{
		"case #1": {
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		"case #2": {
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
		"case #3": {
			nums: []int{2, 3, 1, 5, 4},
			k:    2,
			want: 4,
		},
		"case #4": {
			nums: []int{2, 3, 1, 1, 5, 5, 4},
			k:    3,
			want: 4,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.FindKthLargest(test.nums, test.k)

			if test.want != got {
				t.Fatalf("FindKthLargest(%v, %d): expected = %d, got = %d", test.nums, test.k, test.want, got)
			}
		})
	}
}
