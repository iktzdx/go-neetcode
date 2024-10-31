package kclosestpointstoorigin

import (
	"sort"
)

type point struct {
	dist   int
	coords []int
}

func KClosest(points [][]int, k int) [][]int {
	pts := make([]point, len(points))

	for i, pt := range points {
		pts[i] = point{calcDist(pt[0], pt[1]), []int{pt[0], pt[1]}}
	}

	sort.Slice(pts, func(i, j int) bool {
		return pts[i].dist < pts[j].dist
	})

	result := make([][]int, 0)

	for i := range k {
		result = append(result, pts[i].coords)
	}

	return result
}

func calcDist(x, y int) int {
	return (x * x) + (y * y)
}
