package combinationsum

type dfsFunc func(idx int, subset []int, total int)

func CombinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var dfs dfsFunc
	dfs = func(idx int, subset []int, total int) {
		if idx >= len(candidates) || total > target {
			return
		}

		if total == target {
			curr := make([]int, len(subset))

			copy(curr, subset)
			result = append(result, curr)

			return
		}

		// Choose to include the value.
		subset = append(subset, candidates[idx])
		dfs(idx, subset, total+candidates[idx])

		// Choose to NOT include the value.
		subset = subset[:len(subset)-1]
		dfs(idx+1, subset, total)
	}

	dfs(0, []int{}, 0)

	return result
}
