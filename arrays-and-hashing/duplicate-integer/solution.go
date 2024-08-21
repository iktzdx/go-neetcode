package duplicateinteger

func FindDuplicate(nums []int) int {
	storage := make(map[int]struct{}, len(nums))

	for _, v := range nums {
		if _, ok := storage[v]; ok {
			return v
		}

		storage[v] = struct{}{}
	}

	return -1
}
