package topkelementsinlist

func TopKFrequent(nums []int, k int) []int {
	if len(nums) < 2 {
		return nums
	}

	freqs := make(map[int]int, 0)
	for _, n := range nums {
		freqs[n]++
	}

	bucket := make([][]int, len(nums)+1)
	for num, count := range freqs {
		bucket[count] = append(bucket[count], num)
	}

	result := make([]int, 0)

	for i := len(bucket) - 1; i > 0; i-- {
		if len(bucket[i]) == 0 {
			continue
		}

		for _, num := range bucket[i] {
			if len(result) == k {
				return result
			}

			result = append(result, num)
		}
	}

	return result
}
