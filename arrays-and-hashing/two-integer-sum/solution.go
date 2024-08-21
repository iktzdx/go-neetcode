package twosum

func TwoSum(nums []int, target int) []int {
	idxs := make([]int, 2)

	storage := make(map[int]int, len(nums))
	for i, v := range nums {
		// If there are 2 equal elements in the slice,
		// save the index of the last one.
		storage[v] = i
	}

	for i := range nums {
		sub := target - nums[i]
		if idx, exists := storage[sub]; exists && idx != i {
			idxs[0], idxs[1] = i, idx

			break
		}
	}

	return idxs
}
