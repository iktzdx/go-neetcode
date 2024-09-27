package kokoeatingbananas

import (
	"math"
	"slices"
)

func MinEatingSpeed(piles []int, h int) int {
	left, right := 1, slices.Max(piles)
	result := right

	for left <= right {
		k := left + (right-left)/2

		hours := 0
		for _, pile := range piles {
			hours += int(math.Ceil(float64(pile) / float64(k)))
		}

		if hours <= h {
			result = min(result, k)
			right = k - 1
		} else {
			left = k + 1
		}
	}

	return result
}
