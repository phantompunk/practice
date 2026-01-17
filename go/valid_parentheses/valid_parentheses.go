package kata

// ::KATA START::
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}
	for _, p := range s {
		if val, found := pairs[p]; found {
			stack = append(stack, val)
		} else if len(stack) > 0 && stack[len(stack)-1] == p {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

// ::KATA END::
