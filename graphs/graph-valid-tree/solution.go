package graphvalidtree

type dfsFunc func(curr, prev int) bool

// Time complexity: O(e + v).
func ValidTree(n int, edges [][]int) bool {
	if n == 0 {
		return true
	}

	var dfs dfsFunc

	adj := make(map[int][]int)

	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	visit := make(map[int]bool)

	dfs = func(curr, prev int) bool {
		if visit[curr] {
			return false
		}

		visit[curr] = true

		for _, j := range adj[curr] {
			if j == prev {
				continue
			}

			if !dfs(j, curr) {
				return false
			}
		}

		return true
	}

	return dfs(0, -1) && len(visit) == n
}
