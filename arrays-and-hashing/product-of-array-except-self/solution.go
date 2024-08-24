package productofarrayexceptself

func ProductExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))

	for i := range answer {
		answer[i] = productOfArray(nums, i)
	}

	return answer
}

func productOfArray(nums []int, exclude int) int {
	num := 1

	for i, n := range nums {
		if i == exclude {
			continue
		}

		num *= n
	}

	return num
}
