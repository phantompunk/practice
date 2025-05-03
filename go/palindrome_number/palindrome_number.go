package kata

// Solve without converting to string
func isPalindrome(x int) bool {
	a, b := x, 0
	fn := func(x, y int) int {
		return (y * 10) + x%10
	}

	for a > 0 {
		a, b = a/10, fn(a, b)
	}
	return b == x
}

// Convert to string
// func isPalindrome(x int) bool {
// 	val := fmt.Sprintf("%d", x)
//
// 	l, r := 0, len(val)-1
//
// 	for l <= r {
// 		if val[l] != val[r] {
// 			return false
// 		}
// 		l++
// 		r--
// 	}
// 	return true
// }
