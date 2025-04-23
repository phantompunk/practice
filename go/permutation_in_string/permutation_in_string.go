package kata

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	freqMap := map[rune]int{}
	windowMap := map[rune]int{}
	for _, char := range s1 {
		freqMap[char]++
		windowMap[char]++
	}

	matches := 0
	for i := range 26 {
		count := 0
		if freqMap[i] == windowMap[i] {
		}
		matches
	}

	left, right := 0, len
	for left < right {
		if _, start := freqMap[rune(s1[left])]; start {
			right++
		}
	}
	return false
	// if len(s2) < len(s1) {
	// 	return false
	// }
	//
	// freqMap := map[rune]int
	// for _, char := range s1 {
	// 	freqMap[char]++
	// }
	//
	// for i, char := range s2 {
	// 	if s, start := freqMap[char]; start {
	// 		for _, next := range s2[i:] {
	// 			if _, continues := freqMap[next]; continues {
	//
	// 			}
	// 		}
	// 	}
	// }
}
