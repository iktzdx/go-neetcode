package medianoftwosortedarrays_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/binary-search/median-of-two-sorted-arrays"
)

func Test_FindMedianSortedArrays(t *testing.T) {
	tests := map[string]struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		"case #1 - slices have different length": {
			nums1: []int{1, 2},
			nums2: []int{3},
			want:  2,
		},
		"case #2 - slices have the same length": {
			nums1: []int{1, 3},
			nums2: []int{2, 4},
			want:  2.5,
		},
		"case #3 - the 1st slice is empty": {
			nums1: []int{},
			nums2: []int{2, 3, 4},
			want:  3,
		},
		"case #4 - the 2nd slice is empty": {
			nums1: []int{11, 22, 33},
			nums2: []int{},
			want:  22,
		},
		"case #5 - merged slice contains a single element": {
			nums1: []int{},
			nums2: []int{1},
			want:  1,
		},
		"case #6 - odd length, contains duplicates after merge": {
			nums1: []int{1, 2, 3, 4, 5},
			nums2: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want:  4,
		},
		"case #7 - even length, contains duplicates after merge": {
			nums1: []int{1, 2, 3, 4},
			nums2: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want:  3.5,
		},
		"case #8 - the most right values of the left partitions are different": {
			nums1: []int{1, 2, 4, 4},
			nums2: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want:  4, // [1 1 2 2 3 4 4 4 5 6 7 8], max(4,3) + min(4,4) / 2
		},
		"case #9 - the most left values of the right partitions are different": {
			nums1: []int{1, 2, 3, 3},
			nums2: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want:  3, // [1 1 2 2 3 3 3 4 5 6 7 8], max(3,3) + min(3,4) / 2
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.FindMedianSortedArrays(test.nums1, test.nums2)
			if test.want != got {
				t.Fatalf("FindMedianSortedArrays(%v, %v): expected = %v, got = %v", test.nums1, test.nums2, test.want, got)
			}
		})
	}
}
