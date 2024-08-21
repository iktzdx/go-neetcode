package twosum

func TwoSum(nums []int, target int) []int {
	idxs := make([]int, 2)

	idx := 0
	for i := idx + 1; i < len(nums); i++ {
		if i == idx {
			continue
		}

		sum := nums[idx] + nums[i]
		if target == sum {
			idxs[0], idxs[1] = idx, i

			break
		}

		if i == len(nums)-1 {
			idx++
			i = 0
		}
	}

	return idxs
}
