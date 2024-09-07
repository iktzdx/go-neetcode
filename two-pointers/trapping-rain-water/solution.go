package trappingrainwater

func Trap(height []int) int {
	var out int

	left, right := 0, len(height)-1
	maxLeft, maxRight := height[left], height[right]

	for left < right {
		// If you compare pointers, there is no need to check the result of subtraction.
		if maxLeft < maxRight {
			left++
			maxLeft = max(maxLeft, height[left])
			// The result of computation is never going to be negative,
			// because we are updating max value first.
			out += maxLeft - height[left]
		} else {
			right--
			maxRight = max(maxRight, height[right])
			out += maxRight - height[right]
		}
	}

	return out
}
