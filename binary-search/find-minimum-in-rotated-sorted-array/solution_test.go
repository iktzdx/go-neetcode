package findminimuminrotatedsortedarray_test

import (
	"testing"

	soution "github.com/iktzdx/go-neetcode/binary-search/find-minimum-in-rotated-sorted-array"
)

func Test_FindMin(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want int
	}{
		"case #1 - rotated 1 time": {
			nums: []int{13, 9, 10, 11, 12},
			want: 9,
		},
		"case #2 - rotated 2 times": {
			nums: []int{27, 29, 21, 23, 25},
			want: 21,
		},
		"case #3 - rotated 3 times": {
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		"case #4 - rotated 4 times": {
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		"case #5 - rotated 4 times to original input": {
			nums: []int{11, 13, 15, 17},
			want: 11,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := soution.FindMin(test.nums)
			if test.want != got {
				t.Fatalf("FindMin(%v): expected = %d, got = %d", test.nums, test.want, got)
			}
		})
	}
}
