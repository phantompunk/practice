package kata

func insert(intervals [][]int, newInterval []int) [][]int {
	results := [][]int{}

	for i, interval := range intervals {
		if newInterval[1] < interval[0] {
			results = append(results, newInterval)
			results = append(results, intervals[i:]...)
			return results
		}
		if newInterval[0] > interval[1] {
			results = append(results, interval)
			continue
		}
		newInterval[0] = min(newInterval[0], interval[0])
		newInterval[1] = max(newInterval[1], interval[1])
	}
	return append(results, newInterval)
}
