package subsetsii

import "slices"

type backtrackingFunc func(idx int, subset []int)

func SubsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)

	slices.Sort(nums)

	var backtracking backtrackingFunc

	backtracking = func(idx int, subset []int) {
		if idx >= len(nums) {
			c := make([]int, len(subset))
			copy(c, subset)

			result = append(result, c)

			return
		}

		// All subsets that include nums[idx].
		subset = append(subset, nums[idx])
		backtracking(idx+1, subset)

		subset = subset[:len(subset)-1]

		// Skip duplicates.
		for ; (idx + 1) < len(nums); idx++ {
			if nums[idx] != nums[idx+1] {
				break
			}
		}

		// All subsets that do not include nums[idx].
		backtracking(idx+1, subset)
	}

	backtracking(0, nil)

	return result
}
