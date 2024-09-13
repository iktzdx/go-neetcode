package slidingwindowmaximum

func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 1 || k == 1 {
		return nums
	}

	result := make([]int, 0)

	for left, right := 0, k; right <= len(nums); right++ {
		result = append(result, findMaxValue(nums[left:right]))
		left++
	}

	return result
}

func findMaxValue(nums []int) int {
	var maxValue int

	for _, n := range nums {
		maxValue = max(maxValue, n)
	}

	return maxValue
}
