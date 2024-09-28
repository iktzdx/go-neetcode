package findminimuminrotatedsortedarray

func FindMin(nums []int) int {
	left, right := 0, len(nums)-1
	result := nums[left]

	for left <= right {
		if nums[left] < nums[right] {
			return min(result, nums[left])
		}

		mid := left + (right-left)/2
		result = min(result, nums[mid])

		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}
