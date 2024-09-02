package largestrectangleinhistogram

func LargestRectangleArea(heights []int) int {
	if len(heights) < 2 {
		return heights[0]
	}

	var maxArea int
	var stack [][2]int // ( index:height )

	for i, h := range heights {
		start := i

		for len(stack) > 0 && stack[len(stack)-1][1] > h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			maxArea = max(maxArea, top[1]*(i-top[0]))
			start = top[0]
		}

		stack = append(stack, [2]int{start, h})
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		maxArea = max(maxArea, top[1]*(len(heights)-top[0]))
	}

	return maxArea
}
