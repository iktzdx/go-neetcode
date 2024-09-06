package containerwithmostwater

func MaxArea(height []int) int {
	var maxArea int

	for left := 0; left < len(height)-1; left++ {
		for right := left + 1; right < len(height); right++ {
			maxArea = max(maxArea, min(height[left], height[right])*(right-left))
		}
	}

	return maxArea
}
