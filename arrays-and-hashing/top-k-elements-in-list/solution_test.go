package topkelementsinlist_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/arrays-and-hashing/top-k-elements-in-list"
	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	tests := map[string]struct {
		nums []int
		k    int
		want []int
	}{
		"basic test case": {
			nums: []int{10, 10, 10, 22, 22, 31},
			k:    2,
			want: []int{10, 22},
		},
		"want all the numbers": {
			nums: []int{11, 11, 11, 22, 22, 33},
			k:    3,
			want: []int{11, 22, 33},
		},
		"only one element": {
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		"empty input": {
			nums: []int{},
			k:    0,
			want: []int{},
		},
		"negative numbers": {
			nums: []int{-1, -1},
			k:    1,
			want: []int{-1},
		},
		"with zeroes": {
			nums: []int{12, 12, 0, 0, 0, 43, 56},
			k:    2,
			want: []int{0, 12},
		},
		"input without duplicate numbers": {
			nums: []int{11, 22, 33},
			k:    3,
			want: []int{11, 22, 33},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.TopKFrequent(test.nums, test.k)
			if !assert.ElementsMatch(t, test.want, got) {
				t.Errorf("TopKFrequent(%v, %d): expected = %v, got = %v", test.nums, test.k, test.want, got)
			}
		})
	}
}
