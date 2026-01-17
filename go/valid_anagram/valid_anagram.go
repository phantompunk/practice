package kata

import "slices"

// ::KATA START::
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letters := make(map[rune]int, len(s))
	for _, char := range s {
		letters[char]++
	}

	for _, char := range t {
		letters[char]--
		if letters[char] < 0 {
			return false
		}
	}
	return true
}

// ::KATA END::

func isAnagramSort(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	return sortStr(s) == sortStr(t)
}

func sortStr(s string) string {
	runes := []rune(s)
	slices.Sort(runes)
	return string(runes)
}
