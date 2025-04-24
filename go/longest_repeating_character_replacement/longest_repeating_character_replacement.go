package kata

func characterReplacement(s string, k int) int {
	diff, longest, size := 0, 0, len(s)-1
	for i, char := range s {
		right, total := i+1, 1
		for diff <= k && right <= size {
			if char != rune(s[right]) {
				if diff == k {
					break
				}
				diff++
			}
			total += 1
			right += 1
		}
		diff = 0
		longest = max(longest, total)
	}
	return longest
}
