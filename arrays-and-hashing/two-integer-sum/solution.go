package twosum

func TwoSum(nums []int, target int) []int {
	idxs := make([]int, 2)
	storage := make(map[int]int, len(nums))

	for i, v := range nums {
		sub := target - nums[i]

		idx, exists := storage[sub]
		if exists {
			idxs[0], idxs[1] = idx, i

			break
		}

		storage[v] = i
	}

	return idxs
}
