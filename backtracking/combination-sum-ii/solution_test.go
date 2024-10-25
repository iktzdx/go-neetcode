package combinationsumii_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	solution "github.com/iktzdx/go-neetcode/backtracking/combination-sum-ii"
)

func Test_CombinationSum2(t *testing.T) {
	tests := map[string]struct {
		candidates []int
		target     int
		want       [][]int
	}{
		"case #1": {
			candidates: []int{1},
			target:     2,
			want:       [][]int{},
		},
		"case #2": {
			candidates: []int{4},
			target:     4,
			want:       [][]int{{4}},
		},
		"case #3": {
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			want:       [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
		"case #4": {
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			want:       [][]int{{1, 2, 2}, {5}},
		},
		"case #5": {
			candidates: []int{9, 2, 2, 4, 6, 1, 5},
			target:     8,
			want:       [][]int{{1, 2, 5}, {2, 2, 4}, {2, 6}},
		},
		"case #6": {
			candidates: []int{1, 2, 3, 4, 5},
			target:     7,
			want:       [][]int{{1, 2, 4}, {2, 5}, {3, 4}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.CombinationSum2(test.candidates, test.target)

			if !assert.ElementsMatch(t, test.want, got) {
				t.Fatalf("CombinationSum2(%v, %d): expected = %v, got = %v", test.candidates, test.target, test.want, got)
			}
		})
	}
}
