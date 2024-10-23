package permutations

import "slices"

func Permute(nums []int) [][]int {
	result := make([][]int, 0)

	if len(nums) == 1 {
		return append(result, []int{nums[0]})
	}

	for _, p := range Permute(nums[1:]) {
		for i := range len(p) + 1 {
			pCopy := make([]int, len(p))
			copy(pCopy, p)

			result = append(result, slices.Insert(pCopy, i, nums[0]))
		}
	}

	return result
}
