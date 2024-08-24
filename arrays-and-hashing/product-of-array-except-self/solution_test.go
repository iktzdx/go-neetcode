package productofarrayexceptself_test

import (
	"reflect"
	"testing"

	solution "github.com/iktzdx/go-neetcode/arrays-and-hashing/product-of-array-except-self"
)

func Test_ProductExceptSelf(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want []int
	}{
		"only positive numbers": {
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		"contains negative numbers": {
			nums: []int{-1, -2, 3, 4},
			want: []int{-24, -12, 8, 6},
		},
		"contains zeroes": {
			nums: []int{1, 0, 3, 4},
			want: []int{0, 12, 0, 0},
		},
		"empty input": {
			nums: []int{},
			want: []int{},
		},
		"minimum length": {
			nums: []int{5, 6},
			want: []int{6, 5},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.ProductExceptSelf(test.nums)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("ProductExceptSelf(%v): expected = %v, got = %v", test.nums, test.want, got)
			}
		})
	}
}
