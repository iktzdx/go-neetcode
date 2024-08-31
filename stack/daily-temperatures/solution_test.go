package dailytemperatures_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	solution "github.com/iktzdx/go-neetcode/stack/daily-temperatures"
)

func Test_DailyTemperatures(t *testing.T) {
	tests := map[string]struct {
		temperatures []int
		want         []int
	}{
		"single measurement": {
			temperatures: []int{30},
			want:         []int{0},
		},
		"all measurements are the same": {
			temperatures: []int{31, 31, 31, 31, 31},
			want:         []int{0, 0, 0, 0, 0},
		},
		"there is no a warmer temperature": {
			temperatures: []int{35, 34, 33, 32, 31},
			want:         []int{0, 0, 0, 0, 0},
		},
		"the temperature is rising all the time": {
			temperatures: []int{30, 40, 50, 60, 70, 80, 90, 100},
			want:         []int{1, 1, 1, 1, 1, 1, 1, 0},
		},
		"the temperature drops and rises": {
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want:         []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.DailyTemperatures(test.temperatures)
			if !assert.ElementsMatch(t, test.want, got) {
				t.Fatalf("DailyTemperatures(%v): expected = %v, got = %v", test.temperatures, test.want, got)
			}
		})
	}
}
