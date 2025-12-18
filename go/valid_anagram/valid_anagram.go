package kata

// ::KATA START::
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	chars := map[rune]int{}
	for _, r := range s {
		chars[r]++
	}

	for _, r := range t {
		if _, exists := chars[r]; !exists {
			return false
		}

		chars[r]--
		if chars[r] < 0 {
			return false
		}
	}
	return true
}

// ::KATA END::
