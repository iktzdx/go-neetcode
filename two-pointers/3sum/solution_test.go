package threeintegersum_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	solution "github.com/iktzdx/go-neetcode/two-pointers/3sum"
)

func Test_ThreeSum(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want [][]int
	}{
		"multiple triplets": {
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		"triplet does not sum up": {
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		"one possible triplet": {
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		"contains duplicates": {
			nums: []int{-3, 3, 4, -3, 1, 2},
			want: [][]int{{-3, 1, 2}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.ThreeSum(test.nums)
			if !assert.ElementsMatch(t, test.want, got) {
				t.Fatalf("ThreeSum(%v): expected = %v, got = %v", test.nums, test.want, got)
			}
		})
	}
}
