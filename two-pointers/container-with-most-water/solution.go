package containerwithmostwater

func MaxArea(height []int) int {
	var maxArea int

	left, right := 0, len(height)-1
	for left < right {
		maxArea = max(maxArea, (right-left)*min(height[right], height[left]))

		if height[right] > height[left] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
