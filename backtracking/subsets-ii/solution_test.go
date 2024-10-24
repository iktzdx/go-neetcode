package subsetsii_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	solution "github.com/iktzdx/go-neetcode/backtracking/subsets-ii"
)

func Test_SubsetsWithDup(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want [][]int
	}{
		"case #1": {
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
		"case #2": {
			nums: []int{1, 2, 2},
			want: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
		"case #3": {
			nums: []int{1, 2, 1},
			want: [][]int{{}, {1}, {1, 2}, {1, 1}, {1, 2, 1}, {2}},
		},
		"case #4": {
			nums: []int{7, 7},
			want: [][]int{{}, {7}, {7, 7}},
		},
		"case #5": {
			nums: []int{1, 2, 2, 3},
			want: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {1, 2, 2, 3}, {1, 2, 3}, {1, 3}, {2}, {2, 2}, {2, 2, 3}, {2, 3}, {3}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			for i := range test.want {
				slices.Sort(test.want[i])
			}

			got := solution.SubsetsWithDup(test.nums)

			if !assert.ElementsMatch(t, test.want, got) {
				t.Fatalf("SubsetsWithDup(%v): expected = %v, got = %v", test.nums, test.want, got)
			}
		})
	}
}
