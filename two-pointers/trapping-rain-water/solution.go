package trappingrainwater

func Trap(height []int) int {
	var maxLeft, maxRight, out int

	left, right := 0, len(height)-1
	for left < right {
		if height[left] > height[right] {
			val := maxRight - height[right]
			if val > 0 {
				out += val
			}

			maxRight = max(maxRight, height[right])
			right--
		} else {
			val := maxLeft - height[left]
			if val > 0 {
				out += val
			}

			maxLeft = max(maxLeft, height[left])
			left++
		}
	}

	return out
}
