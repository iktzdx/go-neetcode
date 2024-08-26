package longestconsecutivesequence

import (
	"slices"
)

func LongestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	slices.Sort(nums)

	longest := 1
	for i := 0; i < len(nums)-1; i++ {
		length := 0
		for i+length < len(nums) && nums[i] == nums[i+length]-length {
			length++
		}

		longest = max(longest, length)
	}

	return longest
}
