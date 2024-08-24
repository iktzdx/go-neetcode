package productofarrayexceptself

func ProductExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	for i := range answer {
		answer[i] = 1
	}

	i, j := 0, 0
	for i < len(nums) && j < len(nums) {
		if j != i {
			answer[i] *= nums[j]
		}

		if j == len(nums)-1 {
			i++
			j = 0

			continue
		}

		j++
	}

	return answer
}
