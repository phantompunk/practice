package kata

func isValid(input string) bool {
	top := -1
	stack := []rune(input)
	ends := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range input {
		if open, exists := ends[char]; exists {
			if top < 0 || stack[top] != open {
				return false
			}
			top--
		} else {
			top++
			stack[top] = char
		}
	}

	return top == -1
}
