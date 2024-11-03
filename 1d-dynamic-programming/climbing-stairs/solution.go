package climbingstairs

func ClimbStairs(n int) int {
	one, two := 1, 1

	for range n - 1 {
		one, two = one+two, one
	}

	return one
}
