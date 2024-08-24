package productofarrayexceptself

func ProductExceptSelf(nums []int) []int {
	length := len(nums)

	pre := make([]int, length)
	post := make([]int, length)
	out := make([]int, length)

	for i := 0; i < length; i++ {
		if i == 0 {
			pre[i] = 1 * nums[i]
			continue
		}

		pre[i] = pre[i-1] * nums[i]
	}

	for i := length - 1; i >= 0; i-- {
		if i == length-1 {
			post[i] = 1 * nums[i]
			continue
		}

		post[i] = post[i+1] * nums[i]
	}

	for i := 0; i < length; i++ {
		defaultMultiplier := 1

		if i == 0 {
			out[i] = defaultMultiplier * post[i+1]
			continue
		}

		if i == length-1 {
			out[i] = pre[i-1] * defaultMultiplier
			continue
		}

		out[i] = pre[i-1] * post[i+1]
	}

	return out
}
