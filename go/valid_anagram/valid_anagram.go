package kata

func isAnagram(s, t string) bool {
	freqA := map[rune]int{}

	for _, char := range s {
		freqA[char]++
	}

	for _, char := range t {
		if _, exists := freqA[char]; exists {
			freqA[char]--
		} else {
			freqA[char]++
		}
		if freqA[char] == 0 {
			delete(freqA, char)
		}
	}
	return len(freqA) == 0
}
