package longestconsecutivesequence

import "slices"

func LongestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	slices.Sort(nums)

	counter := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1]-1 {
			counter++
		}
	}

	return counter
}
