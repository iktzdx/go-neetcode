package productofarrayexceptself

func ProductExceptSelf(nums []int) []int {
	out := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		out[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		out[i] *= postfix
		postfix *= nums[i]
	}

	return out
}
