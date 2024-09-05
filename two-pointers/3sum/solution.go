package threeintegersum

import "slices"

func ThreeSum(nums []int) [][]int {
	slices.Sort(nums)

	res := make([][]int, 0)
	for i, n := range nums {
		if n > 0 {
			return res // There is no correct answer.
		}

		if i > 0 && n == nums[i-1] {
			continue // Do not use the same value twice.
		}

		left, right := i+1, len(nums)-1
		for left < right {
			threeSum := n + nums[left] + nums[right]

			if threeSum > 0 {
				right--
				continue
			}

			if threeSum < 0 {
				left++
				continue
			}

			res = append(res, []int{n, nums[left], nums[right]})
			left++

			// Shift only left pointer to avoid using the same value.
			for left < right && nums[left] == nums[left-1] {
				left++
			}
		}
	}

	return res
}
