package permutations_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	solution "github.com/iktzdx/go-neetcode/backtracking/permutations"
)

func Test_Permute(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want [][]int
	}{
		"case #1": {
			nums: []int{1},
			want: [][]int{{1}},
		},
		"case #2": {
			nums: []int{0, 1},
			want: [][]int{{0, 1}, {1, 0}},
		},
		"case #3": {
			nums: []int{1, 2, 3},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.Permute(test.nums)

			if !assert.ElementsMatch(t, test.want, got) {
				t.Fatalf("Permute(%v): expected = %v, got = %v", test.nums, test.want, got)
			}
		})
	}
}
