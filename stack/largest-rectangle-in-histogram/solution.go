package largestrectangleinhistogram

func LargestRectangleArea(heights []int) int {
	var maxArea int

	if len(heights) < 2 {
		return heights[0]
	}

	// Append a zero height to ensure we can process all bars in the stack.
	heights = append(heights, 0)

	stack := make([]int, 0) // store only indices

	for i, h := range heights {
		// While the current height is less than the height of the bar at the index
		// on top of the stack, pop the stack and calculate the area.
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			width := i
			if len(stack) > 0 {
				width -= stack[len(stack)-1] + 1
			}

			maxArea = max(maxArea, height*width)
		}

		stack = append(stack, i)
	}

	return maxArea
}
