package numberofconnectedcomponentsinanundirectedgraph

func CountComponents(n int, edges [][]int) int {
	connected := n

	par := make([]int, n) // idx - node, par[idx] - node's parent.
	for idx := range n {
		par[idx] = idx // each node is the parent of itself initially.
	}

	rank := make([]int, n) // idx - parent, rank[idx] - size of connected component.

	find := func(node int) int {
		root := node

		for root != par[root] {
			// path compression
			par[root] = par[par[root]]

			root = par[root]
		}

		return root
	}

	union := func(node1, node2 int) int {
		p1, p2 := find(node1), find(node2)

		if p1 == p2 {
			return 0
		}

		if rank[p2] > rank[p1] {
			p1, p2 = p2, p1
		}

		par[p2] = p1
		rank[p1] += rank[p2]

		return 1
	}

	for _, edge := range edges {
		node1, node2 := edge[0], edge[1]

		connected -= union(node1, node2)
	}

	return connected
}
