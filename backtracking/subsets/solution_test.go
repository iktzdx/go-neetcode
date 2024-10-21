package subsets_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iktzdx/go-neetcode/backtracking/subsets"
)

func Test_Subsets(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want [][]int
	}{
		"case #1": {
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
		"case #2": {
			nums: []int{7},
			want: [][]int{{}, {7}},
		},
		"case #3": {
			nums: []int{1, 2, 3},
			want: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := subsets.Subsets(test.nums)

			if !assert.ElementsMatch(t, test.want, got) {
				t.Fatalf("Subsets(%v): expected = %v, got = %v", test.nums, test.want, got)
			}
		})
	}
}
