package medianoftwosortedarrays

import "math"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	leftDefault, rigthDefault := math.MinInt, math.MaxInt

	total := len(nums1) + len(nums2)

	smaller, bigger := nums1, nums2

	if len(bigger) < len(smaller) {
		smaller, bigger = bigger, smaller
	}

	left, right := 0, len(smaller)
	for left <= right {
		// Calculate the middle of the smaller slice.
		i := left + (right-left)/2
		// Calculate the middle of the bigger slice:
		// subtract the partition size, then get index.
		j := (total+1)/2 - i

		// Set defaults to -Inf or +Inf to handle cases where indexes are out of bounds.
		sLeft, sRight := leftDefault, rigthDefault
		bLeft, bRight := leftDefault, rigthDefault

		if i > 0 {
			sLeft = smaller[i-1]
		}
		if i < len(smaller) {
			sRight = smaller[i]
		}

		if j > 0 {
			bLeft = bigger[j-1]
		}
		if j < len(bigger) {
			bRight = bigger[j]
		}

		// Check if left partition is correct.
		if sLeft <= bRight && bLeft <= sRight {
			// Even number of elements.
			if total%2 == 0 {
				return float64(max(sLeft, bLeft)+min(sRight, bRight)) / 2
			}

			// Odd number of elements.
			return float64(max(sLeft, bLeft))
		} else if sLeft > bRight {
			// Reduce the size of smaller slice.
			right = i - 1
		} else {
			// Increase the size of the left partition.
			left = i + 1
		}
	}

	return -1
}
