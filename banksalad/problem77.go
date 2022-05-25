package main

import (
	"fmt"
	"math"
	"sort"
)

func solution(intervals [][]int) [][]int {
	var merged [][]int

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	for i := range intervals {
		startTime := intervals[i][0]
		if len(merged) != 0 && merged[len(merged)-1][1] >= startTime {
			max := int(math.Max(float64(merged[len(merged)-1][1]), float64(intervals[i][1])))
			merged[len(merged)-1] = []int{
				merged[len(merged)-1][0],
				max,
			}
		} else {
			merged = append(merged, intervals[i])
		}
	}
	return merged
}

func main() {
	array := [][]int{
		{1, 3},
		{5, 8},
		{4, 10},
		{20, 25},
	}
	result := solution(array)
	fmt.Println(result)
}
