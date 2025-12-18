package kata

// ::KATA START::
func longestPalindrome(s string) int {
	freq := map[rune]int{}
	for _, char := range s {
		freq[char]++
	}

	result := 0
	hasMiddle := false
	for _, count := range freq {
		pairs := (count / 2) * 2
		if !hasMiddle && count%2 == 1 {
			hasMiddle = true
			pairs++
		}
		result += pairs
	}
	return result
}

// ::KATA END::
