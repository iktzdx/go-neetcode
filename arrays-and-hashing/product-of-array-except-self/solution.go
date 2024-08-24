package productofarrayexceptself

func ProductExceptSelf(nums []int) []int {
	length := len(nums)

	out := make([]int, length)

	prefix := 1
	for i := 0; i < length; i++ {
		out[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := length - 1; i >= 0; i-- {
		out[i] *= postfix
		postfix *= nums[i]
	}

	return out
}
