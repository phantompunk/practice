package kata

var results []string

func generateParenthesis(n int) []string {
	results = []string{}
	backtrack("", 0, 0, n)

	return results
}

func backtrack(current string, open, closed, number int) {
	if len(current) == number*2 {
		results = append(results, current)
		return
	}

	if open < number {
		backtrack(current+"(", open+1, closed, number)
	}

	if closed < open {
		backtrack(current+")", open, closed+1, number)
	}
}
