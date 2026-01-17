package kata

// ::KATA START::
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

  for left < right {
    for left < right && !isAlpha(s[left]) {
      left++
    }

    for left < right && !isAlpha(s[right]) {
      right--
    }

    if toLower(s[left]) != toLower(s[right]) {
      return false
    }

    left++
    right--
  }

	return true
}

func isAlpha(l byte) bool {
  if (l >= 'A' && l <= 'Z') || (l >= 'a' && l <= 'z') || (l >= '0' && l <= '9') {
    return true
  }
	return false
}

func toLower(l byte) byte {
  if l >= 'A' && l<='Z' {
    return l + ' ' // 32
  }
	return l
}

// ::KATA END::
