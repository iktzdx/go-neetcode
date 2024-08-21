package twosum_test

import (
	"reflect"
	"testing"

	twosum "github.com/iktzdx/go-neetcode/arrays-and-hashing/two-integer-sum"
)

func Test_TwoSum(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		target int
		want   []int
	}{
		"first two elements": {
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		"last two elements": {
			nums:   []int{7, 9, 3, 2, 4},
			target: 6,
			want:   []int{3, 4},
		},
		"equal elements": {
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		"negative numbers": {
			nums:   []int{4, 9, 1, -2, -6},
			target: 2,
			want:   []int{0, 3},
		},
		"negative target": {
			nums:   []int{-5, 2, 3, 4, -2, 3},
			target: -7,
			want:   []int{0, 4},
		},
		"target is zero": {
			nums:   []int{3, 2, -2, 1},
			target: 0,
			want:   []int{1, 2},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := twosum.TwoSum(test.nums, test.target)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("TwoSum(%v, %d), expect = %v, got = %v", test.nums, test.target, test.want, got)
			}
		})
	}
}
