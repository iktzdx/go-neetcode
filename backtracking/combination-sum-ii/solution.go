package combinationsumii

import "slices"

type backtrackFunc func(pos int, subset []int, total int)

func CombinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)

	result := make([][]int, 0)

	var backtrack backtrackFunc
	backtrack = func(pos int, subset []int, total int) {
		if total == target {
			subsetCopy := make([]int, len(subset))
			copy(subsetCopy, subset)

			result = append(result, subsetCopy)

			return
		}

		if total > target || pos >= len(candidates) {
			return
		}

		// Include candidates[pos].
		subset = append(subset, candidates[pos])
		backtrack(pos+1, subset, total+candidates[pos])

		// Skip candidates[pos].
		subset = subset[:len(subset)-1]

		for ; pos+1 < len(candidates); pos++ {
			if candidates[pos] != candidates[pos+1] {
				break
			}
		}

		backtrack(pos+1, subset, total)
	}

	backtrack(0, nil, 0)

	return result
}
