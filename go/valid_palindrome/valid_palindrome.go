package kata

import "unicode"

// ::KATA START::
func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	left, right := 0, len(s)-1
	for left <= right {
		for !isAlphanumeric(rune(s[left])) {
			left++
		}

		for !isAlphanumeric(rune(s[right])) {
			right--
		}

		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
      return false
		}

    left++
    right--
	}

	return true
}

func isAlphanumeric(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsLetter(r)
}

// ::KATA END::
