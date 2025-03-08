package kata

func insert(intervals [][]int, newInterval []int) [][]int {
	results := [][]int{}

	if len(intervals) == 0 {
		results = append(results, newInterval)
		return results
	}

	for _, interval := range intervals {
		if newInterval[0] > interval[1] {
			results = append(results, interval)
		}
	}

	var mStart int
	var mEnd int
	for _, interval := range intervals {
		if mStart == 1 && newInterval[0] > interval[0] && newInterval[0] < interval[1] {
			mStart = interval[0]
		}
		if mStart == 0 && newInterval[0] <= interval[0] {
			mStart = newInterval[0]
		}
		if newInterval[1] > interval[1] {
			mEnd = newInterval[1]
		}
		if newInterval[1] < interval[1] && newInterval[1] >= interval[0] {
			mEnd = interval[1]
		}
	}
	results = append(results, []int{mStart, mEnd})

	for _, interval := range intervals {
		if newInterval[1] < interval[0] {
			results = append(results, interval)
		}
	}
	return results
}
