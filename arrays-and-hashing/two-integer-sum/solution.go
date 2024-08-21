package twosum

func TwoSum(nums []int, target int) []int {
	storage := make(map[int]int, 0)

	for i, v := range nums {
		idx, exists := storage[target-v]
		if exists {
			return []int{idx, i}
		}

		storage[v] = i
	}

	return []int{}
}
