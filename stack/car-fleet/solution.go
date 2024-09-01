package carfleet

import (
	"sort"
)

func CarFleet(target int, position []int, speed []int) int {
	var pairs [][2]int
	for idx, val := range position {
		pairs = append(pairs, [2]int{val, speed[idx]})
	}

	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})

	var stack []int
	for i := len(pairs) - 1; i >= 0; i-- {
		t := (target - pairs[i][0]) / pairs[i][1]

		if len(stack) > 0 && stack[len(stack)-1] >= t {
			continue
		}

		stack = append(stack, t)
	}

	return len(stack)
}
