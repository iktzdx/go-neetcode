package slidingwindowmaximum

func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 1 || k == 1 {
		return nums
	}

	result := make([]int, 0)
	mdq := make([]int, 0) // store indexes

	left, right := 0, 0
	for right < len(nums) {
		for len(mdq) > 0 && nums[mdq[len(mdq)-1]] < nums[right] {
			mdq = mdq[:len(mdq)-1]
		}

		mdq = append(mdq, right)

		if left > mdq[0] {
			mdq = mdq[1:]
		}

		if (right + 1) >= k {
			result = append(result, nums[mdq[0]])
			left++
		}

		right++
	}

	return result
}
