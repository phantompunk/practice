package kata

// ::KATA START::
func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

	ans := []int{}
	for i := len(buckets) - 1; i >= 0 && k > 0; i-- {
		for _, num := range buckets[i] {
			ans = append(ans, num)
			k--
			if k == 0 {
				break
			}
		}
	}

	return ans
}

// ::KATA END::
