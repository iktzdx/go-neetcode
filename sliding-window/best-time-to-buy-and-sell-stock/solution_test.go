package besttimetobuyandsellstock_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/sliding-window/best-time-to-buy-and-sell-stock"
)

func Test_MaxProfit(t *testing.T) {
	tests := map[string]struct {
		prices []int
		want   int
	}{
		"volatile price": {
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		"price is constantly decreasing": {
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		"price is constantly rising": {
			prices: []int{1, 3, 4, 6, 7},
			want:   6,
		},
		"single day": {
			prices: []int{9},
			want:   0,
		},
		"the price remains unchanged": {
			prices: []int{9, 9, 9, 9},
			want:   0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.MaxProfit(test.prices)
			if test.want != got {
				t.Fatalf("MaxProfit(%v): expected = %d, got = %d", test.prices, test.want, got)
			}
		})
	}
}
