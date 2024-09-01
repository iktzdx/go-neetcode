package carfleet_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/stack/car-fleet"
)

func Test_CarFleet(t *testing.T) {
	tests := map[string]struct {
		target   int
		position []int
		speed    []int
		want     int
	}{
		"multiple car fleets with different arrival time": {
			target:   12,
			position: []int{10, 8, 0, 5, 3},
			speed:    []int{2, 4, 1, 1, 3},
			want:     3,
		},
		"only one car": {
			target:   10,
			position: []int{3},
			speed:    []int{3},
			want:     1,
		},
		"multiple fleets become one": {
			target:   100,
			position: []int{0, 2, 4},
			speed:    []int{4, 2, 1},
			want:     1,
		},
		"two cars - one fleet": {
			target:   10,
			position: []int{6, 8},
			speed:    []int{3, 2},
			want:     1,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.CarFleet(test.target, test.position, test.speed)
			if test.want != got {
				t.Fatalf("CarFleet(%v, %v, %v): expected = %v, got = %v", test.target, test.position, test.speed, test.want, got)
			}
		})
	}
}
