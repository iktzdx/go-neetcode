package courseschedule

type dfsFunc func(node int) bool

// CanFinish checks if all courses can be finished given the prerequisites.
//
// Time complexity: O(n + p), where n is the number of courses and p is the number of prerequisites.
func CanFinish(numCourses int, prerequisites [][]int) bool {
	var dfs dfsFunc

	// Create a map to store prerequisites for each course.
	prereq := make(map[int][]int)

	// Store prerequisites for each course in the map.
	for _, p := range prerequisites {
		prereq[p[0]] = append(prereq[p[0]], p[1])
	}

	// Store prerequisites for each node in the map.
	visited := make(map[int]bool)

	// Store prerequisites for each node in the map.
	dfs = func(crs int) bool {
		// Base case: return false if the current course is already visited (cycle detected).
		if visited[crs] {
			return false
		}

		// Base case: return true if the current course has no prerequisites.
		if prereq[crs] == nil {
			return true
		}

		// Mark the current course as visited.
		visited[crs] = true

		// Recursively explore all prerequisites of the current course.
		for _, pre := range prereq[crs] {
			if !dfs(pre) {
				return false
			}
		}

		// Mark the current course as unvisited after exploring.
		visited[crs] = false

		// Prevent revisiting this course in future DFS calls.
		prereq[crs] = nil

		return true
	}

	// Run DFS recursively on each course since
	// the courses may not be connected (e.g., [1,2], [3,4]).
	for crs := range numCourses {
		if !dfs(crs) {
			return false
		}
	}

	return true
}
