package topkelementsinlist_test

import (
	"reflect"
	"testing"

	topkelementsinlist "github.com/iktzdx/go-neetcode/arrays-and-hashing/top-k-elements-in-list"
)

func TestTopKFrequent(t *testing.T) {
	tests := map[string]struct {
		nums []int
		k    int
		want []int
	}{
		"basic test case": {
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
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
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := topkelementsinlist.TopKFrequent(test.nums, test.k)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("TopKFrequent(%v, %d): expected = %v, got = %v", test.nums, test.k, test.want, got)
			}
		})
	}
}
