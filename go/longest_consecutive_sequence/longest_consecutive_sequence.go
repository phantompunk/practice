package kata

func longestConsecutive(nums []int) int {
	hashSet := map[int]struct{}{}
	for _, num := range nums {
		hashSet[num] = struct{}{}
	}

	longest := 0
	for n := range hashSet {
		if _, exists := hashSet[n-1]; !exists {
			length := 0
			for {
				if _, ok := hashSet[n+length]; ok {
					length++
					continue
				}
				longest = max(longest, length)
				break
			}
		}
	}
	return longest
}
