package twosumiiinputarrayissorted_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	solution "github.com/iktzdx/go-neetcode/two-pointers/two-sum-ii-input-array-is-sorted"
)

func Test_TwoSum(t *testing.T) {
	tests := map[string]struct {
		numbers []int
		target  int
		want    []int
	}{
		"first two elements": {
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		"last two elements": {
			numbers: []int{-15, -11, 2, 7},
			target:  9,
			want:    []int{3, 4},
		},
		"first and last element": {
			numbers: []int{2, 3, 5, 7},
			target:  9,
			want:    []int{1, 4},
		},
		"only two numbers": {
			numbers: []int{2, 7},
			target:  9,
			want:    []int{1, 2},
		},
		"closer to the start": {
			numbers: []int{2, 3, 7, 11},
			target:  9,
			want:    []int{1, 3},
		},
		"closer to the end": {
			numbers: []int{-2, 2, 4, 7},
			target:  9,
			want:    []int{2, 4},
		},
		"zero sum": {
			numbers: []int{-1, 0, 0, 15},
			target:  0,
			want:    []int{2, 3},
		},
		"zero sum - negative numbers": {
			numbers: []int{-11, -3, 1, 2, 3, 8, 10},
			target:  0,
			want:    []int{2, 5},
		},
		"duplicate numbers": {
			numbers: []int{-12, 4, 4, 5, 6, 8, 11, 99},
			target:  12,
			want:    []int{2, 6},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.TwoSum(test.numbers, test.target)
			if !assert.ElementsMatch(t, test.want, got) {
				t.Fatalf("ElementsMatch(%v, %v): expected = %v, got = %v", test.numbers, test.target, test.want, got)
			}
		})
	}
}
