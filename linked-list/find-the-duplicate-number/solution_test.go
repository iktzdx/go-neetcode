package findtheduplicatenumber_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/linked-list/find-the-duplicate-number"
)

func Test_FindDuplicate(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want int
	}{
		"duplicates follow each other": {
			nums: []int{1, 3, 4, 2, 2},
			want: 2,
		},
		"duplicates out of order": {
			nums: []int{3, 1, 3, 4, 2},
			want: 3,
		},
		"the same number": {
			nums: []int{3, 3, 3, 3, 3},
			want: 3,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.FindDuplicate(test.nums)
			if test.want != got {
				t.Fatalf("FindDuplicate(%v): expected = %d, got = %d", test.nums, test.want, got)
			}
		})
	}
}
