package subsets

type dfsFunc func(int)

func Subsets(nums []int) [][]int {
	result := make([][]int, 0)
	subset := make([]int, 0)

	var dfs dfsFunc

	dfs = func(idx int) {
		// Base case.
		if idx >= len(nums) {
			c := make([]int, len(subset))
			copy(c, subset)
			result = append(result, c)

			return
		}

		// Decision to include nums[i].
		subset = append(subset, nums[idx])
		dfs(idx + 1)

		// Decision to NOT include nums[i].
		subset = subset[:len(subset)-1]

		dfs(idx + 1)
	}

	dfs(0)

	return result
}
