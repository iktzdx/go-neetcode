package climbingstairs

type dfsFunc func(int) int

func ClimbStairs(n int) int {
	var dfs dfsFunc

	dp := make(map[int]int)

	dfs = func(i int) int {
		if i == n {
			return 1
		}

		if i > n {
			return 0
		}

		if val, ok := dp[i]; ok {
			return val
		}

		dp[i] = dfs(i+1) + dfs(i+2)

		return dp[i]
	}

	return dfs(0)
}
