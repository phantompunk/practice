package kata

// ::KATA START::
func longestPalindrome(s string) int {
	letters := make(map[rune]int)
	for _, r := range s {
		letters[r]++
	}

	result := 0
	hasMiddle := false
	for _, v := range letters {
		result += 2*(v/2)
		if !hasMiddle && v%2==1{
			hasMiddle=true
			result++
		}
	}

	return result
}

// ::KATA END::
