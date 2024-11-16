package numberofconnectedcomponentsinanundirectedgraph

type dfsFunc func(int)

func CountComponents(n int, edges [][]int) int {
	var (
		dfs   dfsFunc
		count int
	)

	adj := make([][]int, n)
	visited := make([]bool, n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]

		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	dfs = func(node int) {
		for _, nei := range adj[node] {
			if !visited[nei] {
				visited[nei] = true

				dfs(nei)
			}
		}
	}

	for node := range n {
		if !visited[node] {
			visited[node] = true

			dfs(node)

			count++
		}
	}

	return count
}
