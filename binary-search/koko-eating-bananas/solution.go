package kokoeatingbananas

import (
	"math"
	"slices"
)

func MinEatingSpeed(piles []int, h int) int {
	var speed int

	for speed = 1; speed <= slices.Max(piles); speed++ {
		hours := 0

		for _, pile := range piles {
			hours += int(math.Ceil(float64(pile) / float64(speed)))
		}

		if hours <= h {
			break
		}
	}

	return speed
}
