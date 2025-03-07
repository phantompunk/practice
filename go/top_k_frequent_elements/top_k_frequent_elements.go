package kata

import "fmt"

func topKFrequent(nums []int, k int) []int {
	frequencyMap := map[int]int{}
	for _, num := range nums {
		frequencyMap[num]++
	}

	bucket := make([][]int, len(nums)+1)
	for k, v := range frequencyMap {
		bucket[v] = append(bucket[v], k)
	}

	results := []int{}
	for i := len(bucket) - 1; i > 0; i-- {
		if bucket[i] != nil {
			fmt.Println("i", i, "bucketI", bucket[i])
			results = append(results, bucket[i]...)
			if len(results) >= k {
				results = results[:k]
				return results
			}
		}
	}

	return results
}
