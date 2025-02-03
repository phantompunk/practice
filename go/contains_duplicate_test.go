package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: false,
		},
		{
			name:     "Single element array",
			input:    []int{1},
			expected: false,
		},
		{
			name:     "No duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: false,
		},
		{
			name:     "One duplicate",
			input:    []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "Multiple duplicates",
			input:    []int{1, 2, 3, 1, 2, 4},
			expected: true,
		},
		{
			name:     "All duplicates",
			input:    []int{1, 1, 1, 1, 1},
			expected: true,
		},
		{
			name:     "Negative numbers",
			input:    []int{-1, -2, -3, -1},
			expected: true,
		},
		{
			name:     "Mixed positive negative",
			input:    []int{-1, 2, -3, 2},
			expected: true,
		},
		{
			name:     "Large Numbers",
			input:    []int{1000000000, 2000000000, 1000000000}, // Test for potential integer overflow issues
			expected: true,
		},
		{
			name:     "Large Array No Duplicates",
			input:    generateLargeArray(10000, false), // Test with a larger range to check performance
			expected: false,
		},
		{
			name:     "Large Array With Duplicate",
			input:    generateLargeArray(10000, true), // Introduce a duplicate
			expected: true,
		},
		{
			name:     "Zero present",
			input:    []int{0, 1, 2, 0},
			expected: true,
		},
		{
			name:     "Duplicate at beginning",
			input:    []int{1, 1, 2, 3},
			expected: true,
		},
		{
			name:     "Duplicate at end",
			input:    []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "Duplicate in middle",
			input:    []int{1, 2, 1, 3},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := containsDuplicate(tc.input)
			if got != tc.expected {
				t.Errorf("For input %v, expected %v, got %v", tc.input, tc.expected, got)
			}
		})
	}
}

// Helper function to create large arrays for testing
func generateLargeArray(size int, withDuplicate bool) []int {
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = i
	}
	if withDuplicate {
		nums[size-1] = size / 2 // Introduce a duplicate somewhere in the middle
	}
	return nums
}
