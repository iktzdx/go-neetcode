package combinationsum_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	combinationsum "github.com/iktzdx/go-neetcode/backtracking/combination-sum"
)

func Test_CombinationSum(t *testing.T) {
	tests := map[string]struct {
		candidates []int
		target     int
		want       [][]int
	}{
		"case #1": {
			candidates: []int{2},
			target:     1,
			want:       [][]int{},
		},
		"case #2": {
			candidates: []int{2, 3, 5},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		"case #3": {
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want:       [][]int{{2, 2, 3}, {7}},
		},
		"case #4": {
			candidates: []int{3, 4, 5},
			target:     16,
			want:       [][]int{{3, 3, 3, 3, 4}, {3, 3, 5, 5}, {4, 4, 4, 4}, {3, 4, 4, 5}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := combinationsum.CombinationSum(test.candidates, test.target)

			if !assert.ElementsMatch(t, test.want, got) {
				t.Fatalf("CombinationSum(%v, %v): expected = %v, got = %v", test.candidates, test.target, test.want, got)
			}
		})
	}
}
