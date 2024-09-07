package trappingrainwater

func Trap(height []int) int {
	length := len(height)

	maxLeft := make([]int, length)
	maxRight := make([]int, length)

	for i := 1; i < length; i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i-1])
	}

	for i := length - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i+1])
	}

	var out int
	for i := 0; i < length; i++ {
		val := min(maxLeft[i], maxRight[i]) - height[i]
		if val > 0 {
			out += val
		}
	}

	return out
}
