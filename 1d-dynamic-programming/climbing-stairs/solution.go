package climbingstairs

type dfsFunc func(int) int

func ClimbStairs(n int) int {
	var dfs dfsFunc

	dfs = func(i int) int {
		if i == n {
			return 1
		}

		if i < n {
			return dfs(i+1) + dfs(i+2)
		}

		return 0
	}

	return dfs(0)
}
