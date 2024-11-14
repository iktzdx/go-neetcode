package coursescheduleii

type dfsFunc func(crs int) bool

// Topological sort algorithm.
// Time complexity: O(p + n).
func FindOrder(numCourses int, prerequisites [][]int) []int {
	var dfs dfsFunc

	output := make([]int, 0)

	// Build adjacency list of prerequisites.
	prereq := make(map[int][]int)
	for _, p := range prerequisites {
		prereq[p[0]] = append(prereq[p[0]], p[1])
	}

	visited := make(map[int]bool)
	cycle := make(map[int]bool)

	// A course has 3 possible states:
	//   * visited -> course has been added to output;
	//   * visiting -> course not added to output, but added to cycle;
	//   * unvisited -> course not added to output or cycle.
	dfs = func(crs int) bool {
		if cycle[crs] {
			return false
		}

		if visited[crs] {
			return true
		}

		cycle[crs] = true

		for _, p := range prereq[crs] {
			if !dfs(p) {
				return false
			}
		}

		cycle[crs] = false
		visited[crs] = true

		output = append(output, crs)

		return true
	}

	for crs := range numCourses {
		if !dfs(crs) {
			return []int{}
		}
	}

	return output
}
