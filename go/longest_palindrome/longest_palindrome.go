package kata

func longestPalindrome(phrase string) int {
	freqMap := map[rune]int{}
	for _, char := range phrase {
		freqMap[char]++
	}

	result := 0
	hasMiddle := false
	for _, count := range freqMap {
		if count%2 == 0 {
			pairs := count / 2
			result += pairs * 2
		} else if !hasMiddle {
			result++
			hasMiddle = true
		}
	}
	return result
}
