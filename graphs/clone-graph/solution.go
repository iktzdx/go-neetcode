package clonegraph

import "github.com/iktzdx/go-neetcode/graphs"

func CloneGraph(node *graphs.Node) *graphs.Node {
	if node == nil {
		return nil
	}

	root := &graphs.Node{
		Val:       node.Val,
		Neighbors: node.Neighbors,
	}

	// The nodes in the queue are copies of the original nodes,
	// but the neighbors of the nodes in the queue are still the original pointers.
	q := make([]*graphs.Node, 0)

	curr := root
	q = append(q, curr)

	// Maintain a map of visited nodes.
	visited := make(map[int]*graphs.Node)
	visited[curr.Val] = root

	for len(q) > 0 {
		// Pop one node.
		curr, q = q[0], q[1:]

		// Create copies of all its Neighbors,
		// but the Neighbors of Neighbors will have the original pointers for now.
		currNeighbors := make([]*graphs.Node, 0)

		for _, neighbor := range curr.Neighbors {
			var nCopy *graphs.Node

			// Use the same node if we already cloned this node before.
			if visited[neighbor.Val] != nil {
				nCopy = visited[neighbor.Val]
			} else {
				nCopy = &graphs.Node{
					Val:       neighbor.Val,
					Neighbors: neighbor.Neighbors, // Still the original pointers.
				}

				visited[neighbor.Val] = nCopy

				// Queue it, because we have to clone its neighbors.
				q = append(q, nCopy)
			}

			// Store copies of the neighbors.
			currNeighbors = append(currNeighbors, nCopy)
		}

		curr.Neighbors = currNeighbors
	}

	return root
}
