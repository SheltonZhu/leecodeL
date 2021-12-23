package _6_merge_intervals

import "sort"

type Intervals [][]int

func (intervals Intervals) Len() int           { return len(intervals) }
func (intervals Intervals) Less(i, j int) bool { return intervals[i][0] < intervals[j][0] }
func (intervals Intervals) Swap(i, j int)      { intervals[i], intervals[j] = intervals[j], intervals[i] }

func merge(intervals [][]int) [][]int {
	sort.Sort(Intervals(intervals))
	var result [][]int
	interval := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if interval[1] >= intervals[i][0] {
			interval[1] = max(intervals[i][1], interval[1])
		} else {
			result = append(result, interval)
			interval = intervals[i]
		}
	}
	result = append(result, interval)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
