package topkelementsinlist

import "sort"

func TopKFrequent(nums []int, k int) []int {
	if len(nums) < 2 {
		return nums
	}

	freqs := make(map[int]int, 0)
	for _, n := range nums {
		freqs[n]++
	}

	keys := make([]int, 0)
	for k := range freqs {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return freqs[keys[i]] > freqs[keys[j]]
	})

	return keys[:k]
}
