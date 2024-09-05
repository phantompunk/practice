package collatzconjecture

import "fmt"

func CollatzConjecture(value int) (int, error) {
	if value <= 0 {
		return 0, fmt.Errorf("Positive number required")
	}

	var steps int = 0
	for {
		if value == 1 {
			break
		} else if value%2 == 0 {
			value = value / 2
		} else {
			value = 3 * value + 1
		}
		steps++
	}
	return steps, nil
}
